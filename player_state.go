package main

// PlayerState represent the state of a player
type PlayerState struct {
	hand          []string
	library       []string
	boardTapped   []string
	boardUntapped []string
	graveyard     []string
	exile         []string
	turn          int
}

// New function creates a new state from a shuffled deck
func New(orderedDeckList []string) *PlayerState {
	s := &PlayerState{
		hand:          make([]string, 0),
		library:       make([]string, 0),
		boardTapped:   make([]string, 0),
		boardUntapped: make([]string, 0),
		graveyard:     make([]string, 0),
		exile:         make([]string, 0),
		turn:          1,
	}
	s.hand = orderedDeckList[0:5]
	s.library = orderedDeckList[6:]
	return s
}

// Clone creates a deepcopy of the state object.
func (s *PlayerState) Clone() *PlayerState {
	newState := &PlayerState{
		hand:          make([]string, len(s.hand)),
		boardTapped:   make([]string, len(s.boardTapped)),
		boardUntapped: make([]string, len(s.boardUntapped)),
		exile:         make([]string, len(s.exile)),
		graveyard:     make([]string, len(s.graveyard)),
		library:       make([]string, len(s.library)),
		turn:          s.turn,
	}
	copy(newState.hand, s.hand)
	copy(newState.boardTapped, s.boardTapped)
	copy(newState.boardUntapped, s.boardUntapped)
	copy(newState.exile, s.exile)
	copy(newState.library, s.library)
	return newState
}

// DrawCard function draws a card and puts it in player's hand
func (s *PlayerState) DrawCard() {
	s.hand = append(s.hand, s.library[0])
	s.library = s.library[1:]
}

// Untap function untaps all permanents
func (s *PlayerState) Untap() {
	s.boardUntapped = append(s.boardUntapped, s.boardTapped...)
	s.boardTapped = make([]string, 0)
}

// PlayCard plays a card from the hand to the board.
func (s *PlayerState) PlayCard(cardID string, tapped bool) {
	for i, card := range s.hand {
		if card == cardID {
			if tapped {
				s.boardTapped = append(s.boardTapped, card)
			} else {
				s.boardUntapped = append(s.boardUntapped, card)
			}
			copy(s.hand[i:], s.hand[i+1:])
			s.hand[len(s.hand)-1] = ""
			s.hand = s.hand[:len(s.hand)-1]
			return
		}
	}
}
