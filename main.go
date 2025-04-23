package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/lexical"
	"github.com/parser"
	"github.com/utils"
)

func redString(src string) string {
	var Red = "\033[31m"
	var Reset = "\033[0m"
	return Red + src + Reset
}

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
	lines := strings.Split(string(src), "\n")
	fmt.Println(tokenList)
	if err != nil {
		lst := tokenList.GetList()
		errTok := lst[len(lst)-1]
		errLine := errTok.GetLine()
		fmt.Println(redString(fmt.Sprintf("Unmatch element in line %d:", errLine)))
		fmt.Println(lines[errLine-1])
		return
	}

	parserRes := parser.Start(&tokenList)
	fmt.Println(parserRes)
	if !parserRes {
		errTokenId := utils.TokenListDeepest - 1
		errTok := tokenList.GetToken(errTokenId)
		errLine := errTok.GetLine()
		fmt.Println(redString(fmt.Sprintf("Syntax error in line %d:", errLine)))
		fmt.Println(lines[errLine-1])

	}

}
