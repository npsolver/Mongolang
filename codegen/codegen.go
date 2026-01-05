package codegen

import (
	"fmt"

	"github.com/npsolver/Mongolang/global"
)

func Generate(start *global.TreeNode) string {

	collectionCommand := start.Children[1].Children[4]

	if collectionCommand.Children[0].Symbol.GetValue() == "find" {
		fmt.Println("called find")
	} else {
		fmt.Println("did not call find")
	}

	return ""
}
