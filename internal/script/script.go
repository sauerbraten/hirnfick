package script

import (
	"io"
	"io/ioutil"
)

type Script struct {
	content []byte
	pos     int
}

func New(input io.Reader) (*Script, error) {
	cleanInput := &Cleaner{input}
	content, err := ioutil.ReadAll(cleanInput)
	if err != nil {
		return nil, err
	}

	return &Script{content: content}, nil
}

func (s *Script) HasRemaining() bool {
	return s.pos < len(s.content)
}

func (s *Script) NextInstruction(current byte) byte {
	token := s.content[s.pos]

	if token == '[' && current == 0x00 {
		s.jumpToClosingBracket()
	}

	if token == ']' && current != 0x00 {
		s.jumpToOpeningBracket()
	}

	s.pos++
	return token
}

func (s *Script) jumpToClosingBracket() {
	for depth := 1; depth > 0; {
		s.pos++
		switch s.content[s.pos] {
		case '[':
			depth++
		case ']':
			depth--
		}
	}
}

func (s *Script) jumpToOpeningBracket() {
	for depth := 1; depth > 0; {
		s.pos--
		switch s.content[s.pos] {
		case ']':
			depth++
		case '[':
			depth--
		}
	}
}
