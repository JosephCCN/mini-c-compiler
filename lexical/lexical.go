package lexical

import (
	"fmt"
	"regexp"
	"strings"
)

func regexCompile(reg string) *regexp.Regexp {
	ret, _ := regexp.Compile(reg)
	return ret
}

func Run(src string) ([]Token, error) {

	table := GetTokenTable()
	tokenList := make([]Token, 0)

	regex := map[string]*regexp.Regexp{
		"integer": regexCompile("[+-]?[0-9]+"),
	}
	order := []string{"integer"}

	for len(src) > 0 {
		src = strings.TrimSpace(src)
		matched := false

		for _, Type := range order {
			if regex[Type].MatchString(src) {
				content := regex[Type].FindString(src)
				src = strings.TrimPrefix(src, content)

				tok := GetToken(content, Type)
				_, err := table.Append(&tok)
				if err != nil {
					return tokenList, err
				}
				tokenList = append(tokenList, tok)
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
