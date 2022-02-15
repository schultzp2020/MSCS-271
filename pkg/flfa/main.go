package main

import (
	"flfa/nfa"
	"fmt"
)

func main() {
	states := []nfa.State{"q0", "q1"}
	alphabet := []nfa.Symbol{'0', '1'}
	delta := nfa.Delta{
		"q0": {
			'0': 0b01,
			'1': 0b10,
		},
		"q1": {
			'0': 0b01,
			'1': 0b10,
		},
	}
	startingStates := nfa.StatesBitMap(0b01)
	acceptingStates := nfa.StatesBitMap(0b10)

	nfa, err := nfa.NewNFA(states, alphabet, delta, startingStates, acceptingStates)
	if err != nil {
		fmt.Print(err)
	}

	finalStates, _, _ := nfa.Solve("111")

	fmt.Printf("%b", finalStates)
}
