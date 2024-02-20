package db

import (
	"time"
)

type HTMLElement string

const (
	A          HTMLElement = "A"
	Abbr       HTMLElement = "ABBR"
	Address    HTMLElement = "ADDRESS"
	Area       HTMLElement = "AREA"
	Article    HTMLElement = "ARTICLE"
	Aside      HTMLElement = "ASIDE"
	Audio      HTMLElement = "AUDIO"
	B          HTMLElement = "B"
	Base       HTMLElement = "BASE"
	Bdi        HTMLElement = "BDI"
	Bdo        HTMLElement = "BDO"
	Blockquote HTMLElement = "BLOCKQUOTE"
	Body       HTMLElement = "BODY"
	Br         HTMLElement = "BR"
	Button     HTMLElement = "BUTTON"
	Canvas     HTMLElement = "CANVAS"
	Caption    HTMLElement = "CAPTION"
	Cite       HTMLElement = "CITE"
	Code       HTMLElement = "CODE"
	Col        HTMLElement = "COL"
	Colgroup   HTMLElement = "COLGROUP"
	Data       HTMLElement = "DATA"
	Datalist   HTMLElement = "DATALIST"
	Dd         HTMLElement = "DD"
	Del        HTMLElement = "DEL"
	Details    HTMLElement = "DETAILS"
	Dfn        HTMLElement = "DFN"
	Dialog     HTMLElement = "DIALOG"
	Div        HTMLElement = "DIV"
	Dl         HTMLElement = "DL"
	Dt         HTMLElement = "DT"
	Em         HTMLElement = "EM"
	Embed      HTMLElement = "EMBED"
	Fieldset   HTMLElement = "FIELDSET"
	Figcaption HTMLElement = "FIGCAPTION"
	Figure     HTMLElement = "FIGURE"
	Footer     HTMLElement = "FOOTER"
	Form       HTMLElement = "FORM"
	H1         HTMLElement = "H1"
	H2         HTMLElement = "H2"
	H3         HTMLElement = "H3"
	H4         HTMLElement = "H4"
	H5         HTMLElement = "H5"
	H6         HTMLElement = "H6"
	Head       HTMLElement = "HEAD"
	Header     HTMLElement = "HEADER"
	Hr         HTMLElement = "HR"
	Html       HTMLElement = "HTML"
	I          HTMLElement = "I"
	Iframe     HTMLElement = "IFRAME"
	Img        HTMLElement = "IMG"
	Input      HTMLElement = "INPUT"
	Ins        HTMLElement = "INS"
	Kbd        HTMLElement = "KBD"
	Label      HTMLElement = "LABEL"
	Legend     HTMLElement = "LEGEND"
	Li         HTMLElement = "LI"
	Link       HTMLElement = "LINK"
	Main       HTMLElement = "MAIN"
	MapElement HTMLElement = "MAP"
	Mark       HTMLElement = "MARK"
	Meta       HTMLElement = "META"
	Meter      HTMLElement = "METER"
	Nav        HTMLElement = "NAV"
	Noscript   HTMLElement = "NOSCRIPT"
	Object     HTMLElement = "OBJECT"
	Ol         HTMLElement = "OL"
	Optgroup   HTMLElement = "OPTGROUP"
	Option     HTMLElement = "OPTION"
	Output     HTMLElement = "OUTPUT"
	P          HTMLElement = "P"
	Param      HTMLElement = "PARAM"
	Picture    HTMLElement = "PICTURE"
	Pre        HTMLElement = "PRE"
	Progress   HTMLElement = "PROGRESS"
	Q          HTMLElement = "Q"
	Rp         HTMLElement = "RP"
	Rt         HTMLElement = "RT"
	Ruby       HTMLElement = "RUBY"
	S          HTMLElement = "S"
	Samp       HTMLElement = "SAMP"
	Script     HTMLElement = "SCRIPT"
	Section    HTMLElement = "SECTION"
	Select     HTMLElement = "SELECT"
	Small      HTMLElement = "SMALL"
	Source     HTMLElement = "SOURCE"
	Span       HTMLElement = "SPAN"
	Strong     HTMLElement = "STRONG"
	Style      HTMLElement = "STYLE"
	Sub        HTMLElement = "SUB"
	Summary    HTMLElement = "SUMMARY"
	Sup        HTMLElement = "SUP"
	Svg        HTMLElement = "SVG"
	Table      HTMLElement = "TABLE"
	Tbody      HTMLElement = "TBODY"
	Td         HTMLElement = "TD"
	Template   HTMLElement = "TEMPLATE"
	Textarea   HTMLElement = "TEXTAREA"
	Tfoot      HTMLElement = "TFOOT"
	Th         HTMLElement = "TH"
	Thead      HTMLElement = "THEAD"
	Time       HTMLElement = "TIME"
	Title      HTMLElement = "TITLE"
	Tr         HTMLElement = "TR"
	Track      HTMLElement = "TRACK"
	U          HTMLElement = "U"
	Ul         HTMLElement = "UL"
	Var        HTMLElement = "VAR"
	Video      HTMLElement = "VIDEO"
	Wbr        HTMLElement = "WBR"
	Text       HTMLElement = "#text"
)

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
	Pages           []*Page
	IsPublished     bool `db:"is_published"`
	Repository      string
	UserID          string    `db:"user_id"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
	LastPublishedAt time.Time `db:"last_published_at"`
}

type Page struct {
	ID           string
	Title        string
	Slug         string
	Dependencies []string
	Nodes        []*TreeNode // Should be binary
	Pages        []*Page
	PageID       string    `db:"page_id"`
	SiteID       string    `db:"site_id"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

type TreeNode struct {
	Tag         HTMLElement
	TextContent string
	Dependency  string
	ClassList   []string
	Children    []*TreeNode
}
