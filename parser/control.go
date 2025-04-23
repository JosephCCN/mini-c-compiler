package parser

import "github.com/utils"

func if_statement(tokList *utils.TokenList) bool {
	return tokList.Match("if") && tokList.Match("(") && logic_experssion(tokList) && tokList.Match(")") &&
		tokList.Match("{") && block_statement(tokList) && tokList.Match("}")
}

func while_statement(tokList *utils.TokenList) bool {
	return tokList.Match("while") && tokList.Match("(") && logic_experssion(tokList) && tokList.Match(")") &&
		tokList.Match("{") && block_statement(tokList) && tokList.Match("}")
}

func for_init(tokList *utils.TokenList) bool {
	tmp := tokList.ShallowCopy()
	if var_declaration(tokList) {
		return true
	}
	*tokList = tmp.ShallowCopy()
	if assignment(tokList) {
		return true
	}
	*tokList = tmp.ShallowCopy()
	return true
}

func for_condition(tokList *utils.TokenList) bool {
	tmp := tokList.ShallowCopy()
	if logic_experssion(tokList) {
		return true
	}
	*tokList = tmp.ShallowCopy()
	return true
}

func for_incr(tokList *utils.TokenList) bool {
	tmp := tokList.ShallowCopy()
	if assignment(tokList) {
		return true
	}
	*tokList = tmp.ShallowCopy()
	return true
}

func for_statement(tokList *utils.TokenList) bool {
	return tokList.Match("for") && tokList.Match("(") && for_init(tokList) && tokList.Match(";") &&
		for_condition(tokList) && tokList.Match(";") && for_incr(tokList) && tokList.Match(")") &&
		tokList.Match("{") && block_statement(tokList) && tokList.Match("}")
}
