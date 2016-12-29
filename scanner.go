package main

import (
	"bufio"
	"bytes"
	"io"
	"unicode"
)

// from https://github.com/benbjohnson/sql-parser/blob/master/scanner.go

type Token int

const (
	ILLEGAL Token = iota
	EOF
	PARAGRAPH
)

const eof = rune(0)

type Scanner struct {
	input *bufio.Reader
}

func NewScanner(i io.Reader) *Scanner {
	return &Scanner{input: bufio.NewReader(i)}
}

func (s *Scanner) Scan() (token Token, literal string) {
	c := s.read()
	switch c {
	case eof:
		return EOF, ""
	}
	if unicode.IsLetter(c) {
		s.unread()
		return s.scanParagraph()
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
