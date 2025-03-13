package lexical

type Token struct {
	name string
	id   int
}

type TokenTable struct {
	table map[string][]Token
}
