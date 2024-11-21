package scanner

import (
	"errors"
	"log"
	"strings"

	"github.com/Npsolver/Mongolang/internal/compiler/scanner/sdfa"
	"github.com/Npsolver/Mongolang/internal/global"
)

var (
	errScanningError = errors.New("cannot scan query")
)

func Scan(s string) ([]*global.Token, error) {
	allTokens := []*global.Token{}
	sdfa := sdfa.NewSDFA("/Users/raiyanjamil/Raiyan/dev/projects/Mongolang/internal/compiler/scanner/sdfa/scanner.dfa")
	traveler := sdfa.NewDFATraveler(strings.Join(strings.Split(s, " "), ""))
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
