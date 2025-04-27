package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/lexical"
	"github.com/parser"
	"github.com/semantic"
	"github.com/utils"
)

func main() {

	semantic.Init()
	rand.Seed(time.Now().UnixNano())

	var srcPath string
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
		lst := tokenList.List()
		errTok := lst[len(lst)-1]
		errLine := errTok.Line()
		fmt.Println(utils.RedString(fmt.Sprintf("Unmatch element in line %d:", errLine)))
		fmt.Println(lines[errLine-1])
		return
	}

	parserRes := parser.Start(&tokenList)
	fmt.Println(parserRes)
	if !parserRes {
		errTokenId := utils.TokenListDeepest - 1
		errTok := tokenList.GetToken(errTokenId)
		errLine := errTok.Line()
		fmt.Println(utils.RedString((fmt.Sprintf("Error in line %d:", errLine))))
		fmt.Println(lines[errLine-1])

	}
	fmt.Println("Sstack:", semantic.Sstack)
	fmt.Println("QStack:\n", semantic.Qstack.ListAll())
	fmt.Println("Symbol Table:\n", semantic.RootSymbolTable.ListAll())

}
