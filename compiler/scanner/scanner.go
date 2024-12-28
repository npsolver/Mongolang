package scanner

import (
	"errors"
	"log"
	"strings"
	"iter"

	"github.com/Npsolver/Mongolang/compiler/scanner/sdfa"
	"github.com/Npsolver/Mongolang/global"
)

var (
	errScanningError = errors.New("cannot scan query")
)

func Scan(s string) ([]*global.Token, error) {
	allTokens := []*global.Token{}
	sdfa := sdfa.NewSDFA("/Users/raiyanjamil/Raiyan/dev/projects/Mongolang/compiler/scanner/sdfa/scanner.dfa")
	traveler := sdfa.NewDFATraveler(Map(strings.Join(strings.Split(s, " "), ""), func(s string) global.TerminalSymbol { return global.NewTerminalSymbol(s) }))
	for {
		tk, err := traveler.NextToken()
		if err != nil {
			log.Println(errScanningError)
			return nil, err // need to fix error handling here
		}
		if tk == nil {
			break
		}
		allTokens = append(allTokens, tk)
	}
	return allTokens, nil
}
