package blackjack

// ruleset maps a card to its value.
var ruleset = map[string]int{
	"ace": 11, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7,
	"eight": 8, "nine": 9, "ten": 10, "jack": 10, "queen": 10, "king": 10,
}

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	value, ok := ruleset[card]

	if ok {
		return value
	}

	return 0
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	sum := ParseCard(card1) + ParseCard(card2)
	var strategy string

	switch {
	case sum == 22:
		strategy = "P"
	case sum == 21:
		if ParseCard(dealerCard) < 10 {
			strategy = "W"
		} else {
			strategy = "S"
		}
	case sum >= 17 && sum <= 20:
		strategy = "S"
	case sum >= 12 && sum <= 16:
		if ParseCard(dealerCard) >= 7 {
			strategy = "H"
		} else {
			strategy = "S"
		}
	case sum <= 11:
		strategy = "H"
	}

	return strategy
}
