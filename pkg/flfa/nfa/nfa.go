package nfa

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

func NewNFA(states []State, alphabet []Symbol, delta Delta, startingStates StatesBitMap, acceptingStates StatesBitMap) (nfa, error) {
	nfa := nfa{states, alphabet, delta, startingStates, acceptingStates}

	err := nfa.validate()
	if err != nil {
		return initializeNFA(), nil
	}

	return nfa, nil
}

func (nfa *nfa) Solve(str string) (StatesBitMap, bool, error) {
	states := nfa.startingStates

	for _, symbol := range str {
		newStates := StatesBitMap(0)
		for i, lastBit := 0, states%2; states != 0; i, states, lastBit = i+1, states>>1, states>>1%2 {
			if lastBit == 1 {
				newStates = nfa.delta[nfa.states[i]][Symbol(symbol)] | newStates
			}
		}
		states = newStates
	}

	return states, true, nil
}

func (nfa *nfa) validate() error {
	return nil
}
