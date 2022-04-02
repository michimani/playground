package main

import (
	"flag"
	"fmt"
	"os"
	"playground/cmd"
)

var (
	funcCode *string = flag.String("f", "", "func code")
)

func usage() {
	u := `
Usage:
  go run . [flags] [values]
Flags:
  -f (string)
    Function code.
`

	fmt.Println(u)
	fmt.Println("Function code list")
	for k := range funcMap {
		fmt.Println(k)
	}
}

func main() {
	flag.Usage = usage
	flag.Parse()

	os.Exit(run())
}

func run() int {
	if funcCode == nil || *funcCode == "" {
		usage()
		return 1
	}

	if _, ok := funcMap[*funcCode]; !ok {
		fmt.Println("undefined func code")
		return 1
	}

	return funcMap[*funcCode]()
}

var funcMap map[string]func() int = map[string]func() int{
	"hello":   cmd.Hello,
	"timeinp": cmd.InAndParseInLocation,
}
