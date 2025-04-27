package semantic

type FunctionTableNode struct {
	returnType    string
	paraNum       int
	paraTable     *SymbolTable
	localVarTable *SymbolTable
}

type FunctionTable struct {
	table []*FunctionTableNode
}

func GetFunctionTable() FunctionTable {
	ret := FunctionTable{
		table: make([]*FunctionTableNode, 0),
	}
	return ret
}

func GetFunctionTableNode(returnType string, paraNum int, paraTable *SymbolTable, localVarTable *SymbolTable) FunctionTableNode {
	ret := FunctionTableNode{
		returnType:    returnType,
		paraNum:       paraNum,
		paraTable:     paraTable,
		localVarTable: localVarTable,
	}
	return ret
}
