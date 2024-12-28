package main

import (
	"fmt"

	"github.com/Npsolver/Mongolang/compiler/scanner"
)

func main() {
	testString := `db.getCollection("something").find({"business_id":{$in:["12345"]}})`
	fmt.Println(testString)

	tks, err := scanner.Scan(testString)

}
