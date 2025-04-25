package utils

type SymbolTableNode struct {
	name        string
	Type        string
	typePointer any
	scope       int
	size        int
}

type SymbolTable struct {
	table []*SymbolTableNode
}

func GetSymbolTable() SymbolTable {
	ret := SymbolTable{
		table: make([]*SymbolTableNode, 0),
	}
	return ret
}

func GetSymbolTableNode(name string, Type string, typePointer any, scope int, size int) SymbolTableNode {
	ret := SymbolTableNode{
		name:        name,
		Type:        Type,
		typePointer: typePointer,
		scope:       scope,
		size:        size,
	}
	return ret
}

func (t *SymbolTable) Insert(node *SymbolTableNode) {
	t.table = append(t.table, node)
}

func (t *SymbolTable) Find(name string) *SymbolTableNode {
	for _, node := range t.table {
		if node.name == name {
			return node
		}
	}
	return nil
}
