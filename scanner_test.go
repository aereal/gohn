package main

import (
	"strings"
	"testing"
)

type expectation struct {
	input   string
	token   Token
	literal string
}

func TestScanner_Scan(t *testing.T) {
	var expectations = []expectation{
		{input: ``, token: EOF},
		{input: "hello", token: PARAGRAPH, literal: "hello"},
		{input: "hello\nworld", token: PARAGRAPH, literal: "hello\nworld"},
	}

	for i, expect := range expectations {
		scanner := NewScanner(strings.NewReader(expect.input))
		token, literal := scanner.Scan()
		if expect.token != token {
			t.Errorf("%d %q token mismatched; expected=%q actual=%q <%q>", i, expect.input, expect.token, token, literal)
		} else if expect.literal != literal {
			t.Errorf("%d %q literal mismatched; expected=%q actual=%q", i, expect.input, expect.literal, literal)
		}
	}
}
