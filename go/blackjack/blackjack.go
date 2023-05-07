package blackjack

const (
	STAND = "S"
	HIT   = "H"
	SPLIT = "P"
	WIN   = "W"
)

func getCardValueMap() map[string]int {
	return map[string]int{
		"ace":   11,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"ten":   10,
		"jack":  10,
		"queen": 10,
		"king":  10,
		"other": 0,
	}
}

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	var cardValueMap map[string]int = getCardValueMap()
	return cardValueMap[card]
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	var cardValueMap map[string]int = getCardValueMap()

	cardSum := cardValueMap[card1] + cardValueMap[card2]
	switch {
	case card1 == "ace" && card2 == "ace":
		return SPLIT
	case cardSum == 21:
		if cardValueMap[dealerCard] >= 10 {
			return STAND
		} else {
			return WIN
		}
	case cardSum >= 17 && cardSum <= 20:
		return STAND
	case cardSum >= 12 && cardSum <= 16:
		if cardValueMap[dealerCard] >= 7 {
			return HIT
		} else {
			return STAND
		}
	case cardSum <= 11:
		return HIT
	}
	return ""
}
