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

// Stand (S)
// Hit (H)
// Split (P)
// Automatically win (W)

func FirstTurn(card1, card2, dealerCard string) string {
	card1Value := ParseCard(card1)
	card2Value := ParseCard(card2)
	dealerCardValue := ParseCard(dealerCard)
	var decision string
	switch true {
	case card1 == "ace" && card2 == "ace":
		decision = "P"
	case (card1Value+card2Value == 21) && (dealerCard != "ace" || dealerCardValue != 10):
		decision = "W"
	case (card1Value+card2Value == 21) && (dealerCard == "ace" || dealerCardValue == 10):
		decision = "S"
	case (card1Value+card2Value >= 17 && card1Value+card2Value <= 20):
		decision = "S"
	case (card1Value+card2Value >= 12 && card1Value+card2Value <= 16):
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
