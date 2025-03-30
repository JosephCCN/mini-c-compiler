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
			"integer":   1,
			"character": 2,
			"string":    3,
		},
		boundId: map[string]int{
			"integer":   1,
			"character": 2,
			"string":    3,
		},
		allowIncrement: map[string]bool{
			"integer":   false,
			"character": false,
			"string":    false,
		},
	}
}

func (self TokenTable) Append(tok *Token) (bool, error) {
	// Append will not append token that already inside the table
	if self.contain(*tok) {
		return false, errors.New("Token already inside table")
	}

	_, err := self.assignIdToToken(tok) //assign id to token
	if err != nil {
		return false, err
	}

	self.table[tok.tokType] = append(self.table[tok.tokType], *tok)
	return true, nil
}

func (self TokenTable) assignIdToToken(tok *Token) (bool, error) {
	if self.contain(*tok) {
		return false, errors.New("Token already inside table")
	}

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
		if t == tok {
			return true
		}
	}
	return false
}
