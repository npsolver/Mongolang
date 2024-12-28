package cfg

import (
	"bufio"
	"log"
	"os"
)

type CFG struct {
	startingNode *node
	stack        []*node
}

func NewCFG(filePath string) *CFG {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fileScanner := bufio.NewScanner(file)

	cfg := &CFG{}
	cfg.scanGrammar(fileScanner)

	return cfg
}
