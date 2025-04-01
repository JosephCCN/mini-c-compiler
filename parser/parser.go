package parser

import (
	"github.com/utils"
)

// this is the only export function
func Start(tokList *utils.TokenList) bool {
	tmp := tokList.Copy()
	if ex_declaration(tokList) {
		return true
	}
	tokList = tmp.Copy()
	return Start(tokList) && ex_declaration(tokList)
}

func ex_declaration(tokList *utils.TokenList) bool {
	tmp := tokList.Copy()
	if function(tokList) {
		return true
	}
	tokList = tmp.Copy()
	return var_declaration(tokList)
}

func function(tokList *utils.TokenList) bool {
	return (types(tokList) && tokList.IsIdentifier() && tokList.Match("(") && tokList.Match(")") && tokList.Match("{") &&
		block_statement(tokList) && return_statement(tokList) && tokList.Match("}"))
}

func types(tokList *utils.TokenList) bool {
	tmp := tokList.Copy()
	types := []string{"int", "double", "char", "string"}
	for _, tp := range types {
		if tokList.Match(tp) {
			return true
		}
		tokList = tmp.Copy()
	}
	return false
}

func block_statement(tokList *utils.TokenList) bool {
	tmp := tokList.Copy()
	if statement(tokList) {
		return true
	}
	tokList = tmp.Copy()
	return block_statement(tokList) && statement(tokList)
}

// LL(1) parsing
func statement(tokList *utils.TokenList) bool {
	return true
}

func return_statement(tokList *utils.TokenList) bool {
	return true
}

func var_declaration(tokList *utils.TokenList) bool {
	return types(tokList) && assignment(tokList)
}

func type_ep(tokList *utils.TokenList) bool {
	tmp := tokList.Copy()
	if tokList.IsInt() {
		return true
	}
	if tokList.IsDouble() {
		return true
	}
	tokList = tmp.Copy()
	if tokList.IsChar() {
		return true
	}
	tokList = tmp.Copy()
	return tokList.IsString()
}

func assignment(tokList *utils.TokenList) bool {
	tmp := tokList.Copy()
	if tokList.IsIdentifier() && tokList.Match("=") && tokList.IsIdentifier() {
		return true
	}
	tokList = tmp.Copy()
	return tokList.IsIdentifier() && tokList.Match("=") && expressions(tokList)
}

func if_statement(tokList *utils.TokenList) bool {
	return true
}
