package semantic

import "github.com/utils"

type SemanticStack struct {
	stack []utils.Token
}

func GetSemanticStack() SemanticStack {
	ret := SemanticStack{
		stack: make([]utils.Token, 0),
	}
	return ret
}

func (s *SemanticStack) Push(tok utils.Token) {
	s.stack = append(s.stack, tok)
}

func (s *SemanticStack) Pop() {
	if len(s.stack) >= 1 {
		s.stack = s.stack[:len(s.stack)-1]
	}
}

func (s *SemanticStack) Top() utils.Token {
	if len(s.stack) >= 1 {
		return s.stack[len(s.stack)-1]
	}
	return utils.Token{}
}
