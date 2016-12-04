package script

import (
	"bytes"
	"io"
	"io/ioutil"
)

type Script struct {
	content []byte
	pos     int
}

func New(input io.Reader) (*Script, error) {
	content, err := ioutil.ReadAll(input)
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

	switch token {
	case '[':
		if current == 0x00 {
			s.pos = s.findMatchingClosingBracket(s.pos) + 1
		} else {
			s.pos++
		}
		return s.NextInstruction(current)

	case ']':
		if current != 0x00 {
			s.pos = s.findMatchingOpeningBracket(s.pos) + 1
		} else {
			s.pos++
		}
		return s.NextInstruction(current)

	default:
		s.pos++
		return token
	}
}

// finds the matching closing bracket to the right
func (s *Script) findMatchingClosingBracket(pos int) int {
	nextClosing := pos + 1 + bytes.IndexByte(s.content[pos+1:], ']')
	nextOpening := pos + 1 + bytes.IndexByte(s.content[pos+1:], '[')

	if nextClosing < nextOpening {
		return nextClosing
	}

	return s.findMatchingClosingBracket(s.findMatchingClosingBracket(nextOpening))
}

// finds the matching opening bracket to the left
func (s *Script) findMatchingOpeningBracket(pos int) int {
	nextOpening := bytes.LastIndexByte(s.content[:pos], '[')
	nextClosing := bytes.LastIndexByte(s.content[:pos], ']')

	if nextOpening > nextClosing {
		return nextOpening
	}

	return s.findMatchingOpeningBracket(s.findMatchingOpeningBracket(nextClosing))
}
