package main

type exaustiveStrat struct {
	cc cardCatalogue
}
type cardCatalogue struct {
}
type cardInfo struct {
	IsBasic       bool
	ComesInTapped bool
	TapsFor       []string
}

func contains(a []string, itemToCheck string) bool {
	for _, val := range a {
		if val == itemToCheck {
			return true
		}
	}
	return false
}

func (es exaustiveStrat) MainPhase(previousState *PlayerState) []PlayerState {

	result := make([]PlayerState, 0)

	cardsPlayed := make([]string, 0)
	for _, cardFromHand := range previousState.hand {
		liForCard, cardIsLand := es.cc.GetDetail(cardFromHand)

		if !cardIsLand {
			continue
		}
		if contains(cardsPlayed, cardFromHand) {
			continue
		}

		nextState := previousState.Clone()

		nextState.PlayCard(cardFromHand, liForCard.ComesInTapped)
		result = append(result, *nextState)
		cardsPlayed = append(cardsPlayed, cardFromHand)
	}

	if len(result) == 0 {
		result = append(result, *previousState)
	}

	return result
}

func (es exaustiveStrat) DiscardPhase(previousState *PlayerState) []PlayerState {

	result := []PlayerState{
		*previousState,
	}
	return result
}

func (cc cardCatalogue) GetDetail(cardID string) (cardInfo, bool) {
	resultInfo, resultOk := cardlist[cardID]
	return resultInfo, resultOk
}
