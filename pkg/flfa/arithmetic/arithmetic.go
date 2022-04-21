package main

import (
	"flfa/ll1"
	"fmt"
)

func main() {
	nonterminalAlphabet := []rune{'S', 't', 'T', 'r', 'R', 'v'}
	terminalAlphabet := []rune{'d', '+', '-', '*', '/', '(', ')', '$'}
	startSymbol := 'S'
	ll1Table := ll1.LL1Table{
		'S': {
			'(': {'t', '$'},
			'd': {'t', '$'},
		},
		't': {
			'(': {'r', 'T'},
			'd': {'r', 'T'},
		},
		'T': {
			'+': {'+', 'r', 'T'},
			'-': {'-', 'r', 'T'},
			')': {},
			'$': {},
		},
		'r': {
			'(': {'v', 'R'},
			'd': {'v', 'R'},
		},
		'R': {
			'+': {},
			'-': {},
			'*': {'*', 'v', 'R'},
			'/': {'/', 'v', 'R'},
			')': {},
			'$': {},
		},
		'v': {
			'(': {'(', 't', ')'},
			'd': {'d'},
		},
	}
	regexesReplaces := []ll1.RegexReplace{
		{Regex: "\\d+\\.?\\d*|\\.\\d+", Replacement: "d"},
	}

	ll1, err := ll1.NewLL1(nonterminalAlphabet, terminalAlphabet, startSymbol, ll1Table, regexesReplaces)
	if err != nil {
		fmt.Print(err)
		return
	}
	var str string

	fmt.Println("Enter an arithmetic expression.")
	fmt.Scan(&str)

	for _, char := range str {
		if char == '$' || char == 'd' {
			fmt.Println("the string entered contains either the character '$' or 'd' which cannot be used")
			return
		}
	}

	err = ll1.Solve(str + "$")
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Printf("the string '%v' is valid", str)
}
