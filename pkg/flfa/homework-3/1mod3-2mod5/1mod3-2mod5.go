package main

import (
	"flfa/nfa"
	"fmt"
)

func main() {
	states := []nfa.State{"q0", "q1", "q2", "q3", "q4", "q5", "q6", "q7", "q8"}
	alphabet := []nfa.Symbol{'0', '1'}
	delta := nfa.Delta{
		"q0": {
			'0': 0b000010010,
			'1': 0b000110111,
		},
		"q1": {
			'0': 0b000000010,
			'1': 0b000010111,
		},
		"q2": {
			'0': 0b000011011,
			'1': 0b000100111,
		},
		"q3": {
			'0': 0b000010111,
			'1': 0b000001000,
		},
		"q4": {
			'0': 0b000010000,
			'1': 0b000100000,
		},
		"q5": {
			'0': 0b001010011,
			'1': 0b010000000,
		},
		"q6": {
			'0': 0b100010011,
			'1': 0b000110111,
		},
		"q7": {
			'0': 0b000100000,
			'1': 0b001010011,
		},
		"q8": {
			'0': 0b010000000,
			'1': 0b100000000,
		},
	}
	startingStates := nfa.StatesBitMap(0b000000001)
	acceptingStates := nfa.StatesBitMap(0b001000100)

	nfa, err := nfa.NewNFA(states, alphabet, delta, startingStates, acceptingStates)
	if err != nil {
		fmt.Print(err)
		return
	}

	var str string

	fmt.Print("What is your binary string? ")
	fmt.Scan(&str)

	finalStates, isAccepting, err := nfa.Solve(str)
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Printf("Final States: %b\nIs Accepted: %v\n", finalStates, isAccepting)
}
