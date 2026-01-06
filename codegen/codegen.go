package codegen

import (
	"fmt"
	"strings"

	"github.com/npsolver/Mongolang/debug"
	"github.com/npsolver/Mongolang/global"
)

func formatID(node *global.TreeNode) string {
	return "\"" + node.Symbol.GetValue() + "\""
}

func formatFieldName(node *global.TreeNode) string {
	return formatID(node.Children[0])
}

func formatOperator(node *global.TreeNode) string {
	return "\"" + "$" + node.Children[1].Symbol.GetValue() + "\""
}

func formatInt(i string) string {
	return i
}

func formatValue(node *global.TreeNode) string {
	if node.Children[0].Symbol.GetName() == "LIST" {
		return formatBSON(node.Children[0])
	} else if node.Children[0].Symbol.GetName() == "SET" {
		return formatBSON(node.Children[0])
	} else if node.Children[0].Symbol.GetName() == "ID" {
		return formatID(node.Children[0])
	} else {
		return node.Children[0].Symbol.GetValue()
	}
}

func formatField(node *global.TreeNode) string {
	if node.Children[0].Symbol.GetName() == "OPERATOR" {
		return formatOperator(node.Children[0]) + ": " + formatValue(node.Children[2])
	} else {
		return formatFieldName(node.Children[0]) + ": " + formatValue(node.Children[2])
	}
}

func formatFields(node *global.TreeNode, hasComma bool) string {
	answer := ""
	if node.Children[0].Symbol.GetName() == "EMPTY" {
		answer = ""
	} else if node.Children[0].Symbol.GetName() == "FIELD" {
		answer = formatField(node.Children[0])
	} else {
		answer = formatFields(node.Children[0], true) + formatField(node.Children[2])
	}
	if hasComma {
		return answer + ",\n"
	} else {
		return answer + "\n"
	}
}

func formatSet(node *global.TreeNode) string {
	return "bson.M{\n" + formatFields(node.Children[1], false) + "}"
}

func formatSets(node *global.TreeNode, hasComma bool) string {
	answer := ""
	if node.Children[0].Symbol.GetName() == "EMPTY" {
		answer = ""
	} else if node.Children[0].Symbol.GetName() == "SET" {
		answer = formatSet(node.Children[0])
	} else {
		answer = formatSets(node.Children[0], true) + formatSet(node.Children[2])
	}
	if hasComma {
		return answer + ",\n"
	} else {
		return answer + "\n"
	}
}

func formatBSON(node *global.TreeNode) string {
	if node.Symbol.GetName() == "SETS" {
		return formatSets(node, false)
	} else if node.Symbol.GetName() == "SET" { // change this to use const vars
		return formatSet(node)
	} else {
		return "bson.A{\n" + formatBSON(node.Children[1]) + "}"
	}
}

func addTabs(s string, tabs int) string {
	output := ""
	for i := 0; i < tabs; i++ {
		output += "\t"
	}
	return output + s
}

func Generate(start *global.TreeNode) string {

	collectionCommand := start.Children[1].Children[4]

	if collectionCommand.Children[0].Symbol.GetValue() == "find" {
		generated := formatBSON(collectionCommand.Children[2].Children[0].Children[0])
		genList := strings.Split(generated, "\n")
		if debug.DEBUG {
			fmt.Println(genList)
		}
		tabbedList := []string{}
		currTabs := 0
		for _, s := range genList {
			if s[len(s)-1:] == "{" {
				tabbedList = append(tabbedList, addTabs(s, currTabs))
				currTabs += 1
			} else if s[len(s)-1:] == "}" || s[len(s)-1:] == "," {
				currTabs -= 1
				tabbedList = append(tabbedList, addTabs(s, currTabs))
			} else {
				tabbedList = append(tabbedList, addTabs(s, currTabs))
			}
		}
		return strings.Join(tabbedList, "\n")
	}

	return ""
}
