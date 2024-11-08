package poker

import (
	"slices"
)

func checkSraight(cards []Card) bool {
	if len(cards) != 5 {
		return false
	}
	cards = SortByNumber(cards)

	numbers := SortOnlyNumbers(cards)
	for i := 0; i < len(numbers)-4; i++ {
		lastN := numbers[i]
		needNumbers := []int{lastN, lastN + 1, lastN + 2, lastN + 3, lastN + 4}
		if slices.Equal(numbers[i:i+5], needNumbers) {
			return true
		}
	}
	return false
}
func checkFlush(cards []Card) bool {
	if len(cards) != 5 {
		return false
	}
	suit := cards[0].CardSuit
	for i := 1; i < len(cards); i++ {
		if cards[i].CardSuit != suit {
			return false
		}
	}
	return true
}
func checkQuads(cards []Card) (bool, []Card) {
	if len(cards) < 4 {
		return false, []Card{}
	}
	for _, quadsCards := range Combinations(cards, 4) {
		quadsCards = SortByNumber(quadsCards)
		if quadsCards[0].CardNumber == quadsCards[3].CardNumber {
			return true, quadsCards
		}
	}
	return false, []Card{}
}
func checkFullHouse(cards []Card) (bool, int) {
	if len(cards) != 5 {
		return false, 0
	}
	if cards[0].CardNumber == cards[1].CardNumber && cards[2].CardNumber == cards[4].CardNumber {
		return true, cards[2].CardNumber
	}
	if cards[0].CardNumber == cards[2].CardNumber && cards[3].CardNumber == cards[4].CardNumber {
		return true, cards[0].CardNumber
	}
	return false, 0
}
func checkThree(cards []Card) (bool, int) {
	if len(cards) < 3 {
		return false, 0
	}
	for _, threeCards := range Combinations(cards, 3) {
		threeCards = SortByNumber(threeCards)
		if threeCards[0].CardNumber == threeCards[2].CardNumber {
			return true, threeCards[0].CardNumber
		}
	}
	return false, 0
}
func checkPair(cards []Card) (bool, int) {
	if len(cards) < 2 {
		return false, 0
	}
	for _, pairCards := range Combinations(cards, 2) {
		if pairCards[0].CardNumber == pairCards[1].CardNumber {
			return true, pairCards[0].CardNumber
		}
	}
	return false, 0
}
func checkTwoPairs(cards []Card) (bool, int) {
	if len(cards) < 4 {
		return false, 0
	}
	allT := Combinations(cards, 2)
	pairs := Combinations(allT, 2)
	for _, p := range pairs {
		p1 := p[0]
		p2 := p[1]
		ch1, _ := checkPair(p1)
		ch2, _ := checkPair(p2)
		if ch1 && ch2 {
			return true, max(p1[0].CardNumber, p2[0].CardNumber)
		}
	}
	return false, 0
}
