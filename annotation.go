package main

type Block interface{}

type Inline interface{}

type UnorderedList struct {
	Items []UnorderedListItem
}

type UnorderedListItem struct {
	Inlines []Inline
}

type Line struct {
	Inlines []Inline
}

type Quotation struct {
	Cite    string
	Content []Block
}

type InlineText struct {
	Literal string
}

type InlineHttp struct {
	Url string
}
