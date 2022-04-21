package ll1

import (
	"fmt"
	"regexp"
)

type LL1Table map[string]map[string][]string

type RegexReplace struct {
	Regex       string
	Replacement string
}

type compiledRegexReplace struct {
	regex       *regexp.Regexp
	replacement string
}

type ll1 struct {
	nonterminalAlphabet []string
	terminalAlphabet    []string
	startSymbol         string
	ll1Table            LL1Table
	regexesReplaces     []compiledRegexReplace
}

func initializeLL1() ll1 {
	return ll1{
		[]string([]string(nil)),
		[]string([]string(nil)),
		string(""),
		LL1Table(LL1Table(nil)),
		[]compiledRegexReplace([]compiledRegexReplace(nil)),
	}
}

func NewLL1(nonterminalAlphabet []string, terminalAlphabet []string, startSymbol string, ll1Table LL1Table, regexesReplaces []RegexReplace) (ll1, error) {
	var compiledRegexesReplaces []compiledRegexReplace

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

func (ll1 *ll1) Solve(str string) error {
	err := ll1.validate()
	if err != nil {
		return err
	}

	for _, regexReplace := range ll1.regexesReplaces {
		str = regexReplace.regex.ReplaceAllString(str, regexReplace.replacement)
	}

	stack := []string{ll1.startSymbol}

	for i, char := range str {
		err := ll1.checkChar(string(char))
		if err != nil {
			return err
		}

		for len(stack) != 0 && string(char) != stack[0] {
			if pushString, ok := ll1.ll1Table[stack[0]][string(char)]; ok {
				stack = stack[1:]
				stack = append(pushString, stack...)
			} else {
				return fmt.Errorf("there is push string while reading '%v' and popping '%v' at the position '%v ~error~ %v'", string(char), stack[0], str[:i], str[i:])
			}
		}

		if len(stack) != 0 && string(char) == stack[0] {
			stack = stack[1:]
		}

		if len(stack) == 0 && i+1 != len([]rune(str)) {
			return fmt.Errorf("there is nothing to pop while reading '%v' at the position '%v ~error~ %v'", string(char), str[:i+1], str[i+1:])
		}
	}

	return nil
}

func (ll1 *ll1) validate() error {
	err := checkNonterminalAlphabet(ll1.nonterminalAlphabet)
	if err != nil {
		return err
	}

	err = checkTerminalAlphabet(ll1.terminalAlphabet)
	if err != nil {
		return err
	}

	err = checkStartSymbol(ll1.startSymbol, ll1.nonterminalAlphabet)
	if err != nil {
		return err
	}

	return nil
}

func (ll1 *ll1) checkChar(char string) error {
	for _, terminalChar := range ll1.terminalAlphabet {
		if char == string(terminalChar) {
			return nil
		}
	}

	return fmt.Errorf("character '%v' is not a terminal character", char)
}

func checkTerminalAlphabet(terminalAlphabet []string) error {
	for _, terminal := range terminalAlphabet {
		if len([]rune(terminal)) != 1 {
			return fmt.Errorf("the terminal character '%v' is more than one character", terminal)
		}
	}

	return nil
}

func checkNonterminalAlphabet(nonterminalAlphabet []string) error {
	for _, nonterminal := range nonterminalAlphabet {
		if len([]rune(nonterminal)) != 1 {
			return fmt.Errorf("the terminal character '%v' is more than one character", nonterminal)
		}
	}

	return nil
}

func checkStartSymbol(startSymbol string, nonterminalAlphabet []string) error {
	for _, nonterminal := range nonterminalAlphabet {
		if startSymbol == nonterminal {
			return nil
		}
	}
	return fmt.Errorf("the starting symbol '%v' is not in the nonterminal alphabet", startSymbol)
}
