package main

import (
	"testing"
)

func newTestState() PlayerState {
	return PlayerState{
		hand:          []string{"a", "b", "b"},
		boardTapped:   make([]string, 0),
		boardUntapped: make([]string, 0),
	}
}

func TestPlayStateCardTapped(t *testing.T) {

	// Arrange
	ts := newTestState()

	// Act
	ts.PlayCard("a", true)

	// Assert
	if len(ts.hand) != 2 || ts.hand[0] != "b" {
		t.Log("Card in hand is wrong.")
		t.Fail()
	}
	if len(ts.boardTapped) != 1 || ts.boardTapped[0] != "a" {
		t.Log("Card played tapped is wrong.")
		t.Fail()
	}
	if len(ts.boardUntapped) != 0 {
		t.Log("Card played untapped is wrong.")
		t.Fail()
	}
}

func TestPlayStateCardUntapped(t *testing.T) {

	// Arrange
	ts := newTestState()

	// Act
	ts.PlayCard("b", false)

	// Assert
	if len(ts.boardUntapped) != 1 {
		t.Log("Only one card should be played at a time..")
		t.Fail()
	}
	if ts.boardUntapped[0] != "b" {
		t.Log("Card played tapped is wrong.")
		t.Fail()
	}
}

func TestDraw(t *testing.T) {

	// Arrange
	ts := PlayerState{
		library:       []string{"a", "b"},
		hand:          make([]string, 0),
		boardTapped:   make([]string, 0),
		boardUntapped: make([]string, 0),
	}

	// Act
	ts.DrawCard()

	// Assert
	if len(ts.hand) != 1 || ts.hand[0] != "a" {
		t.Log("Card in hand is wrong.")
		t.Fail()
	}
	if len(ts.library) != 1 || ts.library[0] != "b" {
		t.Log("Cards in library is wrong.")
		t.Fail()
	}
}

func TestUntap(t *testing.T) {

	// Arrange
	ts := PlayerState{
		boardTapped:   []string{"a", "b"},
		boardUntapped: make([]string, 0),
	}

	// Act
	ts.Untap()

	// Assert
	if len(ts.boardTapped) != 0 {
		t.Log("Stuff's still tapped.")
		t.Fail()
	}
	if len(ts.boardUntapped) != 2 {
		t.Log("Cards in untapped is wrong.")
		t.Fail()
	}
}
