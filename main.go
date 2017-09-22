package main

import (
	"fmt"
)

func main() {
	dotFile := "examples/step2-problem.dot"
	dddModel := Parse(dotFile)

	codeDir := "examples/step2-java/html"
	codeModel := ParseCodeDir(codeDir)
	fmt.Println(dddModel.Compare(codeModel))
}
