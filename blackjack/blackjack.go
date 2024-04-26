// Package blackjack provides functions related to playing blackjack.
package blackjack

var cardsValues = map[string]int{
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
	"ace":   11,
}

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	return cardsValues[card]
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	dealerCardValue := ParseCard(dealerCard)
	cardsValuesSum := ParseCard(card1) + ParseCard(card2)
	var strategy string
	switch {
	case cardsValuesSum == 22:
		strategy = "P"
	case cardsValuesSum == 21:
		strategy = "W"
		if dealerCardValue >= 10 && dealerCardValue <= 11 {
			strategy = "S"
		}
	case cardsValuesSum <= 20 && cardsValuesSum >= 17:
		strategy = "S"
	case cardsValuesSum <= 16 && cardsValuesSum >= 12:
		strategy = "S"
		if dealerCardValue >= 7 {
			strategy = "H"
		}
	case cardsValuesSum <= 11:
		strategy = "H"
	}
	return strategy
}
