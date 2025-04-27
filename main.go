package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/lexical"
	"github.com/parser"
	"github.com/semantic"
	"github.com/utils"
)

func main() {

	semantic.Init()
	// rand.Seed(time.Now().UnixNano())

	var srcPath string
	var outputDir string
	flag.StringVar(&srcPath, "s", "", "source file path")
	flag.StringVar(&outputDir, "d", "", "destination")
	flag.Parse()

	if srcPath == "" {
		fmt.Println(utils.RedString("Missing source"))
		return
	}
	if outputDir == "" {
		fmt.Println(utils.RedString("Missing destination"))
		return
	}
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		os.Mkdir(outputDir, 0777)
	}

	src, err := os.ReadFile(srcPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	tokenList, err := lexical.Run(string(src))
	lines := strings.Split(string(src), "\n")
	os.WriteFile(filepath.Join(outputDir, "lexical"), []byte(tokenList.ListAll()), 0777)
	// fmt.Println(tokenList.ListAll())
	if err != nil {
		lst := tokenList.List()
		errTok := lst[len(lst)-1]
		errLine := errTok.Line()
		fmt.Println(utils.RedString(fmt.Sprintf("Unmatch element in line %d:", errLine)))
		fmt.Println(lines[errLine-1])
		return
	}

	parserRes := parser.Start(&tokenList)
	// fmt.Println(parserRes)
	if !parserRes {
		errTokenId := utils.TokenListDeepest - 1
		errTok := tokenList.GetToken(errTokenId)
		errLine := errTok.Line()
		fmt.Println(utils.RedString((fmt.Sprintf("Error in line %d:", errLine))))
		fmt.Println(lines[errLine-1])

	}
	os.WriteFile(filepath.Join(outputDir, "sstack"), []byte(semantic.Sstack.ListAll()), 0777)
	os.WriteFile(filepath.Join(outputDir, "result"), []byte(semantic.Qstack.ListAll()), 0777)
	os.WriteFile(filepath.Join(outputDir, "symbolTable"), []byte(semantic.RootSymbolTable.ListAll()), 0777)
	// fmt.Println("Sstack:", semantic.Sstack)
	// fmt.Println("QStack:\n", semantic.Qstack.ListAll())
	// fmt.Println("Symbol Table:\n", semantic.RootSymbolTable.ListAll())

}
