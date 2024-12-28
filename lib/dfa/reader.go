package dfa

import (
	"bufio"
	"log"
	"os"
)

func NewFileReader(filePath string) <-chan string {
	ch := make(chan string)
	go func(filepath string, ch chan<- string) {
		file, err := os.Open(filePath)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			ch <- scanner.Text()
		}
	}(filePath, ch)
	return ch
}

func NewCommandLineReader() <-chan string {
	ch := make(chan string)
	// todo
	return ch
}
