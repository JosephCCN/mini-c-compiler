package parser

import (
	"github.com/utils"
)

// this is the only export function
func Start(tokList *utils.TokenList) bool {
	res := ex_declaration(tokList) && start_(tokList)
	return res && tokList.End()
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
	return var_declaration(tokList) && tokList.Match(";")
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

func type_ep(tokList *utils.TokenList) bool {
	tmp := tokList.ShallowCopy()
	if tokList.IsInt() {
		return true
	}
	*tokList = tmp.ShallowCopy()
	if tokList.IsChar() {
		return true
	}
	*tokList = tmp.ShallowCopy()
	if tokList.IsString() {
		return true
	}
	*tokList = tmp.ShallowCopy()
	return tokList.IsDouble()
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
	tmp := tokList.ShallowCopy()
	if types(tokList) {
		*tokList = tmp.ShallowCopy()
		return var_declaration(tokList) && tokList.Match(";")
	}
	*tokList = tmp.ShallowCopy()
	if tokList.IsIdentifier() {
		*tokList = tmp.ShallowCopy()
		return assignment(tokList) && tokList.Match(";")
	}
	*tokList = tmp.ShallowCopy()
	if tokList.Match("if") {
		*tokList = tmp.ShallowCopy()
		return if_statement(tokList)
	}
	*tokList = tmp.ShallowCopy()
	if tokList.Match("for") {
		*tokList = tmp.ShallowCopy()
		return for_statement(tokList)
	}
	*tokList = tmp.ShallowCopy()
	if tokList.Match("while") {
		*tokList = tmp.ShallowCopy()
		return while_statement(tokList)
	}
	return false
}

func return_statement(tokList *utils.TokenList) bool {
	tmp := tokList.ShallowCopy()
	if tokList.Match("return") && type_ep(tokList) && tokList.Match(";") {
		return true
	}
	*tokList = tmp.ShallowCopy()
	return tokList.Match("return") && tokList.IsIdentifier() && tokList.Match(";")
}

func var_declaration(tokList *utils.TokenList) bool {
	return types(tokList) && assignment(tokList)
}

func assignment(tokList *utils.TokenList) bool {
	tmp := tokList.ShallowCopy()
	if tokList.IsIdentifier() && tokList.Match("=") && expressions(tokList) {
		return true
	}
	*tokList = tmp.ShallowCopy()
	return tokList.IsIdentifier() && tokList.Match("=") && tokList.IsIdentifier()
}
