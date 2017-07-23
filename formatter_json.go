package main

import (
	"encoding/json"
)

func (ul UnorderedList) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Name  string
		Items []UnorderedListItem
	}{
		Name:  "UnorderedList",
		Items: ul.Items,
	})
}

func (uli UnorderedListItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Name    string
		Inlines []Inline
		Depth   int
	}{
		Name:    "UnorderedListItem",
		Inlines: uli.Inlines,
		Depth:   uli.Depth,
	})
}

func (l Line) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Name    string
		Inlines []Inline
	}{
		Name:    "Line",
		Inlines: l.Inlines,
	})
}

func (q Quotation) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Name    string
		Content []Block
	}{
		Name:    "Quotation",
		Content: q.Content,
	})
}

func (h Heading) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Name    string
		Level   int
		Content []Inline
	}{
		Name:    "Heading",
		Level:   h.Level,
		Content: h.Content,
	})
}

func (it InlineText) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Name    string
		Literal string
	}{
		Name:    "InlineText",
		Literal: it.Literal,
	})
}

func (ih InlineHttp) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Name    string
		Url     string
		Options HttpOptions
	}{
		Name:    "InlineHttp",
		Url:     ih.Url,
		Options: ih.Options,
	})
}
