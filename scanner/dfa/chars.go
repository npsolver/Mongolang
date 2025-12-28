package dfa

import (
	"bufio"
	"strings"
)

type Char string

func (d *DFA) scanChars(scanner *bufio.Scanner) {
	scanner.Scan()
	scanner.Scan()
	for scanner.Text() != STATES {
		chars := strings.Split(strings.Trim(scanner.Text(), " "), " ")
		for _, c := range chars {
			if len(c) == 3 {
				for i := c[0]; i != c[2]; i++ {
					d.allChars = append(d.allChars, Char(i))
				}
				d.allChars = append(d.allChars, Char(c[2]))
			} else {
				d.allChars = append(d.allChars, Char(c))
			}
		}
		scanner.Scan()
	}
}

func (s Char) append(t Char) Char {
	return s + t
}
