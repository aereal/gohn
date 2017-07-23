package main

import (
	"io"
	"text/scanner"
	"unicode"
)

const NEW_LINE = '\n'
const EOF = -1

var symbolTables = map[string]int{
	"-":              UNORDERED_LIST_MARKER,
	"+":              ORDERED_LIST_MARKER,
	string(NEW_LINE): CR,
	"[":              LBRACKET,
	"]":              RBRACKET,
	"<":              LT,
	">":              GT,
	"*":              HEADING_MARKER,
	":":              COLON,
}

type Token struct {
	token   int
	literal string
}

func (t Token) Name() string {
	switch t.token {
	case TEXT:
		return "TEXT"
	case UNORDERED_LIST_MARKER:
		return "UNORDERED_LIST_MARKER"
	case ORDERED_LIST_MARKER:
		return "ORDERED_LIST_MARKER"
	case CR:
		return "CR"
	case LBRACKET:
		return "LBRACKET"
	case RBRACKET:
		return "RBRACKET"
	case LT:
		return "LT"
	case GT:
		return "GT"
	case HEADING_MARKER:
		return "HEADING_MARKER"
	case EOF:
		return "EOF"
	case COLON:
		return "COLON"
	default:
		return "UNKNOWN"
	}
}

type Lexer struct {
	scanner.Scanner
	result    []Block
	err       *ParseError
	inHttp    bool
	seenColon bool
}

type ParseError struct {
	Message string
	Line    int
	Column  int
}

func (e *ParseError) Error() string {
	return e.Message
}

func isWhitespace(ch rune) bool {
	return unicode.IsSpace(ch) && ch != rune(NEW_LINE)
}

func NewLexer(in io.Reader) *Lexer {
	l := new(Lexer)
	l.Init(in)
	l.Mode &^= scanner.ScanInts | scanner.ScanFloats | scanner.ScanStrings | scanner.ScanComments | scanner.SkipComments
	l.IsIdentRune = l.isIdent
	l.Whitespace = 1<<' ' | 1<<'\t'
	l.inHttp = false
	l.seenColon = false
	return l
}

func (l *Lexer) isIdent(ch rune, size int) bool {
	return unicode.IsGraphic(ch) && !l.isReserved(ch) && !unicode.IsSpace(ch)
}

func (l *Lexer) isReserved(ch rune) bool {
	token, ok := symbolTables[string(ch)]
	if ok {
		switch token {
		case COLON:
			if !l.seenColon {
				l.seenColon = true
				return false // maybe part of URL
			} else {
				return true
			}
		default:
			return true
		}
	} else {
		return false
	}
}

func (l *Lexer) skipBlank() {
	for isWhitespace(l.Peek()) {
		l.Next()
	}
}

func (l *Lexer) Lex(lval *yySymType) int {
	l.skipBlank()
	ch := l.Peek()
	if l.isReserved(ch) {
		s := string(ch)
		token := symbolTables[s]
		if token == LBRACKET {
			l.inHttp = true
			l.seenColon = false // reset
		} else if token == RBRACKET {
			l.inHttp = false
			l.seenColon = false // reset
		}
		_ = l.Next()
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
