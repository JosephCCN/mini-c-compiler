package semantic

import "github.com/utils"

var RootSymbolTable *utils.SymbolTable
var Sstack *SemanticStack

func Init() {
	tmpST := utils.GetSymbolTable()
	RootSymbolTable = &tmpST
	tmpSS := GetSemanticStack()
	Sstack = &tmpSS
}
