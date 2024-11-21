package sdfa

import (
	"strings"

	"github.com/Npsolver/Mongolang/internal/global"
	"github.com/Npsolver/Mongolang/internal/lib/dfa"
)

type SDFA struct {
	dfa.DFA[string, global.TerminalSymbol]
}

func NewSDFA(filePath string) *SDFA {
	ch := dfa.NewFileReader(filePath)
	sdfa := &SDFA{}
	sdfa.ScanSymbols(ch)
	return sdfa
}

func (s *SDFA) ScanSymbols(ch <-chan string) {
	<-ch
	text := (<-ch)
	for text != global.STATES {
		symbols := strings.Split(strings.Trim(text, " "), " ")
		for _, c := range symbols {
			if len(c) == 3 {
				for i := c[0]; i != c[2]; i++ {
					s.AddSymbol(global.NewTerminalSymbol(string(i)))
				}
				s.AddSymbol(global.NewTerminalSymbol(string(c[2])))
			} else {
				s.AddSymbol(global.NewTerminalSymbol(c))
			}
		}
		text = <-ch
	}
}
