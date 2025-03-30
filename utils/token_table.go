package utils

import (
	"errors"
)

type TokenTable struct {
	table          map[string][]Token
	currentId      map[string]int
	boundId        map[string]int
	allowIncrement map[string]bool
}

func GetTokenTable() TokenTable {
	return TokenTable{
		table: make(map[string][]Token),
		currentId: map[string]int{
			"identifier": 0,
			"integer":    1,
			"character":  2,
			"string":     3,
			"double":     4,
			"keyword":    5,
			"operator":   24,
			"punc":       44,
		},
		boundId: map[string]int{
			"identifier": 0,
			"integer":    1,
			"character":  2,
			"string":     3,
			"double":     4,
			"keyword":    23,
			"operator":   43,
			"punc":       64,
		},
		allowIncrement: map[string]bool{
			"identifier": false,
			"integer":    false,
			"character":  false,
			"string":     false,
			"double":     false,
			"keyword":    true,
			"operator":   true,
			"punc":       true,
		},
	}
}

func (self TokenTable) Append(tok *Token) (bool, error) {
	// Append will not append token that already inside the table
	if self.contain(*tok) {
		tok.id = self.findTokenId(*tok)
		return true, nil
	}

	_, err := self.assignIdToToken(tok) //assign id to token
	if err != nil {
		return false, err
	}

	self.table[tok.tokType] = append(self.table[tok.tokType], *tok)
	return true, nil
}

func (self TokenTable) assignIdToToken(tok *Token) (bool, error) {
	if !self.allowIncrement[tok.tokType] {
		tok.id = self.currentId[tok.tokType]
	} else {
		if self.currentId[tok.tokType] >= self.boundId[tok.tokType] {
			return false, errors.New("Token ID out of bound")
		}
		tok.id = self.currentId[tok.tokType]
		self.currentId[tok.tokType]++
	}

	return true, nil
}

func (self TokenTable) contain(tok Token) bool {
	for _, t := range self.table[tok.tokType] {
		if t.content == tok.content {
			return true
		}
	}
	return false
}

func (self TokenTable) findTokenId(tok Token) int {
	for _, t := range self.table[tok.tokType] {
		if t.content == tok.content {
			return t.id
		}
	}
	return -1
}
