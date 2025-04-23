package utils

type TokenList struct {
	list []Token
	cur  int
}

func GetTokenList() TokenList {
	ret := TokenList{
		list: make([]Token, 0),
		cur:  0,
	}
	return ret
}

func (self *TokenList) End() bool {
	return len(self.list) == self.cur
}

func (self *TokenList) Copy() *TokenList {
	ret := &TokenList{
		list: self.list,
		cur:  self.cur,
	}
	return ret
}

func (self *TokenList) ShallowCopy() TokenList {
	ret := TokenList{
		list: self.list,
		cur:  self.cur,
	}
	return ret
}

func (self *TokenList) Push(tok Token) {
	self.list = append(self.list, tok)
}

func (self *TokenList) Pop() Token {
	if self.cur >= len(self.list) {
		return Token{}
	}
	self.cur += 1
	return self.list[self.cur-1]
}

func (self *TokenList) GetList() []Token {
	return self.list
}

func (self *TokenList) Match(content string) bool {
	tok := self.Pop()
	return tok.Match(content)
}

func (self *TokenList) IsInt() bool {
	tok := self.Pop()
	return tok.IsInt()
}

func (self *TokenList) IsDouble() bool {
	tok := self.Pop()
	return tok.IsDouble()
}

func (self *TokenList) IsChar() bool {
	tok := self.Pop()
	return tok.IsChar()
}

func (self *TokenList) IsString() bool {
	tok := self.Pop()
	return tok.IsString()
}

func (self *TokenList) IsIdentifier() bool {
	tok := self.Pop()
	return tok.IsIdentifier()
}
