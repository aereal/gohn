package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

type expectation struct {
	description string
	input       string
	result      []Block
}

func TestParser_Parse(t *testing.T) {
	var expectations = []expectation{
		{
			description: "Simple unordered list",
			input:       "- a\n",
			result: []Block{
				UnorderedList{
					Items: []UnorderedListItem{
						UnorderedListItem{
							Depth: 1,
							Inlines: []Inline{
								InlineText{Literal: "a"},
							},
						},
					},
				},
			},
		},
		{
			description: "Lines",
			input:       "a\nb\n",
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
			description: "multibytes",
			input:       "姉\n弟\n",
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
			description: "Lines with HTTP annotation",
			input:       "[http://example.com/]\n弟\n",
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
			description: "List and Lines",
			input:       "- a\n- b\na\nb\n",
			result: []Block{
				UnorderedList{
					Items: []UnorderedListItem{
						UnorderedListItem{
							Depth: 1,
							Inlines: []Inline{
								InlineText{Literal: "a"},
							},
						},
						UnorderedListItem{
							Depth: 1,
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
			description: "Lines with empty line",
			input:       "a\n\nb\n",
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
			description: "List with http annotation",
			input:       "- [http://example.com/]\n",
			result: []Block{
				UnorderedList{
					Items: []UnorderedListItem{
						UnorderedListItem{
							Depth: 1,
							Inlines: []Inline{
								InlineHttp{Url: "http://example.com/"},
							},
						},
					},
				},
			},
		},
		{
			description: "quotation",
			input:       ">>\na\n- a\n- b\n<<\n",
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
									Depth: 1,
									Inlines: []Inline{
										InlineText{Literal: "a"},
									},
								},
								UnorderedListItem{
									Depth: 1,
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
			description: "quotation with cite",
			input:       ">http://example.com/>\na\n- a\n- b\n<<\n",
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
									Depth: 1,
									Inlines: []Inline{
										InlineText{Literal: "a"},
									},
								},
								UnorderedListItem{
									Depth: 1,
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
			description: "ordered list",
			input:       "+ a\n",
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
		{
			description: "ordered list and unordered list",
			input:       "+ a\n- b\n",
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
				UnorderedList{
					Items: []UnorderedListItem{
						UnorderedListItem{
							Depth: 1,
							Inlines: []Inline{
								InlineText{Literal: "b"},
							},
						},
					},
				},
			},
		},
		{
			description: "nested list",
			input:       "- a\n-- b\n",
			result: []Block{
				UnorderedList{
					Items: []UnorderedListItem{
						UnorderedListItem{
							Depth: 1,
							Inlines: []Inline{
								InlineText{Literal: "a"},
							},
						},
						UnorderedListItem{
							Depth: 2,
							Inlines: []Inline{
								InlineText{Literal: "b"},
							},
						},
					},
				},
			},
		},
		{
			description: "heading",
			input:       "* a\n",
			result: []Block{
				Heading{
					Level: 1,
					Content: []Inline{
						InlineText{Literal: "a"},
					},
				},
			},
		},
		{
			description: "heading with different level",
			input:       "* a\n** b\n",
			result: []Block{
				Heading{
					Level: 1,
					Content: []Inline{
						InlineText{Literal: "a"},
					},
				},
				Heading{
					Level: 2,
					Content: []Inline{
						InlineText{Literal: "b"},
					},
				},
			},
		},
	}

	for i, expect := range expectations {
		actual, err := Parse(strings.NewReader(expect.input))
		if err != nil {
			t.Errorf("! #%d %v: Failed to parse: %#v", i, expect.description, err)
			continue
		}
		ok, msg := matchResult(expect.result, actual)
		if !ok {
			t.Errorf("! #%d %v: %v", i, expect.description, msg)
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
