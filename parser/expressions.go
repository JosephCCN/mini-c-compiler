package parser

import (
	"github.com/utils"
)

func expressions(tokList *utils.TokenList) bool {
	tmp := tokList.Copy()
	if math_experssion(tokList) {
		return true
	}
	tokList = tmp.Copy()
	if logic_experssion(tokList) {
		return true
	}
	return tokList.IsString()
}

func math_experssion(tokList *utils.TokenList) bool {
	tmp := tokList.Copy()
	if math_experssion2(tokList) {
		return true
	}
	tokList = tmp.Copy()
	return math_experssion(tokList) && op1(tokList) && math_experssion2(tokList)
}

func math_experssion2(tokList *utils.TokenList) bool {
	tmp := tokList.Copy()
	if term(tokList) {
		return true
	}
	tokList = tmp.Copy()
	return math_experssion2(tokList) && op2(tokList) && term(tokList)
}

func logic_experssion(tokList *utils.TokenList) bool {
	tmp := tokList.Copy()
	if tokList.Match("!") && logic_experssion(tokList) {
		return true
	}
	tokList = tmp.Copy()
	if logic_experssion2(tokList) {
		return true
	}
	tokList = tmp.Copy()
	return logic_experssion(tokList) && bool_op(tokList) && logic_experssion2(tokList)
}

func logic_experssion2(tokList *utils.TokenList) bool {
	tmp := tokList.Copy()
	if logic_experssion3(tokList) {
		return true
	}
	tokList = tmp.Copy()
	return logic_experssion2(tokList) && cmp_op(tokList) && logic_experssion3(tokList)
}

func logic_experssion3(tokList *utils.TokenList) bool {
	tmp := tokList.Copy()
	if math_experssion(tokList) {
		return true
	}
	tokList = tmp.Copy()
	if logic_experssion3(tokList) && bitwise_op(tokList) && term(tokList) {
		return true
	}
	tokList = tmp.Copy()
	return term(tokList)
}

func term(tokList *utils.TokenList) bool {
	tmp := tokList.Copy()
	if tokList.IsIdentifier() {
		return true
	}
	tokList = tmp.Copy()
	if tokList.IsInt() {
		return true
	}
	tokList = tmp.Copy()
	if tokList.IsDouble() {
		return true
	}
	tokList = tmp.Copy()
	return tokList.Match("(") && math_experssion(tokList) && tokList.Match(")")
}

func op1(tokList *utils.TokenList) bool {
	tmp := tokList.Copy()
	if tokList.Match("+") {
		return true
	}
	tokList = tmp.Copy()
	return tokList.Match("-")
}

func op2(tokList *utils.TokenList) bool {
	tmp := tokList.Copy()
	if tokList.Match("*") {
		return true
	}
	tokList = tmp.Copy()
	return tokList.Match("/")
}

func bool_op(tokList *utils.TokenList) bool {
	tmp := tokList.Copy()
	if tokList.Match("&&") {
		return true
	}
	tokList = tmp.Copy()
	return tokList.Match("||")
}

func cmp_op(tokList *utils.TokenList) bool {
	tmp := tokList.Copy()
	types := []string{">", "<", ">=", "<=", "=="}
	for _, tp := range types {
		if tokList.Match(tp) {
			return true
		}
		tokList = tmp.Copy()
	}
	return false
}

func bitwise_op(tokList *utils.TokenList) bool {
	tmp := tokList.Copy()
	types := []string{"&", "|", "^", "~"}
	for _, tp := range types {
		if tokList.Match(tp) {
			return true
		}
		tokList = tmp.Copy()
	}
	return false
}
