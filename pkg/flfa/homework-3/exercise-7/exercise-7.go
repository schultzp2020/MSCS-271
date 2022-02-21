package main

import (
	"flfa/nfa"
	"fmt"
)

func main() {
	states := []nfa.State{"q0", "q1", "q2"}
	alphabet := []nfa.Symbol{'a', 'b'}
	delta := nfa.Delta{
		"q0": {
			'a': 0b011,
			'b': 0b001,
		},
		"q1": {
			'a': 0b100,
			'b': 0b000,
		},
		"q2": {
			'a': 0b100,
			'b': 0b100,
		},
	}
	startingStates := nfa.StatesBitMap(0b001)
	acceptingStates := nfa.StatesBitMap(0b100)

	nfa, err := nfa.NewNFA(states, alphabet, delta, startingStates, acceptingStates)
	if err != nil {
		fmt.Print(err)
		return
	}

	var str string

	fmt.Print("What is your string? ")
	fmt.Scan(&str)

	finalStates, isAccepting, err := nfa.Solve(str)
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Printf("Final States: %b\nIs Accepted: %v\n", finalStates, isAccepting)
}
