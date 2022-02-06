package dfa

import (
	"fmt"
)

type dfa struct {
	states          []string
	alphabet        []rune
	delta           map[string]map[rune]string
	startingState   string
	acceptingStates []string
}

func New(states []string, alphabet []rune, delta map[string]map[rune]string, startingState string, acceptingStates []string) (dfa, error) {
	err := isStartingStateInStates(states, startingState)

	if err != nil {
		return dfa{}, err
	}

	return dfa{states, alphabet, delta, startingState, acceptingStates}, nil
}

func (*dfa) Solve(str string) {

}

func isStartingStateInStates(states []string, startingState string) error {
	for _, state := range states {
		if state == startingState {
			return nil
		}
	}

	return fmt.Errorf("The starting state is not within the possible states!")
}
