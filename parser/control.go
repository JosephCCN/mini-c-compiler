package parser

import (
	"fmt"

	"github.com/semantic"
	"github.com/utils"
)

func if_statement(tokList *utils.TokenList) bool {
	if !tokList.Match("if") {
		return false
	}
	table := semantic.GetSymbolTable()
	table.SetParent(semantic.CurrentSymbolTable)
	node := semantic.GetSymbolTableNode(utils.RandString(8), "if", &table, semantic.Scope, -1)
	semantic.CurrentSymbolTable.Insert(&node)
	semantic.CurrentSymbolTable = &table
	defer func() {
		semantic.CurrentSymbolTable = semantic.CurrentSymbolTable.Parent()
	}()
	t1 := tokList.Match("(") && logic_experssion(tokList) && tokList.Match(")")
	initScope := semantic.Scope
	if t1 {
		semantic.Sstack.Push(utils.GetToken("if", "keyword", -1))
		semantic.Action()
		lastLabel := semantic.NextLable
		curLabel := semantic.NextLable - 1
		semantic.NextLable += 1
		t2 := tokList.Match("{") && semantic.IncreaseScope() && block_statement(tokList) && tokList.Match("}") && semantic.DecreaseScope()
		semantic.Qstack.Push(semantic.GetQuadruple("goto", "", "", fmt.Sprintf("L%d", lastLabel))) // finish block inside if statement
		semantic.Qstack.Push(semantic.GetQuadruple("label", "", "", fmt.Sprintf("L%d", curLabel))) // this is the else statement block
		if t2 && else_if_statement(tokList, lastLabel) {
			semantic.Qstack.Push(semantic.GetQuadruple("label", "", "", fmt.Sprintf("L%d", lastLabel))) // at the end of all if, else if, else statement
			return true
		}
	}
	semantic.Scope = initScope
	return false
}

func else_if_statement(tokList *utils.TokenList, lastLabel int) bool {
	if !tokList.Match("else if") {
		return false
	}
	tmp := *tokList
	tmpQ := *semantic.Qstack
	tmpS := *semantic.Sstack
	initScope := semantic.Scope
	table := semantic.GetSymbolTable()
	table.SetParent(semantic.CurrentSymbolTable)
	node := semantic.GetSymbolTableNode(utils.RandString(8), "else if", &table, semantic.Scope, -1)
	semantic.CurrentSymbolTable.Insert(&node)
	semantic.CurrentSymbolTable = &table
	defer func() {
		semantic.CurrentSymbolTable = semantic.CurrentSymbolTable.Parent()
	}()
	t1 := tokList.Match("else if") && tokList.Match("(") && logic_experssion(tokList) && tokList.Match(")")
	if t1 {
		semantic.Sstack.Push(utils.GetToken("if", "keyword", -1))
		semantic.Action()
		curLabel := semantic.NextLable - 1
		t2 := tokList.Match("{") && semantic.IncreaseScope() && block_statement(tokList) && tokList.Match("}") && semantic.DecreaseScope()
		semantic.Qstack.Push(semantic.GetQuadruple("goto", "", "", fmt.Sprintf("L%d", lastLabel))) // finish block inside if statement
		semantic.Qstack.Push(semantic.GetQuadruple("label", "", "", fmt.Sprintf("L%d", curLabel))) // this is the else statement block
		if t2 && else_statement(tokList) {
			return true
		}
	}
	*tokList = tmp
	*semantic.Qstack = tmpQ
	*semantic.Sstack = tmpS
	semantic.Scope = initScope
	return else_statement(tokList)
}

func else_statement(tokList *utils.TokenList) bool {
	if !tokList.Match("else") {
		return false
	}
	tmp := *tokList
	tmpS := *semantic.Sstack
	tmpQ := *semantic.Qstack
	initScope := semantic.Scope
	table := semantic.GetSymbolTable()
	table.SetParent(semantic.CurrentSymbolTable)
	node := semantic.GetSymbolTableNode(utils.RandString(8), "else", &table, semantic.Scope, -1)
	semantic.CurrentSymbolTable.Insert(&node)
	semantic.CurrentSymbolTable = &table
	defer func() {
		semantic.CurrentSymbolTable = semantic.CurrentSymbolTable.Parent()
	}()

	if tokList.Match("else") && tokList.Match("{") && semantic.IncreaseScope() && block_statement(tokList) && tokList.Match("}") && semantic.DecreaseScope() {
		return true
	} else {
		*semantic.Qstack = tmpQ
		*semantic.Sstack = tmpS
		*tokList = tmp
		semantic.Scope = initScope
	}
	return true
}

func while_statement(tokList *utils.TokenList) bool {
	if !tokList.Match("while") {
		return false
	}
	table := semantic.GetSymbolTable()
	table.SetParent(semantic.CurrentSymbolTable)
	node := semantic.GetSymbolTableNode(utils.RandString(8), "while", &table, semantic.Scope, -1)
	semantic.CurrentSymbolTable.Insert(&node)
	semantic.CurrentSymbolTable = &table
	defer func() {
		semantic.CurrentSymbolTable = semantic.CurrentSymbolTable.Parent()
	}()

	initScope := semantic.Scope
	curLabel := semantic.NextLable
	semantic.NextLable += 1
	semantic.Qstack.Push(semantic.GetQuadruple("label", "", "", fmt.Sprintf("L%d", curLabel)))
	t1 := tokList.Match("(") && logic_experssion(tokList) && tokList.Match(")")
	if t1 {
		semantic.Sstack.Push(utils.GetToken("if", "keyword", tokList.PrevToken().Line()))
		semantic.Action()
		lastLabel := semantic.NextLable - 1
		t2 := tokList.Match("{") && semantic.IncreaseScope() && block_statement(tokList) && tokList.Match("}") && semantic.DecreaseScope()
		if t2 {
			semantic.CurrentSymbolTable = semantic.CurrentSymbolTable.Parent()
			semantic.Qstack.Push(semantic.GetQuadruple("goto", "", "", fmt.Sprintf("L%d", curLabel)))
			semantic.Qstack.Push(semantic.GetQuadruple("label", "", "", fmt.Sprintf("L%d", lastLabel)))
			return true
		}
		semantic.CurrentSymbolTable = semantic.CurrentSymbolTable.Parent()
	}
	semantic.Scope = initScope
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
	if !tokList.Match("for") {
		return false
	}
	initScope := semantic.Scope
	semantic.IncreaseScope()
	table := semantic.GetSymbolTable()
	table.SetParent(semantic.CurrentSymbolTable)
	node := semantic.GetSymbolTableNode(utils.RandString(8), "for", &table, semantic.Scope, -1)
	semantic.CurrentSymbolTable.Insert(&node)
	semantic.CurrentSymbolTable = &table
	defer func() {
		semantic.CurrentSymbolTable = semantic.CurrentSymbolTable.Parent()
	}()

	t1 := tokList.Match("(") && for_init(tokList) && tokList.Match(";")
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
				t4 := tokList.Match("{") && block_statement(tokList) && tokList.Match("}") && semantic.DecreaseScope()
				if t4 {
					semantic.Qstack.Push(semantic.GetQuadruple("goto", "", "", fmt.Sprintf("L%d", incrLabel)))
					semantic.Qstack.Push(semantic.GetQuadruple("label", "", "", fmt.Sprintf("L%d", curLabel+2)))
					return true
				}
			}
		}
	}
	semantic.Scope = initScope
	return false
}
