package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

type expectation struct {
	input  string
	result []Block
}

func TestParser_Parse(t *testing.T) {
	var expectations = []expectation{
		{
			input: "- a\n",
			result: []Block{
				UnorderedList{
					items: []UnorderedListItem{
						UnorderedListItem{text: "a"},
					},
				},
			},
		},
		{
			input: "a\nb\n",
			result: []Block{
				Line{
					inlines: []Inline{
						InlineText{literal: "a"},
					},
				},
				Line{
					inlines: []Inline{
						InlineText{literal: "b"},
					},
				},
			},
		},
		{
			input: "姉\n弟\n",
			result: []Block{
				Line{
					inlines: []Inline{
						InlineText{literal: "姉"},
					},
				},
				Line{
					inlines: []Inline{
						InlineText{literal: "弟"},
					},
				},
			},
		},
		{
			input: "[http://example.com/]\n弟\n",
			result: []Block{
				Line{
					inlines: []Inline{
						InlineHttp{url: "http://example.com/"},
					},
				},
				Line{
					inlines: []Inline{
						InlineText{literal: "弟"},
					},
				},
			},
		},
		{
			input: "- a\n- b\na\nb\n",
			result: []Block{
				UnorderedList{
					items: []UnorderedListItem{
						UnorderedListItem{text: "a"},
						UnorderedListItem{text: "b"},
					},
				},
				Line{
					inlines: []Inline{
						InlineText{literal: "a"},
					},
				},
				Line{
					inlines: []Inline{
						InlineText{literal: "b"},
					},
				},
			},
		},
	}

	for i, expect := range expectations {
		actual, err := Parse(strings.NewReader(expect.input))
		if err != nil {
			t.Errorf("! #%d Failed to parse: %#v", i, err)
			continue
		}
		ok, msg := matchResult(expect.result, actual)
		if !ok {
			t.Error(msg)
			continue
		}
	}
}

func matchResult(expected []Block, actual []Block) (ok bool, msg string) {
	if len(expected) != len(actual) {
		ok = false
		msg = fmt.Sprintf("Number of blocks mismatched; %d expected, but %d got; <%#v>", len(expected), len(actual), actual)
		return
	}
	if !reflect.DeepEqual(expected, actual) {
		ok = false
		msg = fmt.Sprintf("Parsed result mismatched\nExpected:\n%#v\nActual:\n%#v\n", expected, actual)
		return
	}
	ok = true
	return
}
