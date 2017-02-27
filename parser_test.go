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
					Items: []UnorderedListItem{
						UnorderedListItem{
							Inlines: []Inline{
								InlineText{Literal: "a"},
							},
						},
					},
				},
			},
		},
		{
			input: "a\nb\n",
			result: []Block{
				Line{
					Inlines: []Inline{
						InlineText{Literal: "a"},
					},
				},
				Line{
					Inlines: []Inline{
						InlineText{Literal: "b"},
					},
				},
			},
		},
		{
			input: "姉\n弟\n",
			result: []Block{
				Line{
					Inlines: []Inline{
						InlineText{Literal: "姉"},
					},
				},
				Line{
					Inlines: []Inline{
						InlineText{Literal: "弟"},
					},
				},
			},
		},
		{
			input: "[http://example.com/]\n弟\n",
			result: []Block{
				Line{
					Inlines: []Inline{
						InlineHttp{Url: "http://example.com/"},
					},
				},
				Line{
					Inlines: []Inline{
						InlineText{Literal: "弟"},
					},
				},
			},
		},
		{
			input: "- a\n- b\na\nb\n",
			result: []Block{
				UnorderedList{
					Items: []UnorderedListItem{
						UnorderedListItem{
							Inlines: []Inline{
								InlineText{Literal: "a"},
							},
						},
						UnorderedListItem{
							Inlines: []Inline{
								InlineText{Literal: "b"},
							},
						},
					},
				},
				Line{
					Inlines: []Inline{
						InlineText{Literal: "a"},
					},
				},
				Line{
					Inlines: []Inline{
						InlineText{Literal: "b"},
					},
				},
			},
		},
		{
			input: "a\n\nb\n",
			result: []Block{
				Line{
					Inlines: []Inline{
						InlineText{Literal: "a"},
					},
				},
				Line{
					Inlines: []Inline{},
				},
				Line{
					Inlines: []Inline{
						InlineText{Literal: "b"},
					},
				},
			},
		},
		{
			input: "- [http://example.com/]\n",
			result: []Block{
				UnorderedList{
					Items: []UnorderedListItem{
						UnorderedListItem{
							Inlines: []Inline{
								InlineHttp{Url: "http://example.com/"},
							},
						},
					},
				},
			},
		},
		{
			input: ">>\na\n- a\n- b\n<<\n",
			result: []Block{
				Quotation{
					Content: []Block{
						Line{
							Inlines: []Inline{
								InlineText{Literal: "a"},
							},
						},
						UnorderedList{
							Items: []UnorderedListItem{
								UnorderedListItem{
									Inlines: []Inline{
										InlineText{Literal: "a"},
									},
								},
								UnorderedListItem{
									Inlines: []Inline{
										InlineText{Literal: "b"},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			input: ">http://example.com/>\na\n- a\n- b\n<<\n",
			result: []Block{
				Quotation{
					Cite: "http://example.com/",
					Content: []Block{
						Line{
							Inlines: []Inline{
								InlineText{Literal: "a"},
							},
						},
						UnorderedList{
							Items: []UnorderedListItem{
								UnorderedListItem{
									Inlines: []Inline{
										InlineText{Literal: "a"},
									},
								},
								UnorderedListItem{
									Inlines: []Inline{
										InlineText{Literal: "b"},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			input: "+ a\n",
			result: []Block{
				OrderedList{
					Items: []OrderedListItem{
						OrderedListItem{
							Inlines: []Inline{
								InlineText{Literal: "a"},
							},
						},
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
