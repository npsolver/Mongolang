package edfa

import (
	"bufio"
	"strings"

	"github.com/npsolver/Mongolang/debug"
)

func nodeCreator() func(item *Item) *Node {
	id := 0
	return func(item *Item) *Node {
		id++

		newNode := CreateNode(id)
		newNode.AppendItem(item)

		return newNode
	}
}

func (e *EDFA) dfs(curr *Node) {

	curr.Visited = true

	startingItem := curr.Items[0].Format()
	debug.DebugPrint("Starting Item of node with id = %d\n", curr.ID)
	debug.DebugPrint(startingItem)
	leftSymbols := make(map[string]struct{})
	index := 0
	for index != len(curr.Items) {
		currItem := curr.Items[index]
		if len(currItem.rhsr) == 0 {
			curr.SetTerminating()
			debug.DebugPrint("Inside new node with id = %d\n", curr.ID)
			debug.DebugPrint("Printing items of a terminating node")
			if debug.DEBUG {
				for _, item := range curr.Items {
					item.Print()
				}
			}
			return
		}
		newLeft := currItem.rhsr[0]
		if _, exists := leftSymbols[newLeft]; !exists {
			// add rules that have a newLeft on the left of -> (use map2)
			for _, item := range e.leftSymbolToItems[newLeft] {
				createdItem := CreateItem(item)
				if createdItem.Format() != startingItem {
					curr.AppendItem(createdItem)
				}
			}
			leftSymbols[newLeft] = struct{}{}
		}
		index++
	}

	// currNode -> nextNode
	//       first of rhsr
	// create nextNode's first item
	// check if such a node exists using map1, create one if not
	// add that to the map

	// while transitioning, check if there is a node whose
	// first item is the transitioned one, if no, then
	// create new one and dfs into it (use map1)

	index = 0

	for index != len(curr.Items) {
		currItem := curr.Items[index]
		nextItem := currItem.Shift()
		if _, exists := curr.Bridges[currItem.rhsr[0]]; exists {
			curr.Bridges[currItem.rhsr[0]].AppendItem(nextItem)
		} else if _, exists := e.itemToNode[nextItem.Format()]; exists {
			curr.Bridges[currItem.rhsr[0]] = e.itemToNode[nextItem.Format()]
		} else {
			newNode := e.createNode(nextItem)
			e.AddToItemToNode(newNode)
			curr.Bridges[currItem.rhsr[0]] = e.itemToNode[nextItem.Format()]
		}
		index++
	}

	for _, node := range curr.Bridges {
		if !node.Visited {
			e.dfs(node)
		}
	}

	// for index != len(curr.Items) {
	// 	currItem := curr.Items[index]
	// 	shiftedItem := currItem.Shift()
	// 	if _, exists := curr.Bridges[currItem.rhsr[0]]; !exists {
	// 		newNode := e.createNode(shiftedItem)
	// 		curr.Bridges[currItem.rhsr[0]] = newNode
	// 	} else {
	// 		curr.Bridges[currItem.rhsr[0]].AppendItem(shiftedItem)
	// 	}
	// 	index++
	// }

	debug.DebugPrint("Inside new node with id = %d\n", curr.ID)
	debug.DebugPrint("Printing items")
	if debug.DEBUG {
		for _, item := range curr.Items {
			item.Print()
		}
	}
	debug.DebugPrint("Printing Bridges")
	for s, b := range curr.Bridges {
		debug.DebugPrint("%s --> %d\n", s, b.ID)
	}

}

func (e *EDFA) parseCFS(scanner *bufio.Scanner) {
	e.createNode = nodeCreator()
	scanner.Scan()
	scanner.Scan()
	startingCFS := scanner.Text()

	newItem := CreateItem(scanner.Text())
	e.StartingNode = e.createNode(newItem)
	e.allNodes = append(e.allNodes, e.StartingNode)

	// map1 from item.Format() to node | item is the first item of node's items
	// map2 from left token to list of items

	e.AddToItemToNode(e.StartingNode)
	e.AddToLeftSymbolToItems(startingCFS)

	scanner.Scan()
	for scanner.Text() != END {
		e.AddToLeftSymbolToItems(scanner.Text())
		scanner.Scan()
	}
	e.dfs(e.StartingNode)

}

func (e *EDFA) AddToItemToNode(node *Node) {
	e.itemToNode[node.Items[0].Format()] = node
}

func (e *EDFA) AddToLeftSymbolToItems(cfs string) {
	lhs := strings.Fields(cfs)[0]
	e.leftSymbolToItems[lhs] = append(e.leftSymbolToItems[lhs], cfs)
}
