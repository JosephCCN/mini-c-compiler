package utils

type Token struct {
	content string
	id      int
	tokType string
	line    int
}

func GetToken(content string, tokType string, line int) Token {
	return Token{
		content: content,
		id:      -1,
		tokType: tokType,
		line:    line,
	}
}

func (tok Token) Getcontent() string {
	return tok.content
}

func (tok Token) GetLine() int {
	return tok.line
}

func (tok Token) Match(content string) bool {
	return tok.content == content
}

func (tok Token) IsInt() bool {
	return tok.tokType == "integer"
}

func (tok Token) IsDouble() bool {
	return tok.tokType == "double"
}

func (tok Token) IsString() bool {
	return tok.tokType == "string"
}

func (tok Token) IsChar() bool {
	return tok.tokType == "character"
}

func (tok Token) IsIdentifier() bool {
	return tok.tokType == "identifier"
}

func (tok Token) Type_ep() bool {
	types := []string{"integer", "double", "string", "character"}
	for _, tp := range types {
		if tok.tokType == tp {
			return true
		}
	}
	return false
}
