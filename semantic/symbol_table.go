package semantic

import "fmt"

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

func (t *SymbolTable) Insert(node *SymbolTableNode) bool {
	for _, nd := range t.table {
		if nd.name == node.name && nd.scope == node.scope {
			return false
		}
	}
	t.table = append(t.table, node)
	return true
}

func (t *SymbolTable) Find(name string, scope int) *SymbolTableNode {
	for _, node := range t.table {
		if node.name == name && node.scope == scope {
			return node
		}
	}
	return nil
}

func (t *SymbolTable) ListAll() string {
	ret := ""
	for _, n := range t.table {
		ret += fmt.Sprintf("[%s, %s, %d]\n", n.name, n.Type, n.scope)
	}
	return ret
}
