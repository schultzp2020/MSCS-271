package main

import (
	"flfa/dfa"
	"fmt"
)

func main() {
	states := []dfa.State{"q0", "q1", "q2", "q3", "q4", "q5", "q6"}
	alphabet := []dfa.Symbol{'0', '1'}
	delta := dfa.Delta{
		"q0": {
			'0': "q0",
			'1': "q1",
		},
		"q1": {
			'0': "q2",
			'1': "q3",
		},
		"q2": {
			'0': "q4",
			'1': "q5",
		},
		"q3": {
			'0': "q6",
			'1': "q0",
		},
		"q4": {
			'0': "q1",
			'1': "q2",
		},
		"q5": {
			'0': "q3",
			'1': "q4",
		},
		"q6": {
			'0': "q5",
			'1': "q6",
		},
	}
	startingState := dfa.State("q0")
	acceptingStates := []dfa.State{"q2"}

	dfa, err := dfa.NewDFA(states, alphabet, delta, startingState, acceptingStates)
	if err != nil {
		fmt.Print(err)
		return
	}

	var str string

	fmt.Print("What is your binary string? ")
	fmt.Scan(&str)

	finalState, isAccepted, err := dfa.Solve(str)
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Printf("Final State: %v\nIs Accepted: %v\n", finalState, isAccepted)
}
