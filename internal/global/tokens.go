package global

import (
	"fmt"
	"strings"
)

type Token struct {
	tokenType TerminalSymbol
	val       string
}

func (tk *Token) Print() {
	fmt.Println(tk.tokenType.GetName(), tk.val)
}

func NewToken(name, val string) *Token {
	return &Token{NewTerminalSymbol(strings.ToUpper(name)), val}
}
