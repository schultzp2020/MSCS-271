package main

import (
	"flfa/dfa"
	"fmt"
)

func main() {
	states := []dfa.State{"q0", "q1", "q2", "q3", "q4", "q5", "q6", "q7", "q8", "q9", "q10"}
	alphabet := []dfa.Symbol{'m', 'w', 'g', 'c'}
	delta := dfa.Delta{
		"q0": {
			'm': "q10",
			'w': "q10",
			'g': "q1",
			'c': "q10",
		},
		"q1": {
			'm': "q2",
			'w': "q10",
			'g': "q0",
			'c': "q10",
		},
		"q2": {
			'm': "q1",
			'w': "q5",
			'g': "q10",
			'c': "q3",
		},
		"q3": {
			'm': "q10",
			'w': "q10",
			'g': "q4",
			'c': "q2",
		},
		"q4": {
			'm': "q10",
			'w': "q7",
			'g': "q3",
			'c': "q10",
		},
		"q5": {
			'm': "q10",
			'w': "q2",
			'g': "q6",
			'c': "q10",
		},
		"q6": {
			'm': "q10",
			'w': "q10",
			'g': "q5",
			'c': "q7",
		},
		"q7": {
			'm': "q8",
			'w': "q4",
			'g': "q10",
			'c': "q6",
		},
		"q8": {
			'm': "q7",
			'w': "q10",
			'g': "q9",
			'c': "q10",
		},
		"q9": {
			'm': "q10",
			'w': "q10",
			'g': "q8",
			'c': "q10",
		},
		"q10": {
			'm': "q10",
			'w': "q10",
			'g': "q10",
			'c': "q10",
		},
	}
	startingState := dfa.State("q0")
	acceptingStates := []dfa.State{"q9"}

	dfa, err := dfa.NewDFA(states, alphabet, delta, startingState, acceptingStates)
	if err != nil {
		fmt.Print(err)
		return
	}

	var str string

	fmt.Println("Use the following letters: \nm - man \nw - wolf \ng - goat \nc - cabbage")
	fmt.Print("What is your man, wolf, goat, cabbage guess? ")
	fmt.Scan(&str)

	finalState, isAccepted, err := dfa.Solve(str)
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Printf("Final State: %v\nIs Accepted: %v\n", finalState, isAccepted)
}
