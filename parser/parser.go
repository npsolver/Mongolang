package parser

import (
	"path/filepath"
	"runtime"

	"github.com/npsolver/Mongolang/global"
	"github.com/npsolver/Mongolang/parser/edfa"
)

// func Parse(symbols []*dfa.Symbol) {

// 	terminalSymbols := []string{}
// 	for _, t := range symbols {
// 		terminalSymbols = append(terminalSymbols, t.GetName())
// 	}

// 	fmt.Println(terminalSymbols)
// }

func Parse(symbols []*global.Symbol) *TreeNode {
	BOF := global.NewTerminatingSymbol("BOF", "BOF")
	EOF := global.NewTerminatingSymbol("EOF", "EOF")
	symList := append([]*global.Symbol{BOF}, symbols...)
	symList = append(symList, EOF)

	_, file, _, ok := runtime.Caller(0)
	if !ok {
		panic("unable to get caller info")
	}

	dir := filepath.Dir(file)
	fullPath := filepath.Join(dir, "parser.cfg")

	edfa.NewEDFA(fullPath, symList)

	// Make a set of terminating symbols

	// Parse the

	return nil
}
