package parser

import (
	"github.com/semantic"
	"github.com/utils"
)

func expressions(tokList *utils.TokenList) bool {
	tmp := *tokList
	if logic_experssion(tokList) {
		return true
	}

	*tokList = tmp
	return term(tokList)
}

func math_experssion_(tokList *utils.TokenList) bool {
	tmp := *tokList
	tmpQ := *semantic.Qstack
	tmpS := *semantic.Sstack
	var op utils.Token
	t := op1(tokList) && tokList.Mark(&op, tokList.Cursor()-1) && math_experssion2(tokList)
	if t {
		semantic.Sstack.Push(op)
		return semantic.Action() && math_experssion_(tokList)
	} else {
		*semantic.Qstack = tmpQ
		*semantic.Sstack = tmpS
		*tokList = tmp
	}
	return true
}

func math_experssion(tokList *utils.TokenList) bool {
	return math_experssion2(tokList) && math_experssion_(tokList)
}

func math_experssion2_(tokList *utils.TokenList) bool {
	tmp := *tokList
	tmpQ := *semantic.Qstack
	tmpS := *semantic.Sstack
	var op utils.Token
	t := op2(tokList) && tokList.Mark(&op, tokList.Cursor()-1) && term(tokList)
	if t {
		semantic.Sstack.Push(op)
		return semantic.Action() && math_experssion2_(tokList)
	} else {
		*semantic.Qstack = tmpQ
		*semantic.Sstack = tmpS
		*tokList = tmp
	}
	return true
}

func math_experssion2(tokList *utils.TokenList) bool {
	return term(tokList) && math_experssion2_(tokList)
}

func logic_experssion_(tokList *utils.TokenList) bool {
	tmp := *tokList
	tmpQ := *semantic.Qstack
	tmpS := *semantic.Sstack
	var op utils.Token
	t := bool_op(tokList) && tokList.Mark(&op, tokList.Cursor()-1) && logic_experssion2(tokList)
	if t {
		semantic.Sstack.Push(op)
		semantic.Action()
		logic_experssion_(tokList)
	} else {
		*semantic.Qstack = tmpQ
		*semantic.Sstack = tmpS
		*tokList = tmp
	}
	return true
}

func logic_experssion(tokList *utils.TokenList) bool {
	tmp := *tokList
	tmpQ := *semantic.Qstack
	tmpS := *semantic.Sstack
	if logic_experssion2(tokList) && logic_experssion_(tokList) {
		return true
	}
	*tokList = tmp
	*semantic.Qstack = tmpQ
	*semantic.Sstack = tmpS
	var op utils.Token
	if tokList.Match("!") && tokList.Mark(&op, tokList.Cursor()-1) && logic_experssion(tokList) && logic_experssion_(tokList) {
		semantic.Sstack.Push(op)
		return semantic.Action()
	}
	return false
}

func logic_experssion2_(tokList *utils.TokenList) bool {
	tmp := *tokList
	tmpQ := *semantic.Qstack
	tmpS := *semantic.Sstack
	var op utils.Token
	t := bitwise_op(tokList) && tokList.Mark(&op, tokList.Cursor()-1) && logic_experssion3(tokList)
	if t {
		semantic.Sstack.Push(op)
		return semantic.Action() && logic_experssion2_(tokList)
	} else {
		*semantic.Qstack = tmpQ
		*semantic.Sstack = tmpS
		*tokList = tmp
	}
	return true
}

func logic_experssion2(tokList *utils.TokenList) bool {
	return logic_experssion3(tokList) && logic_experssion2_(tokList)
}

func logic_experssion3_(tokList *utils.TokenList) bool {
	tmp := *tokList
	tmpQ := *semantic.Qstack
	tmpS := *semantic.Sstack
	var op utils.Token
	t := cmp_op(tokList) && tokList.Mark(&op, tokList.Cursor()-1) && math_experssion(tokList)
	if t {
		semantic.Sstack.Push(op)
		return semantic.Action() && logic_experssion3_(tokList)
	} else {
		*semantic.Qstack = tmpQ
		*semantic.Sstack = tmpS
		*tokList = tmp
	}
	return true
}

func logic_experssion3(tokList *utils.TokenList) bool {
	return math_experssion(tokList) && logic_experssion3_(tokList)
}

func term(tokList *utils.TokenList) bool {
	tmp := *tokList
	if tokList.IsIdentifier() {
		semantic.Sstack.Push(tokList.PrevToken())
		return true
	}
	*tokList = tmp
	if tokList.IsInt() {
		semantic.Sstack.Push(tokList.PrevToken())
		return true
	}
	*tokList = tmp
	if tokList.IsDouble() {
		semantic.Sstack.Push(tokList.PrevToken())
		return true
	}
	*tokList = tmp
	if tokList.IsString() {
		semantic.Sstack.Push(tokList.PrevToken())
		return true
	}
	*tokList = tmp
	if tokList.IsChar() {
		semantic.Sstack.Push(tokList.PrevToken())
		return true
	}
	*tokList = tmp
	if tokList.Match("(") && math_experssion(tokList) && tokList.Match(")") {
		return true
	}
	*tokList = tmp
	return false
}

func op1(tokList *utils.TokenList) bool {
	tmp := *tokList
	types := []string{"+", "-"}
	for _, tp := range types {
		if tokList.Match(tp) {
			return true
		}
		*tokList = tmp
	}
	return false
}

func op2(tokList *utils.TokenList) bool {
	tmp := *tokList
	types := []string{"*", "/"}
	for _, tp := range types {
		if tokList.Match(tp) {
			return true
		}
		*tokList = tmp
	}
	return false
}

func bool_op(tokList *utils.TokenList) bool {
	tmp := *tokList
	types := []string{"||", "&&"}
	for _, tp := range types {
		if tokList.Match(tp) {
			return true
		}
		*tokList = tmp
	}
	return false
}

func cmp_op(tokList *utils.TokenList) bool {
	tmp := *tokList
	types := []string{">", "<", ">=", "<=", "=="}
	for _, tp := range types {
		if tokList.Match(tp) {
			return true
		}
		*tokList = tmp
	}
	return false
}

func bitwise_op(tokList *utils.TokenList) bool {
	tmp := *tokList
	types := []string{"&", "|", "^", "~"}
	for _, tp := range types {
		if tokList.Match(tp) {
			return true
		}
		*tokList = tmp
	}
	return false
}
