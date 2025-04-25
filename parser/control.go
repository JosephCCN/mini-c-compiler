package parser

import "github.com/utils"

func if_statement(tokList *utils.TokenList) bool {
	return tokList.Match("if") && tokList.Match("(") && logic_experssion(tokList) && tokList.Match(")") &&
		tokList.Match("{") && block_statement(tokList) && tokList.Match("}") && else_if_statement(tokList)
}

func else_if_statement(tokList *utils.TokenList) bool {
	tmp := *tokList
	t := tokList.Match("else if") && tokList.Match("(") && logic_experssion(tokList) && tokList.Match(")") &&
		tokList.Match("{") && block_statement(tokList) && tokList.Match("}") && else_if_statement(tokList)
	if t {
		return true
	}
	*tokList = tmp
	return else_statement(tokList)
}

func else_statement(tokList *utils.TokenList) bool {
	tmp := *tokList
	if tokList.Match("else") && tokList.Match("{") && block_statement(tokList) && tokList.Match("}") {
		return true
	} else {
		*tokList = tmp
	}
	return true
}

func while_statement(tokList *utils.TokenList) bool {
	return tokList.Match("while") && tokList.Match("(") && logic_experssion(tokList) && tokList.Match(")") &&
		tokList.Match("{") && block_statement(tokList) && tokList.Match("}")
}

func for_init(tokList *utils.TokenList) bool {
	tmp := *tokList
	if var_declaration(tokList) {
		return true
	}
	*tokList = tmp
	if assignment(tokList) {
		return true
	}
	*tokList = tmp
	return true
}

func for_condition(tokList *utils.TokenList) bool {
	tmp := *tokList
	if logic_experssion(tokList) {
		return true
	}
	*tokList = tmp
	return true
}

func for_incr(tokList *utils.TokenList) bool {
	tmp := *tokList
	if assignment(tokList) {
		return true
	}
	*tokList = tmp
	return true
}

func for_statement(tokList *utils.TokenList) bool {
	return tokList.Match("for") && tokList.Match("(") && for_init(tokList) && tokList.Match(";") &&
		for_condition(tokList) && tokList.Match(";") && for_incr(tokList) && tokList.Match(")") &&
		tokList.Match("{") && block_statement(tokList) && tokList.Match("}")
}
