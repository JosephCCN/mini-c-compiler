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
	return declaration(tokList)
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
	return true
}

func return_statement(tokList *utils.TokenList) bool {
	return true
}

func declaration(tokList *utils.TokenList) bool {
	return true
}
