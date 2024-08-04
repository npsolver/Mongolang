package tokenizer

import (
	"fmt"
	"strings"
)

func Scan(s string) []*Token {
	re := strings.NewReader(s)
	for {
		c, s, err := re.ReadRune()
		if err != nil {
			break
		}
		fmt.Println(c, s, err)
	}

	return []*Token{}
}
