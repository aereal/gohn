package main

import (
	"io"
	"text/scanner"
)

var symbolTables = map[string]int{
	"-": UNORDERED_LIST_MARKER,
}

type Block interface{}

type UnorderedList struct {
	items []UnorderedListItem
}

type UnorderedListItem struct {
	text string
}

type Token struct {
	token   int
	literal string
}

type Lexer struct {
	scanner.Scanner
	result []Block
	err    *ParseError
}

type ParseError struct {
	Message string
	Line    int
	Column  int
}

func (e *ParseError) Error() string {
	return e.Message
}

func NewLexer(in io.Reader) *Lexer {
	l := new(Lexer)
	l.Init(in)
	l.Mode &^= scanner.ScanInts | scanner.ScanFloats | scanner.ScanStrings | scanner.ScanComments | scanner.SkipComments
	return l
}

func (l *Lexer) Lex(lval *yySymType) int {
	token := int(l.Scan())
	s := l.TokenText()
	if token == scanner.String || token == scanner.Ident {
		token = TEXT
	}
	if _, ok := symbolTables[s]; ok {
		token = symbolTables[s]
	}
	lval.token = Token{token: token, literal: s}
	return token
}

func (l *Lexer) Error(e string) {
	l.err = &ParseError{
		Message: e,
		Line:    l.Line,
		Column:  l.Column,
	}
}

func Parse(src io.Reader) ([]Block, error) {
	lex := NewLexer(src)
	if ok := yyParse(lex); ok == 0 {
		return lex.result, nil
	} else {
		return nil, lex.err
	}
}
