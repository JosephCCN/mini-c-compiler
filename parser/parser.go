package parser

import (
	"fmt"

	"github.com/semantic"
	"github.com/utils"
)

// this is the only export function
func Start(tokList *utils.TokenList) bool {
	res := ex_declaration(tokList) && start_(tokList)
	return res && tokList.End()
	// return for_statement(tokList)
}

func start_(tokList *utils.TokenList) bool {
	tmp := *tokList
	if ex_declaration(tokList) {
		start_(tokList)
	} else {
		*tokList = tmp
	}
	return true
}

func ex_declaration(tokList *utils.TokenList) bool {
	tmp := *tokList
	if function(tokList) {
		return true
	}
	*tokList = tmp
	return var_declaration(tokList) && tokList.Match(";")
}

func parameter(tokList *utils.TokenList) bool {
	tmp := *tokList
	var typeTok utils.Token
	var id utils.Token
	if types(tokList) && tokList.Mark(&typeTok, tokList.Cursor()-1) && tokList.IsIdentifier() && tokList.Mark(&id, tokList.Cursor()-1) && parameter2(tokList) {
		nd := semantic.GetSymbolTableNode(id.Content(), semantic.KeywordShortTermToFullTerm[typeTok.Content()], nil, semantic.Scope, 0)
		semantic.CurrentSymbolTable.Insert(&nd)
		return true
	}
	*tokList = tmp
	return true
}

func parameter2(tokList *utils.TokenList) bool {
	tmp := *tokList
	var typeTok utils.Token
	var id utils.Token
	if tokList.Match(",") && types(tokList) && tokList.Mark(&typeTok, tokList.Cursor()-1) && tokList.IsIdentifier() && tokList.Mark(&id, tokList.Cursor()-1) && parameter2(tokList) {
		nd := semantic.GetSymbolTableNode(id.Content(), semantic.KeywordShortTermToFullTerm[typeTok.Content()], nil, semantic.Scope, 0)
		semantic.CurrentSymbolTable.Insert(&nd)
		return true
	}
	*tokList = tmp
	return true
}

func function(tokList *utils.TokenList) bool {
	initScope := semantic.Scope
	var typeTok utils.Token
	var id utils.Token
	t1 := types(tokList) && tokList.Mark(&typeTok, tokList.Cursor()-1) && tokList.IsIdentifier() && tokList.Mark(&id, tokList.Cursor()-1) && tokList.Match("(")
	if t1 {
		semantic.IncreaseScope()
		functionID := utils.RandString(8)                                                                // now the function is represented with an ID
		semantic.FunctionReturnType[functionID] = semantic.KeywordShortTermToFullTerm[typeTok.Content()] // save the return type
		functionSymbolTable := semantic.GetSymbolTable()
		functionSymbolTable.SetParent(semantic.CurrentSymbolTable)
		functionNode := semantic.GetSymbolTableNode(functionID, "funct", &functionSymbolTable, semantic.Scope, -1)
		semantic.CurrentSymbolTable.Insert(&functionNode)
		semantic.CurrentSymbolTable = &functionSymbolTable // pass current symbol table to new function symbol table
		t2 := parameter(tokList)
		if t2 {
			t3 := tokList.Match(")") && tokList.Match("{") && block_statement(tokList) && return_statement(tokList, functionID) && tokList.Match("}") && semantic.DecreaseScope()
			if t3 {
				semantic.CurrentSymbolTable = semantic.CurrentSymbolTable.Parent() // return the symbol table to parent symbol table
				return true
			}
		}
	}
	semantic.Scope = initScope
	return false
}

func types(tokList *utils.TokenList) bool {
	tmp := *tokList
	types := []string{"int", "double", "char", "string"}
	for _, tp := range types {
		if tokList.Match(tp) {
			return true
		}
		*tokList = tmp
	}
	return false
}

func type_ep(tokList *utils.TokenList) bool {
	tmp := *tokList
	if tokList.IsInt() {
		return true
	}
	*tokList = tmp
	if tokList.IsChar() {
		return true
	}
	*tokList = tmp
	if tokList.IsString() {
		return true
	}
	*tokList = tmp
	return tokList.IsDouble()
}

func block_statement_(tokList *utils.TokenList) bool {
	tmp := *tokList
	if statement(tokList) {
		block_statement_(tokList)
	} else {
		*tokList = tmp
	}
	return true
}

func block_statement(tokList *utils.TokenList) bool {
	return statement(tokList) && block_statement_(tokList)
}

// LL(1) parsing
func statement(tokList *utils.TokenList) bool {
	tmp := *tokList
	if types(tokList) {
		*tokList = tmp
		return var_declaration(tokList) && tokList.Match(";")
	}
	*tokList = tmp
	if tokList.IsIdentifier() {
		*tokList = tmp
		return assignment(tokList) && tokList.Match(";")
	}
	*tokList = tmp
	if tokList.Match("if") {
		*tokList = tmp
		return if_statement(tokList)
	}
	*tokList = tmp
	if tokList.Match("for") {
		*tokList = tmp
		return for_statement(tokList)
	}
	*tokList = tmp
	if tokList.Match("while") {
		*tokList = tmp
		return while_statement(tokList)
	}
	return false
}

// SR(0) parsing
func return_statement(tokList *utils.TokenList, functionID string) bool {
	return return_statement_sr(tokList, functionID)
}

func return_statement_recursive(tokList *utils.TokenList) bool {
	tmp := *tokList
	if tokList.Match("return") && type_ep(tokList) && tokList.Match(";") {
		return true
	}
	*tokList = tmp
	return tokList.Match("return") && tokList.IsIdentifier() && tokList.Match(";")
}

func return_statement_sr(tokList *utils.TokenList, functionID string) bool {
	stack := make([]utils.Token, 0)
	state := 0
	for 0 <= state && state <= 3 {
		tok := tokList.Pop()
		switch state {
		case 0:
			if tok.Match("return") {
				stack = append(stack, tok)
				state = 1
			} else {
				state = -1
			}
		case 1:
			if tok.IsIdentifier() {
				prevTok := tokList.PrevToken()
				retTok := semantic.CurrentSymbolTable.Find(prevTok.Content())
				if retTok == nil {
					fmt.Printf(utils.RedString("%s is not declared before using it\n"), prevTok.Content())
					return false
				}
				if retTok.Type != semantic.FunctionReturnType[functionID] {
					fmt.Printf(utils.RedString("Return type is not correct\n"))
					return false
				}
				stack = append(stack, tok)
				state = 2
			} else if tok.Type_ep() {
				stack = append(stack, tok)
				state = 3
			} else {
				state = -1
			}
		case 2:
			if tok.Match(";") {
				stack = append(stack, tok)
				state = 4
			} else {
				state = -1
			}
		case 3:
			if tok.Match(";") {
				stack = append(stack, tok)
				state = 5
			} else {
				state = -1
			}
		}
	}
	return state == 4 || state == 5
}

func var_declaration(tokList *utils.TokenList) bool {
	var tp utils.Token
	tmp := *tokList
	if types(tokList) && tokList.Mark(&tp, tokList.Cursor()-1) && assignment(tokList) {
		semantic.Sstack.Push(tp)
		semantic.Sstack.Push(utils.GetToken("d+", "declaration", -1))
		semantic.Action()
		return true
	}
	*tokList = tmp
	if types(tokList) && tokList.Mark(&tp, tokList.Cursor()-1) && tokList.IsIdentifier() {
		id := tokList.PrevToken()
		nd := semantic.GetSymbolTableNode(id.Content(), semantic.KeywordShortTermToFullTerm[tp.Content()], nil, semantic.Scope, -1)
		semantic.CurrentSymbolTable.Insert(&nd)
		return true
	}
	return false
}

func assignment(tokList *utils.TokenList) bool {
	tmp := *tokList
	tmpQ := *semantic.Qstack
	tmpS := *semantic.Sstack
	var id utils.Token
	if tokList.IsIdentifier() && tokList.Mark(&id, tokList.Cursor()-1) && tokList.Match("=") && expressions(tokList) {
		semantic.Sstack.Push(id)
		semantic.Sstack.Push(utils.GetToken("=", "operator", -1))
		return semantic.Action()
	}
	*tokList = tmp
	*semantic.Qstack = tmpQ
	*semantic.Sstack = tmpS
	if tokList.IsIdentifier() && tokList.Mark(&id, tokList.Cursor()-1) && tokList.Match("=") && tokList.IsIdentifier() {
		semantic.Sstack.Push(tokList.PrevToken())
		semantic.Sstack.Push(id)
		semantic.Sstack.Push(utils.GetToken("=", "operator", -1))
		return semantic.Action()
	} else {
		return false
	}
}
