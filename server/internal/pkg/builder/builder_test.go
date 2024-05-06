package builder

import (
	"strings"
	"testing"

	"github.com/DogAndHerDude/web-builder/db"
)

func TestTreeNodeToHTMLTagOnlyText(t *testing.T) {
	node := &db.HTMLNode{
		Tag:         "#text",
		TextContent: "Hello!",
	}
	expected := `
  Hello!
  `
	result, err := htmlNodeToHTMLString(node)
	if err != nil {
		t.Fatalf(`treeToHTMLTag expected not to error, but received %v`, err)
	}

	if !strings.Contains(expected, result) {
		t.Fatalf(`treeToHTMLTag expected templates to match but they did not`)
	}
}

// TestTreeNodeToHTMLTagWithChild creates a template
// with tag and child in it
func TestTreeNodeToHTMLTagWithChild(t *testing.T) {
	children := make([]*db.HTMLNode, 1)
	children[0] = &db.HTMLNode{
		Tag:         "#text",
		TextContent: "Hi",
	}
	node := &db.HTMLNode{
		Tag:      "DIV",
		Children: children,
	}
	expected := `
  <div>
    Hi
  </div>
  `
	result, err := htmlNodeToHTMLString(node)
	if err != nil {
		t.Fatalf(`treeToHTMLTag expected not to error, but received: %v`, err)
	}

	if !strings.Contains(expected, result) {
		t.Fatalf(`treeToHTMLTag expected templates to match but they did not`)
	}
}

// TestTreeNodeToHTMLTagWithClass creates a template with class names
func TestTreeNodeToHTMLTagWithClass(t *testing.T) {
	classes := make([]string, 2)
	classes[0] = "class1"
	classes[1] = "class2"
	children := make([]*db.HTMLNode, 1)
	children[0] = &db.HTMLNode{
		Tag:         "#text",
		TextContent: "Hi",
	}
	node := &db.HTMLNode{
		Tag:       "DIV",
		Children:  children,
		ClassList: classes,
	}
	expected := `
  <div class="class1 class2">
    Hi
  </div>
  `
	result, err := htmlNodeToHTMLString(node)
	if err != nil {
		t.Fatalf(`treeToHTMLTag expected not to error, but received %v`, err)
	}

	if !strings.Contains(expected, result) {
		t.Fatalf(`treeToHTMLTag expected templates to match but they did not`)
	}
}

func TestBuildPageHTMLToReturnContentInHTML(t *testing.T) {
	classes := make([]string, 2)
	classes[0] = "class1"
	classes[1] = "class2"
	children := make([]*db.HTMLNode, 1)
	children[0] = &db.HTMLNode{
		Tag:         "#text",
		TextContent: "Hi",
	}
	node := &db.HTMLNode{
		Tag:       "DIV",
		Children:  children,
		ClassList: classes,
	}
	nodes := make([]*db.HTMLNode, 1)
	nodes[0] = node
	page := &db.Page{
		Title: "Hello",
		Body:  nodes,
	}
	expected := `
  <!doctype html>
  <html>
    <head>
      <meta charset="UTF-8">
      <meta name="viewport" content="width=device-width, initial-scale=1.0">
      <title>Hello</title>
      <script src="https://cdn.tailwindcss.com"></script>
    </head>
    <body>
      <div class="class1 class2">
        Hi
      </div>
    </body>
  </html>
  `
	result, err := buildPageHTML(page)
	if err != nil {
		t.Fatalf(`buildPageHTML expected not to error, but received %v`, err)
	}

	if !strings.Contains(strings.ReplaceAll(strings.ReplaceAll(expected, "\n", ""), " ", ""), strings.ReplaceAll(strings.ReplaceAll(result, "\n", ""), " ", "")) {
		t.Fatalf(`buildPageHTML expected templates to match but they did not`)
	}
}

func TestBuildNodeHTMLTagWithAttributes(t *testing.T) {
	classes := make([]string, 2)
	classes[0] = "class1"
	classes[1] = "class2"
	children := make([]*db.HTMLNode, 1)
	children[0] = &db.HTMLNode{
		Tag:         "#text",
		TextContent: "Hi",
	}
	node := &db.HTMLNode{
		Tag:        "DIV",
		Children:   children,
		ClassList:  classes,
		Attributes: map[string]string{"x-trigger": "click"},
	}
	expected := `
  <div class="class1 class2" x-trigger="click">
    Hi
  </div>
  `
	result, err := htmlNodeToHTMLString(node)
	if err != nil {
		t.Fatalf(`buildPageHTML expected not to error, but received %v`, err)
	}

	if !strings.Contains(strings.ReplaceAll(strings.ReplaceAll(expected, "\n", ""), " ", ""), strings.ReplaceAll(strings.ReplaceAll(result, "\n", ""), " ", "")) {
		t.Fatalf(`buildPageHTML expected templates to match but they did not`)
	}
}
