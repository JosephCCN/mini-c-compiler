package lexical

func GetToken(name string, id int) Token {
	return Token{
		name: name,
		id:   id,
	}
}

func Run(src string) []Token {
	var (
		tokenList []Token
	)
	tokenList = append(tokenList, GetToken("test", 1))

	return tokenList
}
