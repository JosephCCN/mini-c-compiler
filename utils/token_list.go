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

func (self *TokenList) Copy() *TokenList {
	ret := &TokenList{
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
