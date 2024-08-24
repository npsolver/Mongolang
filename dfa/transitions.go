package dfa

import (
	"bufio"
	"strings"
)

func (d *DFA) scanTransitions(scanner *bufio.Scanner) {
	scanner.Scan()
	for scanner.Text() != END {
		transitions := scanner.Text()
		temp := strings.Split(strings.Trim(transitions, " "), " ")
		prev := d.getState(StateName(temp[0]))
		bridges := temp[1 : len(temp)-1]
		next := d.getState(StateName(temp[len(temp)-1]))
		for _, c := range bridges {
			if len(c) == 3 {
				for i := c[0]; i != c[2]; i++ {
					prev.addBridge(Symbol(i), next)
				}
				prev.addBridge(Symbol(c[2]), next)
			} else {
				prev.addBridge(Symbol(c), next)
			}
		}
		scanner.Scan()
	}
}
