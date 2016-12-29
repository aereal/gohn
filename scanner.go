package main

import (
	"bufio"
	"bytes"
	"io"
	"unicode"
	"unicode/utf8"
)

// from https://github.com/benbjohnson/sql-parser/blob/master/scanner.go

type Token int

const (
	ILLEGAL Token = iota
	EOF
	WS
	PARAGRAPH
	UNORDERED_LIST
	ORDERED_LIST
)

type NodeContext int

const (
	cRoot NodeContext = iota
	cParagraph
	cUnorderedList
	cOrderedList
)

const eof = rune(0)
const tUnorderedList rune = '-'
const tOrderedList rune = '+'

type Scanner struct {
	input   *bufio.Reader
	context NodeContext
}

func NewScanner(i io.Reader) *Scanner {
	return &Scanner{input: bufio.NewReader(i)}
}

func (s *Scanner) Scan() (token Token, literal string) {
	c := s.read()
	switch c {
	case eof:
		return EOF, ""
	case tUnorderedList:
		s.context = cUnorderedList
		return UNORDERED_LIST, string(tUnorderedList)
	case tOrderedList:
		s.context = cOrderedList
		return ORDERED_LIST, string(tOrderedList)
	}
	if unicode.IsLetter(c) {
		if s.context == cUnorderedList || s.context == cOrderedList {
			s.unread()
			return s.scanLine()
		} else {
			s.unread()
			return s.scanParagraph()
		}
	} else if unicode.IsSpace(c) {
		s.unread()
		return s.scanWhitespace()
	}
	return ILLEGAL, string(c)
}

func (s *Scanner) scanParagraph() (token Token, literal string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if c := s.read(); c == eof {
			break
			// } else if !unicode.IsLetter(c) {
			// 	s.unread()
			// 	break
		} else {
			_, _ = buf.WriteRune(c)
		}
	}

	return PARAGRAPH, buf.String()
}

func (s *Scanner) scanWhitespace() (token Token, literal string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if c := s.read(); c == eof {
			break
		} else if !unicode.IsSpace(c) {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(c)
		}
	}

	return WS, buf.String()
}

func (s *Scanner) scanLine() (token Token, literal string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if c := s.read(); c == eof {
			break
		} else if c == '\n' {
			nextBytes, _ := s.input.Peek(1)
			if r, _ := utf8.DecodeRune(nextBytes); r != tUnorderedList || r != tOrderedList {
				s.context = cRoot
			}
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(c)
		}
	}

	return PARAGRAPH, buf.String()
}

func (s *Scanner) read() rune {
	c, _, err := s.input.ReadRune()
	if err != nil {
		return eof
	}
	return c
}

func (s *Scanner) unread() {
	_ = s.input.UnreadRune()
}
