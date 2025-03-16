package lexical

type TokenTable struct {
	table map[string][]Token
}

func GetTokenTable() TokenTable {
	return TokenTable{
		table: make(map[string][]Token),
	}
}

func (self TokenTable) Append(tokType string, tok Token) bool {
	for _, t := range self.table[tokType] {
		if t == tok {
			return false
		}
	}
	self.table[tokType] = append(self.table[tokType], tok)
	return true
}
