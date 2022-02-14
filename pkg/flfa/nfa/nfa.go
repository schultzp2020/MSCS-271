package nfa

type args struct {
	str string
}

type State string
type Symbol rune
type Delta map[State]map[Symbol]string

type nfa struct {
	states          []State
	alphabet        []Symbol
	delta           Delta
	startingStates  []State
	acceptingStates []State
}

/*
  Creates an empty NFA.
*/
func initializeNFA() nfa {
	return nfa{
		[]State([]State(nil)),
		[]Symbol([]Symbol(nil)),
		Delta(Delta(nil)),
		[]State([]State(nil)),
		[]State([]State(nil)),
	}
}
