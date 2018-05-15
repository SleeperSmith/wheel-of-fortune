package main

import (
	"math/rand"
	"time"
)

var ri *rand.Rand

func init() {
	rs := rand.NewSource(int64(time.Now().Nanosecond()))
	ri = rand.New(rs)
}

// Shuffle function takes in a map of card id x count
// and returns a list of card id representing its shuffled
// order
func Shuffle(deckList map[string]int) []string {
	total := 0
	for _, cardcount := range deckList {
		total += cardcount
	}

	randomOrder := ri.Perm(total)
	result := make([]string, total, total)
	gi := 0
	for cardid, cardcount := range deckList {
		for li := 0; li < cardcount; li++ {
			result[randomOrder[gi]] = cardid
			gi++
		}
	}

	return result
}

type playStrategy interface {
	MainPhase(state *PlayerState) []PlayerState
	DiscardPhase(state *PlayerState) []PlayerState
}
type analyseStrategy interface {
	Upkeep(state *PlayerState)
	PreMain(state *PlayerState)
	PostMain(state *PlayerState)
}
type engine struct {
	ps playStrategy
	as analyseStrategy
}

// PlayPermutations function does a run of the library and returns the analysis result.
func (e *engine) PlayPermutations(orderedDeckList []string, depth int) []PlayerState {

	goFirst := New(orderedDeckList)
	goSecond := New(orderedDeckList)
	goSecond.DrawCard()

	states := []PlayerState{*goFirst, *goSecond}
	for turn := 1; turn <= depth; turn++ {
		nextStates := make([]PlayerState, 0)
		for _, state := range states {

			state.turn = turn
			state.Untap()
			//e.as.Upkeep(&state)
			state.DrawCard()

			//main phase
			//e.as.PreMain(&state)
			for _, beforeDiscard := range e.ps.MainPhase(&state) {
				//e.as.PostMain(&state)
				for _, afterDiscard := range e.ps.DiscardPhase(&beforeDiscard) {
					nextStates = append(nextStates, afterDiscard)
				}
			}
		}
		states = nextStates
	}

	return states
}
