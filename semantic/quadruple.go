package semantic

import "fmt"

type quadruple struct {
	op   string
	var1 string
	var2 string
	res  string
}

type quadrupleStack struct {
	stack []quadruple
}

func GetQuadruple(op string, var1 string, var2 string, res string) quadruple {
	ret := quadruple{
		op:   op,
		var1: var1,
		var2: var2,
		res:  res,
	}
	return ret
}

func GetQuadrupleStack() quadrupleStack {
	ret := quadrupleStack{
		stack: make([]quadruple, 0),
	}
	return ret
}

func (s *quadrupleStack) Push(q quadruple) {
	s.stack = append(s.stack, q)
}

func (s *quadrupleStack) Pop() quadruple {
	q := s.Top()
	if len(s.stack) >= 1 {
		s.stack = s.stack[:len(s.stack)-1]
	}
	return q
}

func (s *quadrupleStack) Top() quadruple {
	if len(s.stack) >= 1 {
		return s.stack[len(s.stack)-1]
	}
	return quadruple{}
}

func (s *quadrupleStack) ListAll() string {
	ret := ""
	for _, q := range s.stack {
		if q.op == "label" {
			ret += fmt.Sprintf("%s:\n", q.res)
		} else {
			ret += fmt.Sprintf("(%s, %s, %s, %s)\n", q.op, q.var1, q.var2, q.res)
		}
	}
	return ret
}
