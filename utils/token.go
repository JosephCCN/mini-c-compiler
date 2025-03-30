package utils

type Token struct {
	content string
	id      int
	tokType string
}

func GetToken(content string, tokType string) Token {
	return Token{
		content: content,
		id:      -1,
		tokType: tokType,
	}
}

func (tok Token) GetContent() string {
	return tok.content
}

func (tok Token) Match(content string) bool {
	return tok.content == content
}

func (tok Token) IsInt() bool {
	return tok.tokType == "integer"
}

func (tok Token) IsString() bool {
	return tok.tokType == "string"
}

func (tok Token) IsChar() bool {
	return tok.tokType == "character"
}
