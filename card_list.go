package main

var cardlist map[string]cardInfo

func init() {
	cardlist = map[string]cardInfo{
		// basic
		"Forest": cardInfo{
			IsBasic:       true,
			ComesInTapped: false,
			TapsFor:       []string{"G"},
		},
		"Plain": cardInfo{
			IsBasic:       true,
			ComesInTapped: false,
			TapsFor:       []string{"W"},
		},
		"Swamp": cardInfo{
			IsBasic:       true,
			ComesInTapped: false,
			TapsFor:       []string{"B"},
		},
		"Island": cardInfo{
			IsBasic:       true,
			ComesInTapped: false,
			TapsFor:       []string{"U"},
		},
		"Mountain": cardInfo{
			IsBasic:       true,
			ComesInTapped: false,
			TapsFor:       []string{"R"},
		},

		// duals
		"Savannah": cardInfo{
			IsBasic:       false,
			ComesInTapped: false,
			TapsFor:       []string{"W", "G"},
		},
		"Taiga": cardInfo{
			IsBasic:       false,
			ComesInTapped: false,
			TapsFor:       []string{"R", "G"},
		},
		"Tundra": cardInfo{
			IsBasic:       false,
			ComesInTapped: false,
			TapsFor:       []string{"W", "U"},
		},
		"Underground Sea": cardInfo{
			IsBasic:       false,
			ComesInTapped: false,
			TapsFor:       []string{"U", "B"},
		},
		"Badlands": cardInfo{
			IsBasic:       false,
			ComesInTapped: false,
			TapsFor:       []string{"B", "R"},
		},
		"Scrubland": cardInfo{
			IsBasic:       false,
			ComesInTapped: false,
			TapsFor:       []string{"W", "B"},
		},
		"Volcanic Island": cardInfo{
			IsBasic:       false,
			ComesInTapped: false,
			TapsFor:       []string{"R", "U"},
		},
		"Bayou": cardInfo{
			IsBasic:       false,
			ComesInTapped: false,
			TapsFor:       []string{"G", "B"},
		},
		"Plateau": cardInfo{
			IsBasic:       false,
			ComesInTapped: false,
			TapsFor:       []string{"W", "R"},
		},
		"Tropical Island": cardInfo{
			IsBasic:       false,
			ComesInTapped: false,
			TapsFor:       []string{"G", "U"},
		},

		// slowland
		// painland
	}
}
