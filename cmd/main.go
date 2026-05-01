package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/familydabba/markdown2ansi"
)

func main() {
	var input string

	if len(os.Args) > 1 {
		data, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading file:", err)
			os.Exit(1)
		}
		input = string(data)
	} else {
		data, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading stdin:", err)
			os.Exit(1)
		}
		input = string(data)
	}

	result := markdown2ansi.Render(input)
	fmt.Print(result)
}
