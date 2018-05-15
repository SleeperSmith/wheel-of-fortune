package deckstructure

import (
	"testing"
)

func TestShuffle(t *testing.T) {
	testlist := map[string]int{
		"a": 1, "b": 2, "c": 1,
	}

	result := Shuffle(testlist)

	if len(result) != 4 {
		t.Fail()
	}
}
