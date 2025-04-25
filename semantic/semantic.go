package semantic

import "github.com/utils"

var RootSymbolTable *utils.SymbolTable

func Init() {
	*RootSymbolTable = utils.GetSymbolTable()
}
