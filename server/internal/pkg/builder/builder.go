package builder

import (
	"bytes"
	"html/template"
	"strings"
	"sync"

	"github.com/DogAndHerDude/web-builder/internal/app/db"
)

type NodeTemplate struct {
	Tag        string
	Class      string
	Children   []string
	Attributes map[string]string
}

type PageTemplate struct {
	Title    string
	Children []string
}

type PageBuildResult struct {
	Slug    string
	Content string
	Pages   []*PageBuildResult
}

type PageBuildError struct {
	PageID string
	Error  error
}

type BuildResult struct {
	Pages   []*PageBuildResult
	Errors  []PageBuildError
	SiteMap string
}

type SiteBuilder interface {
	BuildSite(site *db.Site) (*BuildResult, error)
	BuildSiteConcurrent(site *db.Site) (*BuildResult, error)
}

type SiteBuilderService struct{}

const (
	pageTemplateName = "PageTemaplate"
	pageTemplate     = `
  <!doctype html>
  <html>
    <head>
      <meta charset="UTF-8">
      <meta name="viewport" content="width=device-width, initial-scale=1.0">
      <title>{{ .Title }}</title>
      <script src="https://cdn.tailwindcss.com"></script>
    </head>
    <body>
     {{ range .Children }}{{ unescaped . }}{{ end }}
    </body>
  </html>
  `
)

const (
	nodeTemplateName = "NodeTemplate"
	nodeTemplate     = `
  {{ tagOpen .Tag .Class .Attributes }}
    {{ range .Children }}{{ unescaped . }}{{ end }}
  {{ tagClose .Tag }}`
)

var funcMap = template.FuncMap{
	"unescaped": func(value string) template.HTML {
		return template.HTML(value)
	},
	"tagOpen": func(name string, class string, attributes map[string]string) template.HTML {
		tag := "<" + template.HTMLEscapeString(name)

		if class != "" {
			tag = tag + ` class="` + template.HTMLEscapeString(class) + `"`
		}

		for k, v := range attributes {
			tag = tag + " " + k + `="` + template.HTMLEscapeString(v) + `"`
		}

		tag = tag + ">"

		return template.HTML(tag)
	},
	"tagClose": func(name string) template.HTML {
		tag := "</" + template.HTMLEscapeString(name) + ">"

		return template.HTML(tag)
	},
}

func htmlNodeToHTMLString(node *db.HTMLNode) (string, error) {
	if node.Tag == db.Text {
		return node.TextContent, nil
	}

	tpl, err := template.New(nodeTemplateName).Funcs(funcMap).Parse(nodeTemplate)
	if err != nil {
		return "", err
	}

	var class string
	children := make([]string, 0)

	if node.ClassList != nil && len(node.ClassList) > 0 {
		class = strings.Join(node.ClassList, " ")
	}

	if node.Children != nil && len(node.Children) > 0 {
		for _, child := range node.Children {
			childStr, childErr := htmlNodeToHTMLString(child)

			if childErr != nil {
				return "", childErr
			}

			children = append(children, childStr)
		}
	}

	data := NodeTemplate{
		Tag:        strings.ToLower(string(node.Tag)),
		Class:      class,
		Children:   children,
		Attributes: node.Attributes,
	}
	var parsedTempalte bytes.Buffer
	parseErr := tpl.Execute(&parsedTempalte, data)
	if parseErr != nil {
		return "", parseErr
	}

	return parsedTempalte.String(), nil
}

func buildPageHTML(page *db.Page) (string, error) {
	tpl, err := template.New(pageTemplateName).Funcs(funcMap).Parse(pageTemplate)
	if err != nil {
		return "", err
	}

	children := make([]string, 0)

	for _, node := range page.Body {
		childStr, err := htmlNodeToHTMLString(node)
		if err != nil {
			return "", err
		}

		children = append(children, childStr)
	}

	data := PageTemplate{
		Title:    page.Title,
		Children: children,
	}
	var parsedTemplate bytes.Buffer
	parseErr := tpl.Execute(&parsedTemplate, data)

	if parseErr != nil {
		return "", parseErr
	}

	return parsedTemplate.String(), nil
}

func buildPageTree(page *db.Page) (*PageBuildResult, error) {
	output := &PageBuildResult{
		Slug: page.Slug,
	}

	if len(page.Body) > 0 {
		pageHTML, err := buildPageHTML(page)
		if err != nil {
			return nil, err
		}

		output.Content = pageHTML
	}

	if len(page.Pages) > 0 {
		for _, subPage := range page.Pages {
			subPageHTML, err := buildPageHTML(subPage)
			if err != nil {
				return nil, err
			}

			subPageOutput := &PageBuildResult{
				Slug:    subPage.Slug,
				Content: subPageHTML,
			}
			output.Pages = append(output.Pages, subPageOutput)
		}
	}

	return output, nil
}

func (s *SiteBuilderService) BuildSite(site *db.Site) (*BuildResult, error) {
	output := make([]*PageBuildResult, 0)

	for _, page := range site.Pages {
		pageOutput, err := buildPageTree(page)
		if err != nil {
			return nil, err
		}

		output = append(output, pageOutput)
	}

	siteOutput := &BuildResult{
		Pages: output,
		// TODO: build sitemap
	}

	return siteOutput, nil
}

func buildPageTreeConcurrent(page *db.Page, output chan<- *PageBuildResult, buildErr chan<- PageBuildError, wg *sync.WaitGroup) {
	defer wg.Done()
	pageOutput := &PageBuildResult{
		Slug: page.Slug,
	}

	if len(page.Body) > 0 {
		pageHTML, err := buildPageHTML(page)
		if err != nil {
			buildErr <- PageBuildError{
				PageID: page.ID,
				Error:  err,
			}
			return
		}

		pageOutput.Content = pageHTML
	}

	if len(page.Pages) > 0 {
		for _, subPage := range page.Pages {
			subPageHTML, err := buildPageHTML(subPage)
			if err != nil {
				buildErr <- PageBuildError{
					PageID: subPage.ID,
					Error:  err,
				}
				return
			}

			subPageOutput := &PageBuildResult{
				Slug:    subPage.Slug,
				Content: subPageHTML,
			}
			pageOutput.Pages = append(pageOutput.Pages, subPageOutput)
		}
	}

	output <- pageOutput
}

func (b *SiteBuilderService) BuildSiteConcurrent(site *db.Site) *BuildResult {
	wg := &sync.WaitGroup{}
	var pages []*PageBuildResult
	buildErrors := []PageBuildError{}
	outputChans := make(chan *PageBuildResult, len(site.Pages))
	errorChans := make(chan PageBuildError, len(site.Pages))

	defer close(outputChans)
	defer close(errorChans)

	for _, page := range site.Pages {
		wg.Add(1)
		go buildPageTreeConcurrent(page, outputChans, errorChans, wg)
	}

	select {
	case page := <-outputChans:
		pages = append(pages, page)
	case err := <-errorChans:
		// Errors get overriden here
		// Should use a stack perhaps?
		// Exit early maybe if one error occurs?
		// can't build the rest of the site though
		buildErrors = append(buildErrors, err)
	}

	wg.Wait()

	siteOutput := &BuildResult{
		Pages:  pages,
		Errors: buildErrors,
		// TODO: build sitemap
	}

	return siteOutput
}

func New() *SiteBuilderService {
	return &SiteBuilderService{}
}
