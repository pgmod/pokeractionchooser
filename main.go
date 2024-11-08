package pokeractionchooser

import (
	"github.com/pgmod/pokeractionchooser/poker"
)

type Action int

const (
	Call Action = iota
	Check
	Bet
	Raise
	Fold
)

type Stage int

const (
	PreFlop Stage = iota
	Flop
	Turn
	River
)

func ActionChooser(stage Stage, myCards poker.Cards, needCall, betCount, handChips, minBuyIn int) (Action, int) {
	// fmt.Println("Stage ", stage)
	// fmt.Println("MyCards ", myCards)
	// fmt.Println("MyCards.pow ", myCards.Pow)
	// fmt.Println("NeedCall ", needCall)
	// fmt.Println("BetCount ", betCount)
	sg := myCards.Sg
	fg := myCards.Fg

	if stage == PreFlop {

		if sg == 1 {
			return Raise, handChips // all in
		} else if sg == 2 {
			if handChips > needCall*2 {
				return Raise, needCall * 2
			} else {
				return Raise, handChips // all in
			}
		} else if sg == 3 {
			if handChips > needCall*4 {
				return Raise, needCall * 4
			} else {
				return Raise, handChips // all in
			}
		} else {
			if fg <= poker.SUITED_SEMI_CONNECTORS {
				if needCall == 0 {
					return Check, 0
				}
				if minBuyIn > needCall*7 {
					return Call, 0
				} else {
					return Fold, 0
				}
			} else if fg > poker.SUITED_SEMI_CONNECTORS {
				if betCount > 0 {
					return Fold, 0
				}
				if needCall > 0 {
					return Call, 0
				} else {
					return Check, 0
				}

			} else {
				if needCall > 0 {
					if needCall > handChips {
						return Fold, 0
					}
					return Call, needCall
				} else {
					return Check, 0
				}
			}
		}
	} else if stage >= Flop {

		mcard := myCards.MaxCard
		switch myCards.Pow {
		case poker.Nothing:
			if betCount > 0 {
				return Fold, 0
			} else {
				return Check, 0
			}
		case poker.HighCard:
			if betCount == 0 {
				return Check, 0

			} else if sg < 3 {
				return Call, 0
			} else {
				return Fold, 0
			}
		case poker.Pair:
			if betCount == 0 {
				return Check, 0

			} else if mcard > 9 {
				return Call, 0
			} else {
				return Fold, 0
			}
		case poker.TwoPair:
			if betCount == 0 {
				return Check, 0
			} else {
				if minBuyIn/4 > needCall {
					return Call, 0
				} else {
					return Fold, 0
				}
			}
		case poker.Three:
			if betCount == 0 {
				return Check, 0

			} else if mcard > 5 {
				return Call, 0
			} else {
				return Fold, 0
			}
		case poker.Straight:
			if mcard >= 10 {
				return Raise, minBuyIn * 2
			} else if betCount == 0 {
				if handChips > needCall {
					if minBuyIn >= handChips {
						return Raise, minBuyIn
					} else {
						return Raise, handChips
					}
				} else {
					return Call, 0
				}
			} else {
				if minBuyIn*3 > needCall*2 {
					return Call, 0
				} else {
					return Fold, 0
				}
			}
		case poker.Flush:
			if mcard >= 8 {
				return Raise, minBuyIn * 2
			} else if betCount == 0 {
				return Raise, minBuyIn
			} else {
				if minBuyIn/3 > needCall {
					return Call, 0
				} else {
					return Fold, 0
				}
			}
		case poker.FullHouse:
			if mcard >= 6 {
				return Raise, minBuyIn * 2
			} else if betCount == 0 {
				return Raise, minBuyIn
			} else {
				if minBuyIn/2 > needCall {
					return Call, 0
				} else {
					return Fold, 0
				}
			}
		case poker.Quads:
			return Raise, handChips
		case poker.StraightFlush:
			return Raise, handChips
		case poker.RoyalFlush:
			return Raise, handChips
		default:
			return Fold, 0
			// TODO dro Check
		}

	} else {
		return Fold, 0
	}

}
