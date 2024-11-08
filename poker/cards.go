package poker

import (
	"fmt"
	"slices"
	"strconv"
)

type Suit struct {
	SuitId     int
	SuitEmoji  string
	SuidLetter rune
}

var (
	CLUBS = Suit{
		SuitId:     1,
		SuitEmoji:  "♣️",
		SuidLetter: 'C',
	}
	SPADES = Suit{
		SuitId:     2,
		SuitEmoji:  "♠️",
		SuidLetter: 'S',
	}
	HEARTS = Suit{
		SuitId:     3,
		SuitEmoji:  "❤️",
		SuidLetter: 'H',
	}
	DIAMONDS = Suit{
		SuitId:     4,
		SuitEmoji:  "♦️",
		SuidLetter: 'D',
	}
	SUITS = []Suit{CLUBS, SPADES, HEARTS, DIAMONDS}
)

type Card struct {
	// CardCode   int
	CardSuit   *Suit
	CardNumber int
}

type Cards struct {
	Fg      int
	Sg      int
	MaxCard int
	Pow     Power
	Dro     int
	Card1   Card
	Card2   Card
}

func InitCards(card1, card2 Card) Cards {
	res := Cards{
		Card1: card1,
		Card2: card2,
	}
	fg, sg, maxCard := GetStartetHandPower(card1, card2)
	res.Fg = fg
	res.Sg = sg
	res.MaxCard = maxCard
	return res
}

func (c *Card) String() string {
	if c == nil {
		return ""
	}
	return c.CardSuit.SuitEmoji + " " + strconv.Itoa(int(c.CardNumber))
}

func PrintCards(cards []Card) {

	for _, card := range cards {
		fmt.Print(card.String() + " ")
	}
	fmt.Println()
}

func SprintCards(cards []Card) string {
	result := ""
	for _, card := range cards {
		result += card.String() + " "
	}
	return result
}

type Power int

var (
	PowerMap = map[Power]string{
		0:  "Ничего",
		1:  "Старшая карта",
		2:  "Пара",
		3:  "Две пары",
		4:  "Сет",
		5:  "Стрит",
		6:  "Флэш",
		7:  "Фул хаус",
		8:  "Каре",
		9:  "Стрит флэш",
		10: "Роял флэш",
	}
)

const (
	Nothing Power = iota
	HighCard
	Pair
	TwoPair
	Three
	Straight
	Flush
	FullHouse
	Quads
	StraightFlush
	RoyalFlush
)

func SortOnlyNumbers(cards []Card) []int {
	result := make([]int, 0)

	for _, card := range cards {
		ok := slices.Index(result, int(card.CardNumber))
		if ok == -1 {
			result = append(result, int(card.CardNumber))
		}
	}
	return result
}

func SortByNumber(cards []Card) []Card {
	len := len(cards)
	result := make([]Card, len)
	copy(result, cards)
	for i := 0; i < len; i++ {
		for j := 0; j < len-i-1; j++ {
			if result[j].CardNumber > result[j+1].CardNumber {
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}
	return result
}

func SortBySuit(cards []Card) []Card {
	len := len(cards)
	result := make([]Card, len)
	copy(result, cards)
	for i := 0; i < len; i++ {
		for j := 0; j < len-i-1; j++ {
			if result[j].CardSuit.SuitId > result[j+1].CardSuit.SuitId {
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}
	return result
}
