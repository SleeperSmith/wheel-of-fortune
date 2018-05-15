package main

import (
	"testing"
)

func TestShuffle(t *testing.T) {

	// Arrange
	testlist := map[string]int{
		"a": 1, "b": 2, "c": 1,
	}

	// Act
	result := Shuffle(testlist)

	// Assert
	if len(result) != 4 {
		t.Fail()
	}
	// Check all cards are present
	rc := map[string]int{
		"a": 0, "b": 0, "c": 0,
	}
	for _, cardid := range result {
		for k := range rc {
			if k == cardid {
				rc[k]++
			}
		}
	}
	if rc["a"] != 1 || rc["b"] != 2 || rc["c"] != 1 {
		t.Log("Result card count is wrong")
	}
}

type dummyStrategy struct {
}

func (ds dummyStrategy) MainPhase(previousState *PlayerState) []PlayerState {
	if len(previousState.boardUntapped) == 0 {
		playA := previousState.Clone()
		playA.PlayCard("a", false)
		playB := previousState.Clone()
		playB.PlayCard("b", false)
		result := []PlayerState{
			*playA, *playB,
		}
		return result
	}
	result := []PlayerState{
		*previousState,
	}
	return result
}

func (ds dummyStrategy) DiscardPhase(previousState *PlayerState) []PlayerState {
	result := []PlayerState{
		*previousState,
	}

	return result
}

func TestPlay(t *testing.T) {

	// Arrange
	testlist := []string{
		"a", "b", "c", "a", "b", "c", "a", "b", "c", "a", "b", "c",
	}
	ti := engine{
		ps: dummyStrategy{},
		as: nil,
	}

	// Act
	result := ti.PlayPermutations(testlist, 3)

	// Assert
	if len(result) != 4 {
		t.Log("Unexpected player state mutation result")
		t.Fail()
	}
}
