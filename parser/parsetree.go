package parser

import "github.com/npsolver/Mongolang/global"

type TreeNode struct {
	symbol *global.Symbol
	next   *TreeNode
}
