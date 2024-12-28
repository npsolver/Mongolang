package cfg

import "github.com/Npsolver/Mongolang/internal/global"

type node struct {
	item    []Item
	bridges map[global.Token]*node
}
