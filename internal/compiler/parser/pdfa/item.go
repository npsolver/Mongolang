package cfg

import "github.com/Npsolver/Mongolang/internal/global"

type Item struct {
	from  global.Token
	left  []global.Token
	right []global.Token
}
