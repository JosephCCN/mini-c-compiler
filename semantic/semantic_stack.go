package semantic

import (
	"fmt"

	"github.com/utils"
)

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

func (s *SemanticStack) Pop() utils.Token {
	var ret utils.Token
	if len(s.stack) >= 1 {
		ret = s.Top()
		s.stack = s.stack[:len(s.stack)-1]
	}
	return ret
}

func (s *SemanticStack) Top() utils.Token {
	if len(s.stack) >= 1 {
		return s.stack[len(s.stack)-1]
	}
	return utils.Token{}
}

func (s *SemanticStack) ListAll() string {
	ret := ""
	for _, q := range s.stack {
		ret += fmt.Sprintf("[%s, %s, %s]\n", q.Content(), q.Type(), q.ID())
	}
	return ret
}
