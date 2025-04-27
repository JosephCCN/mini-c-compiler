package semantic

type FunctionTable struct {
	returnType string
	paraNum    int
	parameter  *SymbolTable
	variable   *SymbolTable
}

func GetFunctionTable(returnType string, paraNum int, parameter *SymbolTable, variable *SymbolTable) FunctionTable {
	ret := FunctionTable{
		returnType: returnType,
		paraNum:    paraNum,
		parameter:  parameter,
		variable:   variable,
	}
	return ret
}
