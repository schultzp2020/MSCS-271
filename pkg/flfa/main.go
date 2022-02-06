package main

import (
	"flfa/dfa"
	"fmt"
)

func main() {
	states := []dfa.State{"q0", "q1"}
	alphabet := []dfa.Symbol{'a', 'b'}
	delta := dfa.Delta{
		"q0": {
			'a': "q1",
			'b': "q1",
		},
		"q1": {
			'a': "q1",
			'b': "q1",
		},
	}
	startingState := dfa.State("q0")
	acceptingStates := []dfa.State{"q1"}

	dfa, err := dfa.New(states, alphabet, delta, startingState, acceptingStates)
	if err != nil {
		fmt.Print(err)
		return
	}

	finalState, isAccepted, err := dfa.Solve("b")
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Printf("Final State: %v\nIs Accepted: %v", finalState, isAccepted)
}
