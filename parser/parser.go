package parser

import (
	"github.com/utils"
)

// this is the only export function
func Start(tokList *utils.TokenList) bool {
	tmp := tokList.Copy()
	if ex_declaration(tokList) {
		return true
	}
	tokList = tmp.Copy()
	if Start(tokList) && ex_declaration(tokList) {
		return true
	}
	return false
}

func ex_declaration(tokList *utils.TokenList) bool {
	tok := tokList.Pop()
	return tok.IsInt()
}
