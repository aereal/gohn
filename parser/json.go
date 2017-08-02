package parser

import (
	"encoding/json"
)

func (doc *Document) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Nodes []Block `json:"nodes"`
	}{
		Nodes: doc.Nodes,
	})
}

func (ul UnorderedList) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Kind  string              `json:"kind"`
		Items []UnorderedListItem `json:"items"`
	}{
		Kind:  "UNORDERED_LIST",
		Items: ul.Items,
	})
}

func (uli UnorderedListItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Kind  string   `json:"kind"`
		Nodes []Inline `json:"nodes"`
		Depth int      `json:"depth"`
	}{
		Kind:  "UNORDERED_LIST_ITEM",
		Nodes: uli.Inlines,
		Depth: uli.Depth,
	})
}

func (l Line) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Kind  string   `json:"kind"`
		Nodes []Inline `json:"nodes"`
	}{
		Kind:  "LINE",
		Nodes: l.Inlines,
	})
}

func (q Quotation) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Kind  string    `json:"kind"`
		Nodes []Block   `json:"nodes"`
		Cite  Reference `json:"cite"`
	}{
		Kind:  "QUOTATION",
		Nodes: q.Content,
		Cite:  q.Cite,
	})
}

func (h Heading) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Kind  string   `json:"kind"`
		Level int      `json:"level"`
		Nodes []Inline `json:"nodes"`
	}{
		Kind:  "HEADING",
		Level: h.Level,
		Nodes: h.Content,
	})
}

func (it InlineText) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Kind  string `json:"kind"`
		Value string `json:"value"`
	}{
		Kind:  "TEXT",
		Value: it.Literal,
	})
}

func (ih InlineHttp) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Kind      string    `json:"kind"`
		Reference Reference `json:"reference"`
	}{
		Kind:      "HTTP",
		Reference: ih.Reference,
	})
}

func (ref Reference) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Url     string      `json:"url"`
		Options HttpOptions `json:"options"`
	}{
		Url:     ref.Url,
		Options: ref.Options,
	})
}
