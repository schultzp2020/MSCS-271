package dfa

import (
	"fmt"
)

type args struct {
	str string
}

type State string
type Symbol rune
type Delta map[State]map[Symbol]State

type dfa struct {
	states          []State
	alphabet        []Symbol
	delta           Delta
	startingState   State
	acceptingStates []State
}

func New(states []State, alphabet []Symbol, delta Delta, startingState State, acceptingStates []State) (dfa, error) {
	err := isDeltaComplete(states, alphabet, delta)
	if err != nil {
		return dfa{}, err
	}

	err = isStateInStates(states, startingState, args{str: "starting"})
	if err != nil {
		return dfa{}, err
	}

	err = areStatesInStates(states, acceptingStates, args{str: "accepting"})
	if err != nil {
		return dfa{}, err
	}

	return dfa{states, alphabet, delta, startingState, acceptingStates}, nil
}

func (dfa *dfa) Solve(str string) (State, bool, error) {
	state := dfa.startingState

	for _, symbol := range str {
		err := validateSymbol(dfa.alphabet, symbol)
		if err != nil {
			return "", false, err
		}

		state = dfa.delta[state][Symbol(symbol)]
	}

	err := isStateInStates(dfa.acceptingStates, state, args{})

	return state, err == nil, nil
}

func isStateInStates(states []State, state State, args args) error {
	for _, possibleState := range states {
		if possibleState == state {
			return nil
		}
	}

	return fmt.Errorf("the %v state '%v' is not within the possible states", args.str, state)
}

func isDeltaComplete(states []State, alphabet []Symbol, delta Delta) error {
	for _, state := range states {
		for _, symbol := range alphabet {
			if newState, ok := delta[state][symbol]; ok {
				err := isStateInStates(states, newState, args{str: "new"})
				if err != nil {
					return err
				}

			} else {
				return fmt.Errorf("delta is not defined for the state '%v' and the symbol '%v'", state, string(symbol))
			}
		}
	}

	return nil
}

func areStatesInStates(allowedStates []State, states []State, args args) error {
	for _, state := range states {
		err := isStateInStates(allowedStates, state, args)
		if err != nil {
			return err
		}
	}

	return nil
}

func validateSymbol(alphabet []Symbol, symbol rune) error {
	for _, acceptedSymbol := range alphabet {
		if rune(acceptedSymbol) == symbol {
			return nil
		}
	}

	return fmt.Errorf("the symbol '%v' is not within the alphabet", symbol)
}
