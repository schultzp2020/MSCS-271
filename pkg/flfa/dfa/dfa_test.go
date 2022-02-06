package dfa

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDFAConstructor(t *testing.T) {
	dfaNil := dfa{
		[]State([]State(nil)),
		[]Symbol([]Symbol(nil)),
		Delta(Delta(nil)),
		State(""),
		[]State([]State(nil)),
	}

	var tests = []struct {
		dfa  dfa
		want error
	}{
		{
			dfa{
				[]State{"q0", "q1"},
				[]Symbol{'a', 'b'},
				Delta{
					"q0": {
						'a': "q1",
						'b': "q1",
					},
					"q1": {
						'a': "q1",
						'b': "q1",
					},
				},
				State("q0"),
				[]State{"q1"},
			},
			nil,
		},
		{
			dfa{
				[]State{"q0", "q1"},
				[]Symbol{'a', 'b'},
				Delta{
					"q3": {
						'a': "q1",
						'b': "q1",
					},
					"q1": {
						'a': "q1",
						'b': "q1",
					},
				},
				State("q0"),
				[]State{"q1"},
			},
			errors.New("delta is not defined for the state 'q0' and the symbol 'a'"),
		},
		{
			dfa{
				[]State{"q0", "q1"},
				[]Symbol{'a', 'b'},
				Delta{
					"q0": {
						'c': "q1",
						'b': "q1",
					},
					"q1": {
						'a': "q1",
						'b': "q1",
					},
				},
				State("q0"),
				[]State{"q1"},
			},
			errors.New("delta is not defined for the state 'q0' and the symbol 'a'"),
		},
		{
			dfa{
				[]State{"q0", "q1"},
				[]Symbol{'a', 'b'},
				Delta{
					"q0": {
						'a': "q3",
						'b': "q1",
					},
					"q1": {
						'a': "q1",
						'b': "q1",
					},
				},
				State("q0"),
				[]State{"q1"},
			},
			errors.New("the new state 'q3' is not within the possible states"),
		},
		{
			dfa{
				[]State{"q0", "q1"},
				[]Symbol{'a', 'b'},
				Delta{
					"q0": {
						'a': "q1",
						'b': "q1",
					},
					"q1": {
						'a': "q1",
						'b': "q1",
					},
				},
				State("q3"),
				[]State{"q1"},
			},
			errors.New("the starting state 'q3' is not within the possible states"),
		},
		{
			dfa{
				[]State{"q0", "q1"},
				[]Symbol{'a', 'b'},
				Delta{
					"q0": {
						'a': "q1",
						'b': "q1",
					},
					"q1": {
						'a': "q1",
						'b': "q1",
					},
				},
				State("q0"),
				[]State{"q3"},
			},
			errors.New("the accepting state 'q3' is not within the possible states"),
		},
	}

	for _, tt := range tests {
		dfa, err := New(tt.dfa.states, tt.dfa.alphabet, tt.dfa.delta, tt.dfa.startingState, tt.dfa.acceptingStates)

		assert.Equal(t, tt.want, err)

		if err != nil {
			assert.Equal(t, dfaNil, dfa)
		} else {
			assert.Equal(t, tt.dfa, dfa)
		}
	}
}
