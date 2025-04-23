package lexical

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/utils"
)

type Token utils.Token

func findShortestMatch(lst []string) string {
	if lst == nil {
		return ""
	}
	sort.Strings(lst)
	return lst[0]
}

func Run(src string) (utils.TokenList, error) {

	table := utils.GetTokenTable()
	tokenList := utils.GetTokenList()

	regex := map[string]*regexp.Regexp{
		"integer":    regexp.MustCompile("^[+-]?[0-9]+"),
		"double":     regexp.MustCompile("^[-+]?[0-9]*\\.?[0-9]+"),
		"character":  regexp.MustCompile("^'.'"),
		"string":     regexp.MustCompile("^\".*\"?"),
		"identifier": regexp.MustCompile("^[_a-zA-Z][_a-zA-Z0-9]*"),
		"keyword":    regexp.MustCompile("^int|^char|^string|^main|^for|^while|^else if|^if|^else|^return|^include|^define"),
		"operator":   regexp.MustCompile(`^=|^-|^\+|^\*|^/|^>=|^<=|^>|^<|^==|^&&|^\|\||^!`),
		"punc":       regexp.MustCompile(`^{|^}|^;|^\(|^\)|^,|^\[|^\]`),
	}
	order := []string{"keyword", "identifier", "operator", "double", "integer", "string", "character", "punc"}

	for len(src) > 0 {
		src = strings.TrimSpace(src)
		matched := false

		for _, Type := range order {
			if regex[Type].MatchString(src) {
				//content := findShortestMatch(regex[Type].FindStringSubmatch(src))
				content := findShortestMatch(regex[Type].FindStringSubmatch(src))
				src = strings.TrimPrefix(src, content)

				tok := utils.GetToken(content, Type)
				_, err := table.Append(&tok)
				if err != nil {
					return tokenList, err
				}
				tokenList.Push(tok)
				matched = true
				break
			}
		}

		if !matched {
			return tokenList, fmt.Errorf("cannot analysis following token: %s", src)
		}

	}
	return tokenList, nil

}
