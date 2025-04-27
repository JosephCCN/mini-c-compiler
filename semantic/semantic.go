package semantic

var RootSymbolTable *SymbolTable
var Sstack *SemanticStack
var Qstack *quadrupleStack
var Scope int
var nextT int
var NextLable int
var TypeSize map[string]int
var TypeConvert map[string]map[string]map[string]string
var KeywordShortTermToFullTerm map[string]string

func Init() {
	tmpST := GetSymbolTable()
	RootSymbolTable = &tmpST
	tmpSS := GetSemanticStack()
	Sstack = &tmpSS
	tmpQs := GetQuadrupleStack()
	Qstack = &tmpQs
	TypeSize = map[string]int{
		"integer": 4,
		"double":  8,
		"char":    1,
		"string":  0,
	}
	KeywordShortTermToFullTerm = map[string]string{
		"int":    "integer",
		"double": "double",
		"char":   "character",
		"string": "string",
	}
	conversionInit()
	Scope = 0
	nextT = 0
	NextLable = 0
}

func conversionInit() {
	plus := map[string]map[string]string{
		"integer": map[string]string{
			"integer": "integer",
			"double":  "double",
			"char":    "integer",
		},
		"double": map[string]string{
			"integer": "double",
			"double":  "double",
		},
		"char": map[string]string{
			"integer": "integer",
			"char":    "char",
		},
	}
	multiply := map[string]map[string]string{
		"integer": map[string]string{
			"integer": "integer",
			"double":  "double",
		},
		"double": map[string]string{
			"integer": "double",
			"double":  "double",
		},
	}

	TypeConvert = map[string]map[string]map[string]string{
		"+":  plus,
		"-":  plus, // the type conversion rule is the same as addition
		"*":  multiply,
		"/":  multiply, // the type conversion rule is the same as multiplication
		"&&": plus,
		"||": plus,
		"<":  plus,
		">":  plus,
		">=": plus,
		"<=": plus,
		"==": plus,
		"&":  plus,
		"|":  plus,
		"^":  plus,
		"~":  plus,
	}
}
