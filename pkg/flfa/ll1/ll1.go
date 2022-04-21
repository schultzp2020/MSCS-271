package ll1

import (
	"fmt"
	"regexp"
)

type LL1Table map[rune]map[rune][]rune

type RegexReplace struct {
	Regex       string
	Replacement string
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

/*
  Creates an empty ll1.
*/
func initializeLL1() ll1 {
	return ll1{
		[]rune([]rune(nil)),
		[]rune([]rune(nil)),
		rune(' '),
		LL1Table(LL1Table(nil)),
		[]compiledRegexReplace([]compiledRegexReplace(nil)),
	}
}

/*
  Creates a ll1 and validates it.
  If the ll1 fails validation, then an empty ll1 is returned.
*/
func NewLL1(nonterminalAlphabet []rune, terminalAlphabet []rune, startSymbol rune, ll1Table LL1Table, regexesReplaces []RegexReplace) (ll1, error) {
	var compiledRegexesReplaces []compiledRegexReplace

	// Compiles the given regexes
	for _, regexReplace := range regexesReplaces {
		compiledRegex, err := regexp.Compile(regexReplace.Regex)
		if err != nil {
			return initializeLL1(), err
		}

		compiledRegexesReplaces = append(compiledRegexesReplaces, compiledRegexReplace{compiledRegex, regexReplace.Replacement})
	}

	ll1 := ll1{nonterminalAlphabet, terminalAlphabet, startSymbol, ll1Table, compiledRegexesReplaces}

	err := ll1.validate()
	if err != nil {
		return initializeLL1(), err
	}

	return ll1, nil
}

/*
  Validates and solves a ll1 given a string.
*/
func (ll1 *ll1) Solve(str string) error {
	err := ll1.validate()
	if err != nil {
		return err
	}

	for _, regexReplace := range ll1.regexesReplaces {
		str = regexReplace.regex.ReplaceAllString(str, regexReplace.replacement)
	}

	stack := []rune{ll1.startSymbol}

	for i, char := range str {
		err := ll1.validateChar(char)
		if err != nil {
			return err
		}

		for len(stack) != 0 && char != stack[0] {
			if pushString, ok := ll1.ll1Table[stack[0]][char]; ok {
				stack = stack[1:]
				stack = append(pushString, stack...)
			} else {
				return fmt.Errorf("there is push string while reading '%v' and popping '%v' at the position '%v ~error~ %v'", string(char), stack[0], str[:i], str[i:])
			}
		}

		if len(stack) != 0 && char == stack[0] {
			stack = stack[1:]
		}

		if len(stack) == 0 && i+1 != len([]rune(str)) {
			return fmt.Errorf("there is nothing to pop while reading '%v' at the position '%v ~error~ %v'", string(char), str[:i+1], str[i+1:])
		}
	}

	return nil
}

/*
	Validates the ll1.
*/
func (ll1 *ll1) validate() error {
	err := validateStartSymbol(ll1.startSymbol, ll1.nonterminalAlphabet)
	if err != nil {
		return err
	}

	return nil
}

/*
	Validates a given character against the ll1's terminal alphabet.
*/
func (ll1 *ll1) validateChar(char rune) error {
	for _, terminalChar := range ll1.terminalAlphabet {
		if char == terminalChar {
			return nil
		}
	}

	return fmt.Errorf("character '%v' is not a terminal character", char)
}

/*
	Validates the ll1's starting symbol against the ll1's nonterminal alphabet.
*/
func validateStartSymbol(startSymbol rune, nonterminalAlphabet []rune) error {
	for _, nonterminal := range nonterminalAlphabet {
		if startSymbol == nonterminal {
			return nil
		}
	}
	return fmt.Errorf("the starting symbol '%v' is not in the nonterminal alphabet", startSymbol)
}
