package edfa

import (
	"bufio"
	"fmt"
	"strings"
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

	fmt.Println("Inside new node")

	leftSymbols := make(map[string]struct{})
	index := 0
	for index != len(curr.items) {
		currItem := curr.items[index]
		if len(currItem.rhsr) == 0 {
			curr.SetTerminating()
			return
		}
		newLeft := currItem.rhsr[0]
		if _, exists := leftSymbols[newLeft]; !exists {
			// add rules that have a newLeft on the left of -> (use map2)
			for _, item := range e.leftSymbolToItems[newLeft] {
				curr.AppendItem(CreateItem(item))
			}
			leftSymbols[newLeft] = struct{}{}
		}
		index++
	}

	for _, item := range curr.items {
		item.Print()
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

	for index != len(curr.items) {
		currItem := curr.items[index]
		nextNodeFirstItem := currItem.Shift()
		if _, exists := e.itemToNode[nextNodeFirstItem.Format()]; !exists {
			newNode := e.createNode(nextNodeFirstItem)
			e.itemToNode[nextNodeFirstItem.Format()] = newNode
			e.AddToItemToNode(newNode)
			curr.bridges[currItem.rhsr[0]] = e.itemToNode[nextNodeFirstItem.Format()]
			e.dfs(curr.bridges[currItem.rhsr[0]])
		}
		index++
	}

}

func (e *EDFA) parseCFS(scanner *bufio.Scanner) {
	e.createNode = nodeCreator()
	scanner.Scan()
	scanner.Scan()
	startingCFS := scanner.Text()

	newItem := CreateItem(scanner.Text())
	e.startingNode = e.createNode(newItem)
	e.allNodes = append(e.allNodes, e.startingNode)

	// map1 from item.Format() to node | item is the first item of node's items
	// map2 from left token to list of items

	e.AddToItemToNode(e.startingNode)
	e.AddToLeftSymbolToItems(startingCFS)

	scanner.Scan()
	for scanner.Text() != END {
		e.AddToLeftSymbolToItems(scanner.Text())
		scanner.Scan()
	}
	e.dfs(e.startingNode)

}

func (e *EDFA) AddToItemToNode(node *Node) {
	e.itemToNode[node.items[0].Format()] = node
}

func (e *EDFA) AddToLeftSymbolToItems(cfs string) {
	lhs := strings.Split(cfs, " ")[0]
	e.leftSymbolToItems[lhs] = append(e.leftSymbolToItems[lhs], cfs)
}
