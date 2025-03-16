package main

import (
	"fmt"

	"github.com/lexical"
)

func main() {

	// var (
	// 	srcPath string
	// )
	// flag.StringVar(&srcPath, "s", "", "source file path")
	// flag.Parse()

	// if srcPath == "" {
	// 	fmt.Println("Missing source")
	// 	return
	// }

	// src, err := os.ReadFile(srcPath)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	fmt.Println(lexical.Run(string("-12")))

}
