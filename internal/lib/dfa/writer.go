package dfa

import (
	"bufio"
	"os"
)

func NewFileWriter(filename string) (chan<- string, error) {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	buf := bufio.NewWriter(f)
	ch := make(chan string)

	go func() {
		defer func() {
			buf.Flush() // ensure all buffered data is written
			f.Close()   // close file last
		}()

		for msg := range ch {
			buf.WriteString(msg)
		}
	}()

	return ch, nil
}

func NewCommandLineWriter() chan<- string {
	ch := make(chan string)
	// todo
	return ch
}
