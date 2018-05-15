package deckstructure

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
func Shuffle(decklist map[string]int) []string {
	total := 0
	for _, cardcount := range decklist {
		total += cardcount
	}

	randomOrder := ri.Perm(total)
	result := make([]string, total, total)
	gi := 0
	for cardid, cardcount := range decklist {
		for li := 0; li < cardcount; li++ {
			result[randomOrder[gi]] = cardid
			gi++
		}
	}

	return result
}
