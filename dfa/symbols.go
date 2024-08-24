package dfa

import (
	"bufio"
	"strings"
)

type Symbol string

func (d *DFA) scanSymbols(scanner *bufio.Scanner) {
	scanner.Scan()
	scanner.Scan()
	for scanner.Text() != STATES {
		symbols := strings.Split(strings.Trim(scanner.Text(), " "), " ")
		for _, c := range symbols {
			if len(c) == 3 {
				for i := c[0]; i != c[2]; i++ {
					d.allSymbols = append(d.allSymbols, Symbol(i))
				}
				d.allSymbols = append(d.allSymbols, Symbol(c[2]))
			} else {
				d.allSymbols = append(d.allSymbols, Symbol(c))
			}
		}
		scanner.Scan()
	}
}

func (s Symbol) append(t Symbol) Symbol {
	return s + t
}
