package dfa

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDFA(t *testing.T) {
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
			fmt.Errorf("delta is not defined for the state 'q0' and the symbol 'a'"),
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
			fmt.Errorf("delta is not defined for the state 'q0' and the symbol 'a'"),
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
			fmt.Errorf("the new state 'q3' is not within the possible states"),
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
			fmt.Errorf("the starting state 'q3' is not within the possible states"),
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
			fmt.Errorf("the accepting state 'q3' is not within the possible states"),
		},
	}

	for _, tt := range tests {
		dfa, err := NewDFA(tt.dfa.states, tt.dfa.alphabet, tt.dfa.delta, tt.dfa.startingState, tt.dfa.acceptingStates)

		assert.Equal(t, tt.want, err)

		if err != nil {
			emptyDFA := InitializeDFA()
			assert.Equal(t, emptyDFA, dfa)
		} else {
			assert.Equal(t, tt.dfa, dfa)
		}
	}
}

type input struct {
	language string
}

type output struct {
	finalState  State
	isAccepting bool
	err         error
}

type dfaSolveWant struct {
	input  input
	output output
}

func TestDFASolve(t *testing.T) {
	var tests = []struct {
		dfa  dfa
		want []dfaSolveWant
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
			[]dfaSolveWant{
				{input{""}, output{"q0", false, nil}},
				{input{"b"}, output{"q1", true, nil}},
				{input{"a"}, output{"q1", true, nil}},
				{input{"abbbaaa"}, output{"q1", true, nil}},
			},
		},
	}

	for _, tt := range tests {
		err := tt.dfa.Validate()
		assert.Equal(t, nil, err)

		for _, dfaSolve := range tt.want {
			input := dfaSolve.input
			output := dfaSolve.output

			finalState, isAccepting, err := tt.dfa.Solve(input.language)

			assert.Equal(t, output.finalState, finalState)
			assert.Equal(t, output.isAccepting, isAccepting)
			assert.Equal(t, output.err, err)
		}
	}
}
