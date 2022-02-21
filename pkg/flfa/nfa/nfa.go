package nfa

import (
	"fmt"
	"strconv"
)

type args struct {
	str string
}

type State string
type Symbol rune
type Delta map[State]map[Symbol]StatesBitMap
type StatesBitMap uint64

type nfa struct {
	states          []State
	alphabet        []Symbol
	delta           Delta
	startingStates  StatesBitMap
	acceptingStates StatesBitMap
}

/*
  Creates an empty NFA.
*/
func initializeNFA() nfa {
	return nfa{
		[]State([]State(nil)),
		[]Symbol([]Symbol(nil)),
		Delta(Delta(nil)),
		StatesBitMap(0),
		StatesBitMap(0),
	}
}

/*
  Creates an NFA and validates it.
  If the NFA fails validation, then an empty NFA is returned.
*/
func NewNFA(states []State, alphabet []Symbol, delta Delta, startingStates StatesBitMap, acceptingStates StatesBitMap) (nfa, error) {
	nfa := nfa{states, alphabet, delta, startingStates, acceptingStates}

	err := nfa.validate()
	if err != nil {
		return initializeNFA(), nil
	}

	return nfa, nil
}

/*
  Validates and solves an NFA given a string using parallel bit mapping.
	If the NFA fails validation, then an empty NFA is returned.
	If the given string contains a symbol not in the language, then the current state and false is returned.
*/
func (nfa *nfa) Solve(str string) (StatesBitMap, bool, error) {
	currentStates := nfa.startingStates

	err := nfa.validate()
	if err != nil {
		return currentStates, false, err
	}

	for _, symbol := range str {
		err := nfa.validateSymbol(Symbol(symbol))
		if err != nil {
			return currentStates, false, err
		}

		nextStates := StatesBitMap(0)

		for i := 0; currentStates != 0; i++ {
			if currentStates%2 == 1 {
				nextStates |= nfa.delta[nfa.states[i]][Symbol(symbol)]
			}

			currentStates >>= 1
		}

		currentStates = nextStates
	}

	isAccepting := currentStates&nfa.acceptingStates != 0

	return currentStates, isAccepting, nil
}

/*
	Validates the entire NFA.
*/
func (nfa *nfa) validate() error {
	err := nfa.validateDelta()
	if err != nil {
		return err
	}

	err = nfa.validateStartingState()
	if err != nil {
		return err
	}

	err = nfa.validateAcceptingStates()
	if err != nil {
		return err
	}

	return nil
}

/*
	Validates the NFA's delta.
*/
func (nfa *nfa) validateDelta() error {
	// The last error catches if delta has less states and pinpoints it
	if len(nfa.delta) > len(nfa.states) {
		return fmt.Errorf("delta contains too many states")
	}

	i := 0
	for _, state := range nfa.states {
		if nfa.states[i] != state {
			return fmt.Errorf("the states in the states array and delta map must be in the same order, states %v does not match delta %v", nfa.states[i], state)
		}

		if _, ok := nfa.delta[state]; ok {
			// The last error catches if delta has less transitions and pinpoints it
			if len(nfa.delta[state]) > len(nfa.alphabet) {
				return fmt.Errorf("delta contains too many transitions for the state '%v'", state)
			}
		}

		for _, symbol := range nfa.alphabet {
			if newState, ok := nfa.delta[state][symbol]; ok {
				err := checkStates(nfa.states, newState, args{str: "new"})
				if err != nil {
					return err
				}

			} else {
				return fmt.Errorf("delta is not defined for the state '%v' and the symbol '%v'", state, string(symbol))
			}
		}

		i++
	}

	return nil
}

/*
	Validates the NFA's starting state.
*/
func (nfa *nfa) validateStartingState() error {
	err := checkStates(nfa.states, nfa.startingStates, args{str: "starting"})

	return err
}

/*
	Validates the NFA's accepting states.
*/
func (nfa *nfa) validateAcceptingStates() error {
	err := checkStates(nfa.states, nfa.acceptingStates, args{str: "accepting"})

	return err
}

/*
	Validates a given symbol against the NFA's alphabet.
*/
func (nfa *nfa) validateSymbol(symbol Symbol) error {
	for _, acceptedSymbol := range nfa.alphabet {
		if acceptedSymbol == symbol {
			return nil
		}
	}

	return fmt.Errorf("the symbol '%v' is not within the alphabet", string(symbol))
}

/*
	Checks if a state is in a state array.
*/
func checkStates(states []State, statesBitMap StatesBitMap, args args) error {
	if len(states) >= len(strconv.FormatUint(uint64(statesBitMap), 2)) {
		return nil
	}

	return fmt.Errorf("the %v states bit map '%v' is too long", args.str, statesBitMap)
}
