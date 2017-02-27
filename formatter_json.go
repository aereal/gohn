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
		Name string
		Text string
	}{
		Name: "UnorderedListItem",
		Text: uli.Text,
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
		Name string
		Url  string
	}{
		Name: "InlineHttp",
		Url:  ih.Url,
	})
}
