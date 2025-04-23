package parser

import (
	"github.com/utils"
)

func expressions(tokList *utils.TokenList) bool {
	tmp := tokList.ShallowCopy()
	if math_experssion(tokList) {
		return true
	}
	*tokList = tmp.ShallowCopy()
	if logic_experssion(tokList) {
		return true
	}
	*tokList = tmp.ShallowCopy()
	return term(tokList)
}

func math_experssion_(tokList *utils.TokenList) bool {
	tmp := tokList.ShallowCopy()
	t := op1(tokList) && math_experssion2(tokList)
	if t {
		math_experssion_(tokList)
	} else {
		*tokList = tmp.ShallowCopy()
	}
	return true
}

func math_experssion(tokList *utils.TokenList) bool {
	return math_experssion2(tokList) && math_experssion_(tokList)
}

func math_experssion2_(tokList *utils.TokenList) bool {
	tmp := tokList.ShallowCopy()
	t := op2(tokList) && term(tokList)
	if t {
		math_experssion2_(tokList)
	} else {
		*tokList = tmp.ShallowCopy()
	}
	return true
}

func math_experssion2(tokList *utils.TokenList) bool {
	return term(tokList) && math_experssion2_(tokList)
}

func logic_experssion_(tokList *utils.TokenList) bool {
	tmp := tokList.ShallowCopy()
	t := bool_op(tokList) && logic_experssion2(tokList)
	if t {
		logic_experssion_(tokList)
	} else {
		*tokList = tmp.ShallowCopy()
	}
	return true
}

func logic_experssion(tokList *utils.TokenList) bool {
	tmp := tokList.ShallowCopy()
	if logic_experssion2(tokList) && logic_experssion_(tokList) {
		return true
	}
	*tokList = tmp.ShallowCopy()
	return tokList.Match("!") && logic_experssion(tokList) && logic_experssion_(tokList)
}

func logic_experssion2_(tokList *utils.TokenList) bool {
	tmp := tokList.ShallowCopy()
	t := bitwise_op(tokList) && logic_experssion3(tokList)
	if t {
		logic_experssion2_(tokList)
	} else {
		*tokList = tmp.ShallowCopy()
	}
	return true
}

func logic_experssion2(tokList *utils.TokenList) bool {
	return logic_experssion3(tokList) && logic_experssion2_(tokList)
}

func logic_experssion3_(tokList *utils.TokenList) bool {
	tmp := tokList.ShallowCopy()
	t := cmp_op(tokList) && term(tokList)
	if t {
		logic_experssion3_(tokList)
	} else {
		*tokList = tmp.ShallowCopy()
	}
	return true
}

func logic_experssion3(tokList *utils.TokenList) bool {
	return term(tokList) && logic_experssion3_(tokList)
}

func term(tokList *utils.TokenList) bool {
	tmp := tokList.ShallowCopy()
	if tokList.IsIdentifier() {
		return true
	}
	*tokList = tmp.ShallowCopy()
	if tokList.IsInt() {
		return true
	}
	*tokList = tmp.ShallowCopy()
	if tokList.IsDouble() {
		return true
	}
	*tokList = tmp.ShallowCopy()
	if tokList.IsString() {
		return true
	}
	*tokList = tmp.ShallowCopy()
	if tokList.Match("(") && math_experssion(tokList) && tokList.Match(")") {
		return true
	}
	*tokList = tmp.ShallowCopy()
	return false
}

func op1(tokList *utils.TokenList) bool {
	tmp := tokList.ShallowCopy()
	types := []string{"+", "-"}
	for _, tp := range types {
		if tokList.Match(tp) {
			return true
		}
		*tokList = tmp.ShallowCopy()
	}
	return false
}

func op2(tokList *utils.TokenList) bool {
	tmp := tokList.ShallowCopy()
	types := []string{"*", "/"}
	for _, tp := range types {
		if tokList.Match(tp) {
			return true
		}
		*tokList = tmp.ShallowCopy()
	}
	return false
}

func bool_op(tokList *utils.TokenList) bool {
	tmp := tokList.ShallowCopy()
	types := []string{"||", "&&"}
	for _, tp := range types {
		if tokList.Match(tp) {
			return true
		}
		*tokList = tmp.ShallowCopy()
	}
	return false
}

func cmp_op(tokList *utils.TokenList) bool {
	tmp := tokList.ShallowCopy()
	types := []string{">", "<", ">=", "<=", "=="}
	for _, tp := range types {
		if tokList.Match(tp) {
			return true
		}
		*tokList = tmp.ShallowCopy()
	}
	return false
}

func bitwise_op(tokList *utils.TokenList) bool {
	tmp := tokList.ShallowCopy()
	types := []string{"&", "|", "^", "~"}
	for _, tp := range types {
		if tokList.Match(tp) {
			return true
		}
		*tokList = tmp.ShallowCopy()
	}
	return false
}
