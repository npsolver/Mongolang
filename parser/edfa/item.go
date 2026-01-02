package edfa

import (
	"fmt"
	"strings"

	"github.com/npsolver/Mongolang/global"
)

// An example of an Item:
// START -> BOF * COMMAND EOF

// lhs: START
// rhsl : [BOF]
// rhsr : [COMMAND, EOF]

type Item struct {
	lhs  string
	rhsl []string
	rhsr []string
}

func CreateItem(cfs string) *Item {
	cfsSymbols := strings.Split(cfs, " ")
	lhs := cfsSymbols[0]
	rhsl := []string{}
	rhsr := []string{}
	for i := 2; i < len(cfsSymbols); i++ {
		rhsr = append(rhsr, cfsSymbols[i])
	}

	newNodeItem := &Item{lhs, rhsl, rhsr}
	return newNodeItem
}

func (item *Item) Shift() *Item {
	lhs := item.lhs
	rhsl := append([]string(nil), item.rhsl...)
	rhsl = append(rhsl, item.rhsr[0])
	rhsr := append([]string(nil), item.rhsr[1:]...)
	return &Item{lhs, rhsl, rhsr}
}

func (i *Item) Format() string {
	symbolList := []string{}
	symbolList = append(symbolList, i.lhs, global.Arrow)
	symbolList = append(symbolList, i.rhsl...)
	symbolList = append(symbolList, global.Separator)
	symbolList = append(symbolList, i.rhsr...)
	return strings.Join(symbolList, " ")
}

func (i *Item) Print() {
	fmt.Println(i.Format())
}
