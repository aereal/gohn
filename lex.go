package main

import (
	"io"
	"text/scanner"
	"unicode"
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

type Line struct {
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

func isIdent(ch rune, size int) bool {
	return unicode.IsGraphic(ch)
}

func NewLexer(in io.Reader) *Lexer {
	l := new(Lexer)
	l.Init(in)
	l.Mode &^= scanner.ScanInts | scanner.ScanFloats | scanner.ScanStrings | scanner.ScanComments | scanner.SkipComments
	l.IsIdentRune = isIdent
	return l
}

func (l *Lexer) skipBlank() {
	for unicode.IsSpace(l.Peek()) {
		l.Next()
	}
}

func (l *Lexer) Lex(lval *yySymType) int {
	l.skipBlank()
	ch := l.Peek()
	if _, ok := symbolTables[string(ch)]; ok {
		_ = l.Next()
		s := string(ch)
		token := symbolTables[s]
		lval.token = Token{token: token, literal: s}
		return token
	} else {
		token := int(l.Scan())
		s := l.TokenText()
		if token == scanner.String || token == scanner.Ident {
			token = TEXT
		}
		lval.token = Token{token: token, literal: s}
		return token
	}
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
