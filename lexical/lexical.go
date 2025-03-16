package lexical

func GetToken(content string, id int) Token {
	return Token{
		content: content,
		id:      id,
	}
}

func Run(src string) []Token {

	tokTable := GetTokenTable()
	tok := GetToken("test", 2)
	tokTable.Append("int", tok)

	var (
		tokenList []Token
	)
	tokenList = append(tokenList, GetToken("test", 1))

	return tokenList
}
