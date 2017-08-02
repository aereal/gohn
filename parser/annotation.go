package parser

type Document struct {
	Nodes []Block
}

type Block interface{}

type Inline interface{}

type UnorderedList struct {
	Items []UnorderedListItem
}

type UnorderedListItem struct {
	Depth   int
	Inlines []Inline
}

type OrderedList struct {
	Items []OrderedListItem
}

type OrderedListItem struct {
	Inlines []Inline
}

type Line struct {
	Inlines []Inline
}

type Quotation struct {
	Cite    Reference
	Content []Block
}

type Heading struct {
	Level   int
	Content []Inline
}

type InlineText struct {
	Literal string
}

type InlineHttp struct {
	Reference Reference
}

type Reference struct {
	Url     string
	Options HttpOptions
}

type HttpOptions []string
