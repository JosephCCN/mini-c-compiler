package parser

import (
	"fmt"

	"github.com/semantic"
	"github.com/utils"
)

func if_statement(tokList *utils.TokenList) bool {
	t1 := tokList.Match("if") && tokList.Match("(") && logic_experssion(tokList) && tokList.Match(")")
	if t1 {
		semantic.Sstack.Push(utils.GetToken("if", "keyword", -1))
		semantic.Action()
		lastLabel := semantic.NextLable
		curLabel := semantic.NextLable - 1
		semantic.NextLable += 1
		t2 := tokList.Match("{") && block_statement(tokList) && tokList.Match("}")
		semantic.Qstack.Push(semantic.GetQuadruple("goto", "", "", fmt.Sprintf("L%d", lastLabel))) // finish block inside if statement
		semantic.Qstack.Push(semantic.GetQuadruple("label", "", "", fmt.Sprintf("L%d", curLabel))) // this is the else statement block
		if t2 && else_if_statement(tokList, lastLabel) {
			semantic.Qstack.Push(semantic.GetQuadruple("label", "", "", fmt.Sprintf("L%d", lastLabel))) // at the end of all if, else if, else statement
			return true
		}
	}
	return false
}

func else_if_statement(tokList *utils.TokenList, lastLabel int) bool {
	tmp := *tokList
	tmpQ := *semantic.Qstack
	tmpS := *semantic.Sstack
	t1 := tokList.Match("else if") && tokList.Match("(") && logic_experssion(tokList) && tokList.Match(")")
	if t1 {
		semantic.Sstack.Push(utils.GetToken("if", "keyword", -1))
		semantic.Action()
		curLabel := semantic.NextLable - 1
		t2 := tokList.Match("{") && block_statement(tokList) && tokList.Match("}")
		semantic.Qstack.Push(semantic.GetQuadruple("goto", "", "", fmt.Sprintf("L%d", lastLabel))) // finish block inside if statement
		semantic.Qstack.Push(semantic.GetQuadruple("label", "", "", fmt.Sprintf("L%d", curLabel))) // this is the else statement block
		if t2 && else_statement(tokList) {
			return true
		}
	}
	*tokList = tmp
	*semantic.Qstack = tmpQ
	*semantic.Sstack = tmpS
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
	curLabel := semantic.NextLable
	semantic.NextLable += 1
	semantic.Qstack.Push(semantic.GetQuadruple("label", "", "", fmt.Sprintf("L%d", curLabel)))
	t1 := tokList.Match("while") && tokList.Match("(") && logic_experssion(tokList) && tokList.Match(")")
	if t1 {
		semantic.Sstack.Push(utils.GetToken("if", "keyword", tokList.PrevToken().Line()))
		semantic.Action()
		lastLabel := semantic.NextLable - 1
		t2 := tokList.Match("{") && block_statement(tokList) && tokList.Match("}")
		if t2 {
			semantic.Qstack.Push(semantic.GetQuadruple("goto", "", "", fmt.Sprintf("L%d", curLabel)))
			semantic.Qstack.Push(semantic.GetQuadruple("label", "", "", fmt.Sprintf("L%d", lastLabel)))
			return true
		}
	}
	return false
}

func for_init(tokList *utils.TokenList) bool {
	tmp := *tokList
	tmpQ := *semantic.Qstack
	tmpS := *semantic.Sstack
	if var_declaration(tokList) {
		return true
	}
	*tokList = tmp
	if assignment(tokList) {
		return true
	}
	*tokList = tmp
	*semantic.Qstack = tmpQ
	*semantic.Sstack = tmpS
	return true
}

func for_condition(tokList *utils.TokenList) bool {
	tmp := *tokList
	tmpQ := *semantic.Qstack
	tmpS := *semantic.Sstack
	if logic_experssion(tokList) {
		return true
	}
	*tokList = tmp
	*semantic.Qstack = tmpQ
	*semantic.Sstack = tmpS
	return true
}

func for_incr(tokList *utils.TokenList) bool {
	tmp := *tokList
	tmpQ := *semantic.Qstack
	tmpS := *semantic.Sstack
	if assignment(tokList) {
		return true
	}
	*tokList = tmp
	*semantic.Qstack = tmpQ
	*semantic.Sstack = tmpS
	return true
}

func for_statement(tokList *utils.TokenList) bool {
	t1 := tokList.Match("for") && tokList.Match("(") && for_init(tokList) && tokList.Match(";")
	if t1 {
		curLabel := semantic.NextLable
		semantic.NextLable += 1
		semantic.Qstack.Push(semantic.GetQuadruple("label", "", "", fmt.Sprintf("L%d", curLabel)))
		t2 := for_condition(tokList) && tokList.Match(";")
		if t2 {
			semantic.Sstack.Push(utils.GetToken("if", "keyword", tokList.PrevToken().Line()))
			semantic.Action()
			condLabelQ := semantic.Qstack.Pop()
			incrLabel := semantic.NextLable
			semantic.NextLable += 1
			semantic.Qstack.Push(semantic.GetQuadruple("label", "", "", fmt.Sprintf("L%d", incrLabel)))
			t3 := for_incr(tokList) && tokList.Match(")")
			if t3 {
				semantic.Qstack.Push(semantic.GetQuadruple("goto", "", "", fmt.Sprintf("L%d", curLabel)))
				semantic.Qstack.Push(condLabelQ)
				t4 := tokList.Match("{") && block_statement(tokList) && tokList.Match("}")
				if t4 {
					semantic.Qstack.Push(semantic.GetQuadruple("goto", "", "", fmt.Sprintf("L%d", incrLabel)))
					semantic.Qstack.Push(semantic.GetQuadruple("label", "", "", fmt.Sprintf("L%d", curLabel+2)))
					return true
				}
			}
		}
	}
	return false
}
