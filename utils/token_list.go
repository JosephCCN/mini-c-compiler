package utils

var TokenListDeepest int

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

func (self *TokenList) Mark(tok *Token, pos int) bool {
	if pos < len(self.list) {
		*tok = self.list[pos]
		return true
	}
	return false
}

func (self *TokenList) At(pos int) Token {
	if pos < len(self.list) {
		return self.list[pos]
	}
	return Token{}
}

func (self *TokenList) PrevToken() Token {
	if self.cur == 0 {
		return Token{}
	}
	return self.list[self.cur-1]
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
	TokenListDeepest = max(self.cur, TokenListDeepest)
	return self.list[self.cur-1]
}

func (self *TokenList) GetToken(pos int) Token {
	return self.list[pos]
}

func (self *TokenList) List() []Token {
	return self.list
}

func (self *TokenList) Cursor() int {
	return self.cur
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
