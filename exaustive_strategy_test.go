package main

import (
	"testing"
)

func TestNoDisard(t *testing.T) {
	testStrategyInstance := exaustiveStrat{}
	noDiscardPs := PlayerState{
		hand:      []string{"a", "a", "a", "a", "a", "a", "a"},
		graveyard: []string{},
	}

	result := testStrategyInstance.DiscardPhase(&noDiscardPs)

	if len(result) != 1 {
		t.Log("Mutation of state occured at discard step with hand less than 7.")
		t.Fail()
	}

	if len(result[0].hand) != 7 || len(result[0].graveyard) != 0 {
		t.Log("Card were moved around unecessarily")
		t.Fail()
	}
}

func TestDisard(t *testing.T) {
	testStrategyInstance := exaustiveStrat{}
	noDiscardPs := PlayerState{
		hand:      []string{"a", "a", "a", "a", "a", "a", "b", "c", "d"},
		graveyard: []string{},
	}

	result := testStrategyInstance.DiscardPhase(&noDiscardPs)

	// combinations:
	// aa, ab, ac, ad, bc, bd, cd,
	if len(result) != 7 {
		t.Log("Wrong number of discard combinations")
		t.Fail()
	}

	if len(result[0].hand) != 7 || len(result[0].graveyard) != 2 {
		t.Log("Wrong number of cards discarded / kept")
		t.Fail()
	}
}
