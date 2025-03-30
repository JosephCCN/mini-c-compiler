package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/lexical"
	"github.com/parser"
)

func main() {

	var (
		srcPath string
	)
	flag.StringVar(&srcPath, "s", "", "source file path")
	flag.Parse()

	if srcPath == "" {
		fmt.Println("Missing source")
		return
	}

	src, err := os.ReadFile(srcPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	tokenList, err := lexical.Run(string(src))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tokenList)
	// for _, tok := range tokenList.GetList() {
	// 	if tok.IsInt() {
	// 		fmt.Println("yes")
	// 	}
	// }

	fmt.Println(parser.Start(&tokenList))

}
