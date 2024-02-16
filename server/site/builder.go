package site

import (
	"bytes"
	"html/template"
	"strings"
)

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
  {{ tagOpen .Tag .Class }}
    {{ range .Children }}{{ unescaped . }}{{ end }}
  {{ tagClose .Tag }}`
)

var funcMap = template.FuncMap{
	"unescaped": func(value string) template.HTML {
		return template.HTML(value)
	},
	"tagOpen": func(name string, class string) template.HTML {
		tag := "<" + template.HTMLEscapeString(name)

		if class != "" {
			tag = tag + ` class="` + template.HTMLEscapeString(class) + `"`
		}

		tag = tag + ">"

		return template.HTML(tag)
	},
	"tagClose": func(name string) template.HTML {
		tag := "</" + template.HTMLEscapeString(name) + ">"

		return template.HTML(tag)
	},
}

func treeNodeToHTML(node *TreeNode) (string, error) {
	if node.Tag == "#text" {
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
			childStr, childErr := treeNodeToHTML(child)

			if childErr != nil {
				return "", childErr
			}

			children = append(children, childStr)
		}
	}

	data := NodeTemplate{
		Tag:      strings.ToLower(string(node.Tag)),
		Class:    class,
		Children: children,
	}
	var parsedTempalte bytes.Buffer
	parseErr := tpl.Execute(&parsedTempalte, data)
	if parseErr != nil {
		return "", parseErr
	}

	return parsedTempalte.String(), nil
}

func buildPageHTML(page *Page) (string, error) {
	tpl, err := template.New(pageTemplateName).Funcs(funcMap).Parse(pageTemplate)
	if err != nil {
		return "", err
	}

	children := make([]string, 0)

	for _, node := range page.Nodes {
		childStr, err := treeNodeToHTML(node)
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

func buildPageTree(page *Page) (*PageOutput, error) {
	output := &PageOutput{
		Slug: page.Slug,
	}

	if len(page.Nodes) > 0 {
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

			subPageOutput := &PageOutput{
				Slug:    subPage.Slug,
				Content: subPageHTML,
			}
			output.SubPages = append(output.SubPages, subPageOutput)
		}
	}

	return output, nil
}

func BuildSite(site *Site) (*SiteOutput, error) {
	output := make([]*PageOutput, 0)

	for _, page := range site.Pages {
		pageOutput, err := buildPageTree(page)
		if err != nil {
			return nil, err
		}

		output = append(output, pageOutput)
	}

	siteOutput := &SiteOutput{
		Pages: output,
		// TODO: build sitemap
	}

	return siteOutput, nil
}
