package scanner

import (
	"errors"
	"log"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/npsolver/Mongolang/global"
	"github.com/npsolver/Mongolang/scanner/dfa"
)

var (
	errScanningError = errors.New("cannot scan query")
)

func Scan(s string) ([]*global.Symbol, []string, error) {
	allSymbols := []*global.Symbol{}

	_, file, _, ok := runtime.Caller(0)
	if !ok {
		panic("unable to get caller info")
	}

	dir := filepath.Dir(file)
	fullPath := filepath.Join(dir, "scanner.dfa")

	scanDFA := dfa.NewDFA(fullPath)
	terminatingStates := scanDFA.GetTerminatingStates()

	traveler := scanDFA.NewDFATraveler(strings.Join(strings.Split(s, " "), ""))
	for {
		tk, err := traveler.NextSymbol()
		if err != nil {
			log.Println(errScanningError)
			return nil, nil, err // need to fix error handling here
		}
		if tk == nil {
			break
		}
		allSymbols = append(allSymbols, tk)
	}
	return allSymbols, terminatingStates, nil
}
