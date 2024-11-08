package poker

var (
	// пара
	POCKET_PAIRS = 1

	// одномастные коннекторы
	SUITED_CONNECTORS = 2

	// разномастные коннекторы
	NOTSUITED_CONNECTORS = 3

	// одномастные полуконнекторы
	SUITED_SEMI_CONNECTORS = 4

	// разномастные полуконнекторы
	NOTSUITED_SEMI_CONNECTORS = 5

	// несвязанные одномастные
	UNRELATED_SUITED = 6

	// несвязанные разномастные
	UNRELATED_OFFSUIT = 7
	NAME_MAP          = map[int]string{
		1: "пара",
		2: "одномастные коннекторы",
		3: "разномастные коннекторы",
		4: "одномастные полуконнекторы",
		5: "разномастные полуконнекторы",
		6: "несвязанные одномастные",
		7: "несвязанные разномастные",
	}
	Deck = []Card{}
)

func init() {
	// init deck
	Deck = make([]Card, 52)
	for i := 0; i < 52; i++ {
		Deck[i] = Card{
			CardNumber: int(i%13 + 2),
			CardSuit:   &SUITS[i/13],
		}
	}
}
func GetCardsPower(cards []Card) (Power, int, []Card) {
	if len(cards) < 2 {
		return 0, 0, []Card{}
	}
	cards = SortByNumber(cards)
	suts := make(map[int][]Card)
	for _, card := range cards {
		suts[int(card.CardSuit.SuitId)] = append(suts[int(card.CardSuit.SuitId)], card)
	}

	isSraight := checkSraight(cards)
	isFlush := checkFlush(cards)
	isStraightFlush := isSraight && isFlush

	if isStraightFlush {
		if cards[4].CardNumber == 14 {
			return 10, 14, cards
		}
		return 9, cards[4].CardNumber, cards
	}

	isQuads, QuadCards := checkQuads(cards)
	if isQuads {
		return 8, QuadCards[0].CardNumber, QuadCards
	}

	isFullHouse, hight := checkFullHouse(cards)
	if isFullHouse {
		return 7, hight, cards
	}
	if isFlush {
		return 6, cards[4].CardNumber, cards
	}
	if isSraight {
		return 5, cards[4].CardNumber, cards
	}
	isThree, threeCard := checkThree(cards)
	if isThree {
		return 4, threeCard, cards
	}

	isTwoPairs, hight := checkTwoPairs(cards)
	if isTwoPairs {
		return 3, hight, cards
	}

	isPair, pairCard := checkPair(cards)
	if isPair {
		return 2, pairCard, cards
	} else {
		return 1, cards[len(cards)-1].CardNumber, cards
	}
}
func GetStartetHandPower(c1, c2 Card) (int, int, int) {
	var first_group_id int
	var second_group_id int
	var max_card_number int
	if c1.CardNumber == c2.CardNumber {
		// пара
		first_group_id = POCKET_PAIRS
		max_card_number = max(c1.CardNumber, c2.CardNumber)
	} else if (c1.CardNumber+1 == c2.CardNumber ||
		c1.CardNumber-1 == c2.CardNumber) &&
		c1.CardSuit.SuitId == c2.CardSuit.SuitId {
		// одномастные коннекторы
		first_group_id = SUITED_CONNECTORS
		max_card_number = max(c1.CardNumber, c2.CardNumber)
	} else if c1.CardNumber+1 == c2.CardNumber || c1.CardNumber-1 == c2.CardNumber {
		// разномастные коннекторы
		first_group_id = NOTSUITED_CONNECTORS
		max_card_number = max(c1.CardNumber, c2.CardNumber)
	} else if (c1.CardNumber+2 == c2.CardNumber ||
		c1.CardNumber-2 == c2.CardNumber) &&
		c1.CardSuit.SuitId == c2.CardSuit.SuitId {
		// одномастные полуконнекторы
		first_group_id = SUITED_SEMI_CONNECTORS
		max_card_number = max(c1.CardNumber, c2.CardNumber)
	} else if c1.CardNumber+2 == c2.CardNumber || c1.CardNumber-2 == c2.CardNumber {
		// разномастные полуконнекторы
		first_group_id = NOTSUITED_SEMI_CONNECTORS
		max_card_number = max(c1.CardNumber, c2.CardNumber)
	} else if c1.CardSuit.SuitId == c2.CardSuit.SuitId {
		// несвязанные одномастные
		first_group_id = UNRELATED_SUITED
		max_card_number = max(c1.CardNumber, c2.CardNumber)
	} else {
		// несвязанные разномастные
		first_group_id = UNRELATED_OFFSUIT
		max_card_number = max(c1.CardNumber, c2.CardNumber)
	}

	// first group
	if first_group_id == POCKET_PAIRS && c1.CardNumber >= 13 { // король и более
		second_group_id = 1
		// second group
	} else if (first_group_id == POCKET_PAIRS && c1.CardNumber >= 11) ||
		(first_group_id == SUITED_CONNECTORS && c1.CardNumber+c2.CardNumber == 27) { // валет и более или король с тузом(27 в сумме)
		second_group_id = 2

	} else if (first_group_id == POCKET_PAIRS && c1.CardNumber == 10) ||
		(first_group_id == NOTSUITED_CONNECTORS && c1.CardNumber+c2.CardNumber == 27) ||
		(first_group_id == SUITED_CONNECTORS && c1.CardNumber+c2.CardNumber == 25) ||
		(first_group_id == SUITED_SEMI_CONNECTORS && c1.CardNumber+c2.CardNumber == 27) {
		second_group_id = 3
	} else {
		second_group_id = 4
	}

	return first_group_id, second_group_id, max_card_number
}
func GetAllCardsPower(allCards []Card) (Power, int, []Card) {

	maxPower := Power(0)
	maxCard := 0
	maxCards := []Card{}
	maxPower2 := Power(0)
	maxCard2 := 0
	maxCards2 := []Card{}

	for _, cards := range Combinations(allCards, 5) {
		power, maxcard, cards := GetCardsPower(cards)
		if power > maxPower {
			maxPower = power
			maxCard = int(maxcard)
			maxCards = cards
		}
	}
	for _, cards := range Combinations(allCards, 5) {
		power, maxcard, cards := GetCardsPower(cards)
		if power > maxPower {
			maxPower = power
			maxCard = int(maxcard)
			maxCards = cards
		}
	}
	for i, card := range allCards {
		if card.CardNumber == 14 {
			allCards[i].CardNumber = 1
		}
	}
	for _, cards := range Combinations(allCards, 5) {
		power, maxcard, cards := GetCardsPower(cards)
		if power > maxPower2 {
			maxPower2 = power
			maxCard2 = int(maxcard)
			maxCards2 = cards
		}
	}
	if maxPower < maxPower2 {
		maxPower = maxPower2
		maxCard = maxCard2
		maxCards = maxCards2
	}
	// fmt.Println("maxPower", combinations[maxPower], "maxCard", maxCard, "Cards", SprintCards(maxCards))
	return maxPower, maxCard, maxCards
}

func CountFlashDro(cards []Card) int { //подсчет одной недостабющей карты
	closedCount := 52 - len(cards)
	// NLH
	// 2 mine and 3-5 on board
	// flop - 5
	suts := make(map[int][]Card)
	m_count := 0
	for _, card := range cards {
		suts[int(card.CardSuit.SuitId)] = append(suts[int(card.CardSuit.SuitId)], card)
		if len(suts[int(card.CardSuit.SuitId)]) > m_count {
			m_count = len(suts[int(card.CardSuit.SuitId)])
		}
	}

	if len(cards) == 5 {
		if m_count == 4 {
			return 9/closedCount*100 + 9/(closedCount-1)*100
		}
	} else if len(cards) == 6 {
		if m_count == 4 {
			return 9 / closedCount * 100
		}
	}
	return 0
}
func CountStraightDro(cards []Card) int { //подсчет одной недостабющей карты

	closedCount := float32(52 - len(cards))
	cards = SortByNumber(cards)
	numbers := SortOnlyNumbers(cards)
	auts := float32(0)
	for _, card := range Deck {
		if ContainsInArray(int(card.CardNumber), numbers) {
			continue
		} else {
			for _, c := range Combinations(append(cards, card), 5) {
				// fmt.Println("c", SprintCards(c))
				r := checkSraight(c)
				if r {
					auts += 1
				}
			}
		}
	}
	var res float32
	if len(cards) == 6 {
		res = auts / closedCount
	} else if len(cards) == 5 {
		res = auts/closedCount + auts/(closedCount-1)

	} else {

		res = auts / closedCount
	}

	return int(res * 100)
}

// turn - 6
// river - 7

// TODO get cards power dro chance
// TODO алгоритм подсчета аутов
