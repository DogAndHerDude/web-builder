package db

import (
	"time"
)

type HTMLElement int

const (
	A HTMLElement = iota
	Abbr
	Address
	Area
	Article
	Aside
	Audio
	B
	Base
	Bdi
	Bdo
	Blockquote
	Body
	Br
	Button
	Canvas
	Caption
	Cite
	Code
	Col
	Colgroup
	Data
	Datalist
	Dd
	Del
	Details
	Dfn
	Dialog
	Div
	Dl
	Dt
	Em
	Embed
	Fieldset
	Figcaption
	Figure
	Footer
	Form
	H1
	H2
	H3
	H4
	H5
	H6
	Head
	Header
	Hr
	Html
	I
	Iframe
	Img
	Input
	Ins
	Kbd
	Label
	Legend
	Li
	Link
	Main
	Map
	Mark
	Menu
	Menuitem
	Meta
	Meter
	Nav
	Noscript
	Object
	Ol
	Optgroup
	Option
	Output
	P
	Param
	Picture
	Pre
	Progress
	Q
	Rp
	Rt
	Ruby
	S
	Samp
	Script
	Section
	Select
	Slot
	Small
	Source
	Span
	Strong
	Style
	Sub
	Summary
	Sup
	Table
	Tbody
	Td
	Template
	Textarea
	Tfoot
	Th
	Thead
	Time
	Title
	Tr
	Track
	U
	Ul
	Var
	Video
	Wbr
	Text
	// Internal template elements
)

func (e HTMLElement) String() string {
	elements := [...]string{
		"A", "ABBR", "ADDRESS", "AREA", "ARTICLE", "ASIDE", "AUDIO",
		"B", "BASE", "BDI", "BDO", "BLOCKQUOTE", "BODY", "BR", "BUTTON",
		"CANVAS", "CAPTION", "CITE", "CODE", "COL", "COLGROUP", "DATA",
		"DATALIST", "DD", "DEL", "DETAILS", "DFN", "DIALOG", "DIV", "DL", "DT",
		"EM", "EMBED", "FIELDSET", "FIGCAPTION", "FIGURE", "FOOTER", "FORM",
		"H1", "H2", "H3", "H4", "H5", "H6", "HEAD", "HEADER", "HR", "HTML",
		"I", "IFRAME", "IMG", "INPUT", "INS", "KBD", "LABEL", "LEGEND", "LI",
		"LINK", "MAIN", "MAP", "MARK", "MENU", "MENUITEM", "META", "METER",
		"NAV", "NOSCRIPT", "OBJECT", "OL", "OPTGROUP", "OPTION", "OUTPUT", "P",
		"PARAM", "PICTURE", "PRE", "PROGRESS", "Q", "RP", "RT", "RUBY", "S",
		"SAMP", "SCRIPT", "SECTION", "SELECT", "SLOT", "SMALL", "SOURCE", "SPAN",
		"STRONG", "STYLE", "SUB", "SUMMARY", "SUP", "TABLE", "TBODY", "TD",
		"TEMPLATE", "TEXTAREA", "TFOOT", "TH", "THEAD", "TIME", "TITLE", "TR",
		"TRACK", "U", "UL", "VAR", "VIDEO", "WBR",
	}

	if int(e) < len(elements) {
		return elements[e]
	}

	return "UNKNOWN"
}

type TemplateElement string

const (
	Outlet TemplateElement = "OUTLET"
)

type TemplateOrHTMLElement interface {
	isElement()
}

func (HTMLElement) isElement()     {}
func (TemplateElement) isElement() {}

type User struct {
	ID        string
	Email     string
	Password  string
	Salt      string
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type Site struct {
	ID              string
	Title           string
	Header          []*HTMLNode
	Pages           []*Page
	Footer          []*HTMLNode
	SharedNodes     []*HTMLNode
	Template        *SiteTemplate
	IsPublished     bool `db:"is_published"`
	Repository      string
	UserID          string    `db:"user_id"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
	LastPublishedAt time.Time `db:"last_published_at"`
}

type PageType string

const (
	Static    PageType = "STATIC"
	Portfolio PageType = "PORTFOLIO"
	Blog      PageType = "BLOG"
)

type Page struct {
	ID           string
	Type         PageType
	Title        string
	Slug         string
	Dependencies []string
	Body         []*HTMLNode // Should be binary
	Pages        []*Page
	PageID       string    `db:"page_id"`
	SiteID       string    `db:"site_id"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

type NodeType string

const (
	Container NodeType = "CONTAINER"
	Component NodeType = "COMPONENT"
)

type HTMLNode struct {
	Tag              HTMLElement
	TextContent      string
	Dependency       string
	Attributes       map[string]string // key=value values
	ComponentID      string
	ComponentVersion string
	ClassList        []string
	Children         []*HTMLNode
}

type TemplateNode struct {
	Tag              TemplateOrHTMLElement
	TextContent      string
	Dependency       string
	Attributes       map[string]string // key=value values
	ComponentID      string
	ComponentVersion string
	ClassList        []string
	Children         []*HTMLNode
}

type CustomTemplate struct {
	ID        string
	Nodes     []*TemplateNode
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SiteTemplate struct {
	ID            string
	Nodes         []*TemplateNode
	Pallete       []string
	FontFamily    string
	FontFamilyURL string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
