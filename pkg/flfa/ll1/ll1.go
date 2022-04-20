package ll1

import (
	"regexp"
)

type LL1Table map[rune]map[rune][]rune

type RegexReplace struct {
	regex       string
	replacement string
}

type compiledRegexReplace struct {
	regex       *regexp.Regexp
	replacement string
}

type ll1 struct {
	nonterminalAlphabet []rune
	terminalAlphabet    []rune
	startSymbol         rune
	ll1Table            LL1Table
	regexesReplaces     []compiledRegexReplace
}

func initializeLL1() ll1 {
	return ll1{
		[]rune([]rune(nil)),
		[]rune([]rune(nil)),
		rune('S'),
		LL1Table(LL1Table(nil)),
		[]compiledRegexReplace([]compiledRegexReplace(nil)),
	}
}

func NewLL1(nonterminalAlphabet []rune, terminalAlphabet []rune, startSymbol rune, ll1Table LL1Table, regexesReplaces []RegexReplace) (ll1, error) {
	var compiledRegexesReplaces []compiledRegexReplace

	for _, regexReplace := range regexesReplaces {
		compiledRegex, err := regexp.Compile(regexReplace.regex)
		if err != nil {
			ll1 := initializeLL1()
			return ll1, err
		}

		compiledRegexesReplaces = append(compiledRegexesReplaces, compiledRegexReplace{compiledRegex, regexReplace.replacement})
	}

	ll1 := ll1{nonterminalAlphabet, terminalAlphabet, startSymbol, ll1Table, compiledRegexesReplaces}

	return ll1, nil
}

func (ll1 *ll1) Solve(str string) error {
	for _, regexReplace := range ll1.regexesReplaces {
		str = regexReplace.regex.ReplaceAllString(str, regexReplace.replacement)
	}

	stack := []rune{ll1.startSymbol}

	for _, strChar := range str {
		stackChar, stack := stack[0], stack[1:]
		for strChar != stackChar {
			if pushRunes, ok := ll1.ll1Table[strChar][stackChar]; ok {
				stack = append(pushRunes, stack...)
				// Potential out of index error
				stackChar, stack = stack[0], stack[1:]
			}
		}
	}

	return nil
}
