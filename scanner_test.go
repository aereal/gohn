package main

import (
	"fmt"
	"strings"
	"testing"
)

type result struct {
	token   Token
	literal string
}

type expectation struct {
	input   string
	results []result
}

func TestScanner_Scan(t *testing.T) {
	var expectations = []expectation{
		{
			input: ``,
			results: []result{
				{token: EOF},
			},
		},
		{
			input: "hello",
			results: []result{
				{token: PARAGRAPH, literal: "hello"},
				{token: EOF},
			},
		},
		{
			input: "hello\nworld",
			results: []result{
				{token: PARAGRAPH, literal: "hello\nworld"},
				{token: EOF},
			},
		},
		{
			input: "- perl\n- ruby",
			results: []result{
				{token: UNORDERED_LIST, literal: "-"},
				{token: WS, literal: " "},
				{token: PARAGRAPH, literal: "perl"},
				{token: WS, literal: "\n"},
				{token: UNORDERED_LIST, literal: "-"},
				{token: WS, literal: " "},
				{token: PARAGRAPH, literal: "ruby"},
				{token: EOF},
			},
		},
	}

	for _, expect := range expectations {
		scanner := NewScanner(strings.NewReader(expect.input))
		results := slurpTokens(scanner)
		ok, errorMsg := matchTokens(expect.results, results)
		if !ok {
			t.Error(errorMsg)
		}
	}
}

func slurpTokens(s *Scanner) (results []result) {
	for {
		token, literal := s.Scan()
		result := result{token: token, literal: literal}
		results = append(results, result)
		if token == EOF {
			break
		}
	}
	return
}

func matchTokens(expected []result, actual []result) (ok bool, msg string) {
	if len(expected) != len(actual) {
		ok = false
		msg = fmt.Sprintf("Number of tokens mismatched; expected=%d actual=%d <%q>", len(expected), len(actual), actual)
		return
	}
	for i, _ := range expected {
		if expected[i].token != actual[i].token {
			ok = false
			msg = fmt.Sprintf("Token mismatched; expected=%q actual=%q <%q>", expected[i].token, actual[i].token, actual[i].literal)
			return
		}
		if expected[i].literal != actual[i].literal {
			ok = false
			msg = fmt.Sprintf("Literal mismatched; expected=%q actual=%q", expected[i].literal, actual[i].literal)
			return
		}
	}
	ok = true
	return
}
