package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/sauerbraten/hirnfick/internal/machine"
	"github.com/sauerbraten/hirnfick/internal/script"
)

func main() {
	// init VM
	m := machine.New()

	// setup buffered stdin reader
	input := bufio.NewReader(os.Stdin)

	// open file
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	// read script
	script, err := script.New(file)
	if err != nil {
		log.Fatalln(err)
	}

	// read & execute tokens
	for script.HasRemaining() {
		instr := script.NextInstruction(m.GetByte())
		switch instr {
		case '>':
			m.IncrAddr()
		case '<':
			m.DecrAddr()
		case '+':
			m.IncrByte()
		case '-':
			m.DecrByte()
		case '.':
			fmt.Printf("%c", m.GetByte())
		case ',':
			b, err := input.ReadByte()
			if err != nil {
				log.Fatalln(err)
			}
			m.PutByte(b)
		}
	}
}
