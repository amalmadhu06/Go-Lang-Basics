package main

import "fmt"

func main() {
	fmt.Println(FirstTurn("ace", "ace", "jack"))
	//== "P"
	fmt.Println(FirstTurn("ace", "king", "ace"))
	//== "S"
	fmt.Println(FirstTurn("five", "queen", "ace"))
	//== "H"

}

// card		value	card	value
// ace		11		eight	8
// two		2		nine	9
// three	3		ten		10
// four		4		jack	10
// five		5		queen	10
// six		6		king	10
// seven	7		other	0

func ParseCard(card string) int {
	value := 0
	switch card {
	case "ace":
		value = 11
	case "two":
		value = 2
	case "three":
		value = 3
	case "four":
		value = 4
	case "five":
		value = 5
	case "six":
		value = 6
	case "seven":
		value = 7
	case "eight":
		value = 8
	case "nine":
		value = 9
	case "ten", "jack", "queen", "king":
		value = 10
	// case "jack":
	// 	value = 10
	// case "queen":
	// 	value = 10
	// case "king":
	// 	value = 10
	default:
		value = 0
	}

	return value
}

func FirstTurn(card1, card2, dealerCard string) string {
	card1Value := ParseCard(card1)
	card2Value := ParseCard(card2)
	dealerCardValue := ParseCard(dealerCard)
	var decision string
	switch true {
	case card1 == "ace" && card2 == "ace":
		decision = "P" //both aces --> split
	case (card1Value+card2Value == 21) && (dealerCard != "ace" && dealerCardValue != 10):
		decision = "W"
	case (card1Value+card2Value == 21) && (dealerCard == "ace" || dealerCardValue == 10):
		decision = "S"
	case (card1Value+card2Value >= 17 && card1Value+card2Value <= 20):
		decision = "S"
	case (card1Value+card2Value >= 12 && card1Value+card2Value <= 16) && (dealerCardValue < 7):
		decision = "S"
	case (card1Value+card2Value >= 12 && card1Value+card2Value <= 16) && (dealerCardValue >= 7):
		decision = "H"
	case (card1Value+card2Value <= 11):
		decision = "H"
	default:
		decision = "S"
	}

	return decision
}

// Depending on your two cards and the card of the dealer, there is a strategy for the first turn of the game, in which you have the following options:

// Stand (S)
// Hit (H)
// Split (P)
// Automatically win (W)
// Although not optimal yet, you will follow the strategy your friend Alex has been developing, which is as follows:

// If you have a pair of aces you must always split them.
// If you have a Blackjack (two cards that sum up to a value of 21), and the dealer does not have an ace, a figure or a ten then you automatically win. If the dealer does have any of those cards then you'll have to stand and wait for the reveal of the other card.
// If your cards sum up to a value within the range [17, 20] you should always stand.
// If your cards sum up to a value within the range [12, 16] you should always stand unless the dealer has a 7 or higher, in which case you should always hit.
// If your cards sum up to 11 or lower you should always hit.
