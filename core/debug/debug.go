package debug

import (
	"fmt"
)

const DEBUG = false

func DebugPrint(format string, a ...any) {
	if DEBUG {
		fmt.Printf(format, a...)
	}
}
