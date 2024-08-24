package scanner

import (
	"errors"
	"log"
	"strings"

	"github.com/Npsolver/Mongolang/dfa"
)

var (
	errScanningError = errors.New("cannot scan query")
)

func Scan(s string) ([]*dfa.Token, error) {
	allTokens := []*dfa.Token{}
	scanDFA := dfa.NewDFA("/Users/raiyanjamil/Raiyan/dev/projects/Mongolang/scanner/scanner.dfa")
	traveler := scanDFA.NewDFATraveler(strings.Join(strings.Split(s, " "), ""))
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
