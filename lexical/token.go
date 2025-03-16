package lexical

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
