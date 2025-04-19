package parser

import (
	"github.com/utils"
)

// this is the only export function
func Start(tokList *utils.TokenList) bool {
	return ex_declaration(tokList) && start_(tokList)
}

func start_(tokList *utils.TokenList) bool {
	tmp := tokList.ShallowCopy()
	if ex_declaration(tokList) {
		start_(tokList)
	} else {
		*tokList = tmp.ShallowCopy()
	}
	return true
}

func ex_declaration(tokList *utils.TokenList) bool {
	tmp := tokList.ShallowCopy()
	if function(tokList) {
		return true
	}
	*tokList = tmp.ShallowCopy()
	if var_declaration(tokList) {
		return true
	}
	*tokList = tmp.ShallowCopy()
	return false
}

func function(tokList *utils.TokenList) bool {
	t := types(tokList) && tokList.IsIdentifier() && tokList.Match("(") && tokList.Match(")") && tokList.Match("{") &&
		block_statement(tokList) && return_statement(tokList) && tokList.Match("}")
	return t
}

func types(tokList *utils.TokenList) bool {
	tmp := tokList.ShallowCopy()
	types := []string{"int", "double", "char", "string"}
	for _, tp := range types {
		if tokList.Match(tp) {
			return true
		}
		*tokList = tmp.ShallowCopy()
	}
	return false
}

func block_statement_(tokList *utils.TokenList) bool {
	tmp := tokList.ShallowCopy()
	if statement(tokList) {
		block_statement_(tokList)
	} else {
		*tokList = tmp.ShallowCopy()
	}
	return true
}

func block_statement(tokList *utils.TokenList) bool {
	t := statement(tokList) && block_statement_(tokList)
	return t
}

// LL(1) parsing
func statement(tokList *utils.TokenList) bool {
	return var_declaration(tokList)
}

func return_statement(tokList *utils.TokenList) bool {
	return true
}

func var_declaration(tokList *utils.TokenList) bool {
	return types(tokList) && assignment(tokList)
}

func assignment(tokList *utils.TokenList) bool {
	tmp := tokList.ShallowCopy()
	if tokList.IsIdentifier() && tokList.Match("=") && tokList.IsIdentifier() {
		return true
	}
	*tokList = tmp.ShallowCopy()
	return tokList.IsIdentifier() && tokList.Match("=") && expressions(tokList)
}

func if_statement(tokList *utils.TokenList) bool {
	return true
}
