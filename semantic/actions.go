package semantic

import (
	"fmt"

	"github.com/utils"
)

func Action() bool {
	op := Sstack.Pop()

	switch op.Content() {
	case "+", "-", "*", "/", "&", "|", "^", "~", "&&", "||", ">", "<", ">=", "<=", "==":
		v2 := Sstack.Pop()
		v1 := Sstack.Pop()
		return operator(v1, v2, op.Content())
	case "=":
		lvalue := Sstack.Pop()
		rvalue := Sstack.Pop()
		return assignemnt(lvalue, rvalue)
	case "d+":
		tp := Sstack.Pop()
		v := Sstack.Pop()
		return declaration(v, tp)
	case "if":
		v := Sstack.Pop()
		return if_statement(v)
	default:
		break
	}

	return false
}

func getType(v utils.Token, recursive bool) string {
	if v.Type() == "identifier" {
		node := CurrentSymbolTable.Find(v.Content())
		if node != nil {
			return node.Type
		}
		if recursive {
			if CurrentSymbolTable.Parent() == CurrentSymbolTable {
				return ""
			}
			tmp := CurrentSymbolTable
			CurrentSymbolTable = CurrentSymbolTable.Parent()
			res := getType(v, recursive)
			CurrentSymbolTable = tmp
			return res
		} else {
			return ""
		}
	} else if v.Type() == "keyword" {
		return KeywordShortTermToFullTerm[v.Content()]
	}
	return v.Type()
}

func typeConversion(t1 string, t2 string, op string) (string, error) {
	res := TypeConvert[op][t1][t2]
	if res != "" {
		return res, nil
	}
	return "", fmt.Errorf("Operator %s cannot apply on %s and %s", op, t1, t2)
}

func if_statement(v utils.Token) bool {
	Qstack.Push(GetQuadruple("if", v.Content(), "", fmt.Sprintf("L%d", NextLable)))
	Qstack.Push(GetQuadruple("goto", "", "", fmt.Sprintf("L%d", NextLable+1)))
	Qstack.Push(GetQuadruple("label", "", "", fmt.Sprintf("L%d", NextLable)))
	NextLable += 2
	return true
}

func operator(v1 utils.Token, v2 utils.Token, op string) bool {
	v1Type := getType(v1, true)
	v2Type := getType(v2, true)

	if v1Type == "" {
		fmt.Println(utils.RedString(fmt.Sprintf("%s is used before declaration", v1.Content())))
		return false
	}
	if v2Type == "" {
		fmt.Println(utils.RedString(fmt.Sprintf("%s is used before declaration", v2.Content())))
		return false
	}

	convertedType, err := typeConversion(v1Type, v2Type, op)
	if err != nil {
		fmt.Println(utils.RedString(err.Error()))
		return false
		// os.Exit(0)
	}
	tok := utils.GetToken(fmt.Sprintf("t%d", nextT), convertedType, v1.Line())
	q := GetQuadruple(op, v1.Content(), v2.Content(), tok.Content())
	nextT += 1
	Qstack.Push(q)
	Sstack.Push(tok)
	return true
}

func declaration(v utils.Token, tp utils.Token) bool {
	tpType := getType(tp, false)
	_, err := typeConversion(tpType, v.Type(), "+")
	if err != nil {
		fmt.Println(utils.RedString("Wrong type assignment"))
		return false
	}
	node := GetSymbolTableNode(v.Content(), tpType, nil, Scope, TypeSize[tpType])
	return CurrentSymbolTable.Insert(&node)
}

func assignemnt(l utils.Token, r utils.Token) bool {
	lType := getType(l, false)
	rType := getType(r, false)
	if lType == "" { // lvalue is not declared, assume this assignment is declaration
		Sstack.Push(utils.GetToken(l.Content(), rType, l.Line())) //replace the type of lvalue from identifier to type of right value
		q := GetQuadruple("=", l.Content(), "", r.Content())
		Qstack.Push(q)
	} else {
		_, err := typeConversion(lType, rType, "+") // l = r can be seen as l = l + r, where l is equivilant to 0
		if err != nil {
			fmt.Println("equal:", utils.RedString(err.Error()))
			return false
			// os.Exit(0)
		}
	}
	return true
}
