package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten", "jack", "queen", "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	var decision string
	var dealerCardValue int = ParseCard(dealerCard)
	var myCardValue int = ParseCard(card1) + ParseCard(card2)
	switch {
	case myCardValue == 22:
		decision = "P"
	case myCardValue == 21 && !(dealerCardValue > 9):
		decision = "W"
	case myCardValue == 21 && dealerCardValue > 9:
		decision = "S"
	case (myCardValue >= 17 && myCardValue <= 20):
		decision = "S"
	case (myCardValue >= 12 && myCardValue <= 16) && dealerCardValue <= 6:
		decision = "S"
	case (myCardValue >= 12 && myCardValue <= 16) && dealerCardValue >= 7:
		decision = "H"
	case myCardValue <= 11:
		decision = "H"
	}
	return decision
}
