package parser

import (
	"fmt"
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

	parseEDFA := edfa.NewEDFA(fullPath, symList)

	// Make a set of terminating symbols
	// done in edfa.terminaitngSymbols

	// Parse the symbols
	edfaStack := []*edfa.Node{parseEDFA.StartingNode}
	symbolStack := []*global.Symbol{}
	treeNodeStack := []*TreeNode{}

	index := 0
	for (index != len(symList)) || (len(edfaStack) > 1) {

		fmt.Printf("In loop with index = %d\n", index)
		for _, e := range edfaStack {
			fmt.Printf("%d ", e.ID)
		}
		fmt.Printf("\n")
		for _, s := range symbolStack {
			fmt.Printf("%s %s\n", s.GetName(), s.GetValue())
		}
		fmt.Printf("\n")

		// if you can reduce, then reduce, otherwise add a symbol

		// Transition, shift or reduce

		// Reduce if current is terminating
		// Transition right after reducing only if

		// Reduce
		if edfaStack[len(edfaStack)-1].IsTerminating() { // handle empty case~~~
			fmt.Println("Reducing using item")
			edfaStack[len(edfaStack)-1].Items[0].Print()

			rhsLen := edfaStack[len(edfaStack)-1].Items[0].GetRhslLen()
			newSymbol := global.NewNonTerminatingSymbol(edfaStack[len(edfaStack)-1].Items[0].GetLhs())
			newTreeNode := &TreeNode{newSymbol, []*TreeNode{}}
			for i := 1; i <= rhsLen; i++ {
				edfaStack = edfaStack[:len(edfaStack)-1]
				symbolStack = symbolStack[:len(symbolStack)-1]
				lastTreeNode := treeNodeStack[len(treeNodeStack)-1]
				treeNodeStack = treeNodeStack[:len(treeNodeStack)-1]
				newTreeNode.AppendChild(lastTreeNode)
			}
			if newSymbol.GetName() != edfaStack[len(edfaStack)-1].Items[0].GetLhs() {
				nextNode := edfaStack[len(edfaStack)-1].Bridges[newSymbol.GetName()]
				edfaStack = append(edfaStack, nextNode)
				symbolStack = append(symbolStack, newSymbol)
			}
			treeNodeStack = append(treeNodeStack, newTreeNode)
			// symbolStack[len(symbolStack)-1] = newSymbol
			// lastTreeNode := treeNodeStack[len(treeNodeStack)-1]
			// newTreeNode.AppendChild(lastTreeNode)
			// treeNodeStack[len(treeNodeStack)-1] = newTreeNode
		} else { // Add
			currSymbol := symList[index]
			fmt.Printf("Adding symbol %s with value %s to stack\n", currSymbol.GetName(), currSymbol.GetValue())
			// add to stacks
			nextNode, exists := edfaStack[len(edfaStack)-1].Bridges[currSymbol.GetName()]
			if !exists {
				nextNode = edfaStack[len(edfaStack)-1].Bridges[edfa.EMPTY]
				treeNodeStack = append(treeNodeStack, &TreeNode{global.NewTerminatingSymbol(edfa.EMPTY, edfa.EMPTY), nil})
			} else {
				treeNodeStack = append(treeNodeStack, &TreeNode{currSymbol, nil})
				index++
			}
			edfaStack = append(edfaStack, nextNode)
			symbolStack = append(symbolStack, currSymbol)
		}

		// use bridge to transition to another node
		// find next node
		// add currSymbol to symbol stack

		// if new node is accepting, reduce
		// remove last few items? from symbolstack (use current stack's item to find it)
		// and add the reduced symbol to symbolstack

		// otherwise, continue
	}

	fmt.Println("Finished tree:")
	treeNodeStack[0].Print()

	return nil
}
