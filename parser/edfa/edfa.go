package edfa

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/npsolver/Mongolang/global"
)

const (
	EMPTY = "EMPTY"
	END   = ".END"
)

type EDFA struct {
	terminatingSymbols map[string]struct{}
	StartingNode       *Node
	allNodes           []*Node
	itemToNode         map[string]*Node
	leftSymbolToItems  map[string][]string
	createNode         func(item *Item) *Node
}

func NewEDFA(filePath string, symList []*global.Symbol) *EDFA {
	fmt.Println("Reached here")
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fileScanner := bufio.NewScanner(file)

	edfa := &EDFA{}
	edfa.terminatingSymbols = make(map[string]struct{})
	edfa.itemToNode = make(map[string]*Node)
	edfa.leftSymbolToItems = make(map[string][]string)
	for _, sym := range symList {
		edfa.terminatingSymbols[sym.GetName()] = struct{}{}
	}

	fmt.Println("printing all terminaitng symbols")
	for item := range edfa.terminatingSymbols {
		fmt.Println(item)
	}

	edfa.parseCFS(fileScanner)

	return edfa
}

func (e *EDFA) isTerminating(symbol string) bool {
	_, exists := e.terminatingSymbols[symbol]
	return exists
}
