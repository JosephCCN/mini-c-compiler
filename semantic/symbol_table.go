package semantic

import (
	"fmt"
	"strings"
)

type SymbolTableNode struct {
	name        string
	Type        string
	typePointer any
	scope       int
	size        int
}

type SymbolTable struct {
	table  []*SymbolTableNode
	parent *SymbolTable
}

func GetSymbolTable() SymbolTable {
	ret := SymbolTable{
		table:  make([]*SymbolTableNode, 0),
		parent: nil,
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

func (t *SymbolTable) SetParent(st *SymbolTable) {
	t.parent = st
}

func (t *SymbolTable) Parent() *SymbolTable {
	return t.parent
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

func (t *SymbolTable) Find(name string) *SymbolTableNode {
	for _, node := range t.table {
		if node.name == name {
			return node
		}
	}
	return nil
}

func (t *SymbolTable) ListAll() string {
	ret := ""
	for _, n := range t.table {
		if strings.HasPrefix(n.Type, "funct") || n.Type == "for" || n.Type == "while" || n.Type == "if" || n.Type == "else if" || n.Type == "else" {
			t, ok := n.typePointer.(*SymbolTable)
			if ok {
				ret += fmt.Sprintf("%s %s:\n%s", n.Type, n.name, t.ListAll())
			}
		} else {
			ret += fmt.Sprintf("[%s, %s, %d]\n", n.name, n.Type, n.scope)
		}
	}
	return ret
}
