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

/*
  Creates an empty DFA.
*/
func InitializeDFA() dfa {
	return dfa{
		[]State([]State(nil)),
		[]Symbol([]Symbol(nil)),
		Delta(Delta(nil)),
		State(""),
		[]State([]State(nil)),
	}
}

/*
  Creates a DFA and validates it.
  If the DFA fails validation, then an empty DFA is returned.
*/
func NewDFA(states []State, alphabet []Symbol, delta Delta, startingState State, acceptingStates []State) (dfa, error) {
	dfa := dfa{states, alphabet, delta, startingState, acceptingStates}

	err := dfa.Validate()
	if err != nil {
		emptyDFA := InitializeDFA()
		return emptyDFA, err
	}

	return dfa, nil
}

/*
  Validates and solves a DFA given a string.
	If the DFA fails validation, then an empty DFA is returned.
	If the given string contains a symbol not in the language, then the current state and false is returned.
*/
func (dfa *dfa) Solve(str string) (State, bool, error) {
	err := dfa.Validate()
	if err != nil {
		return "", false, err
	}

	state := dfa.startingState

	for _, symbol := range str {
		err := dfa.ValidateSymbol(Symbol(symbol))
		if err != nil {
			return state, false, err
		}

		state = dfa.delta[state][Symbol(symbol)]
	}

	ok := dfa.IsStateAccepting(state)

	return state, ok, nil
}

/*
	Validates the entire DFA.
*/
func (dfa *dfa) Validate() error {
	err := dfa.ValidateDelta()
	if err != nil {
		return err
	}

	err = dfa.ValidateStartingState()
	if err != nil {
		return err
	}

	err = dfa.ValidateAcceptingStates()
	if err != nil {
		return err
	}

	return nil
}

/*
	Validates the DFA's delta.
*/
func (dfa *dfa) ValidateDelta() error {
	// The last error catches if delta has less states and pinpoints it
	if len(dfa.delta) > len(dfa.states) {
		return fmt.Errorf("delta contains too many states")
	}

	for _, state := range dfa.states {
		if _, ok := dfa.delta[state]; ok {
			// The last error catches if delta has less transitions and pinpoints it
			if len(dfa.delta[state]) > len(dfa.alphabet) {
				return fmt.Errorf("delta contains too many transitions for the state '%v'", state)
			}
		}

		for _, symbol := range dfa.alphabet {
			if newState, ok := dfa.delta[state][symbol]; ok {
				err := checkStateInStates(dfa.states, newState, args{str: "new"})
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

/*
	Validates the DFA's starting state.
*/
func (dfa *dfa) ValidateStartingState() error {
	err := checkStateInStates(dfa.states, dfa.startingState, args{str: "starting"})

	return err
}

/*
	Validates the DFA's accepting states.
*/
func (dfa *dfa) ValidateAcceptingStates() error {
	err := checkStatesInStates(dfa.states, dfa.acceptingStates, args{str: "accepting"})

	return err
}

/*
	Validates a given symbol against the DFA's alphabet.
*/
func (dfa *dfa) ValidateSymbol(symbol Symbol) error {
	for _, acceptedSymbol := range dfa.alphabet {
		if acceptedSymbol == symbol {
			return nil
		}
	}

	return fmt.Errorf("the symbol '%v' is not within the alphabet", string(symbol))
}

/*
	Checks if a state is an accepting state.
*/
func (dfa *dfa) IsStateAccepting(state State) bool {
	err := checkStateInStates(dfa.acceptingStates, state, args{})

	return err == nil
}

/*
	Checks if a state is in a state array.
*/
func checkStateInStates(states []State, state State, args args) error {
	for _, possibleState := range states {
		if possibleState == state {
			return nil
		}
	}

	return fmt.Errorf("the %v state '%v' is not within the possible states", args.str, state)
}

/*
	Checks if a state array is a subset of another state array.
*/
func checkStatesInStates(allowedStates []State, states []State, args args) error {
	for _, state := range states {
		err := checkStateInStates(allowedStates, state, args)
		if err != nil {
			return err
		}
	}

	return nil
}
