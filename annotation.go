package main

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
	Cite    string
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
	Url string
}
