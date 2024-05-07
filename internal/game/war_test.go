package game

import (
	"reflect"
	"testing"

	deck "github.com/DarickMcBride/card-games-cli/pkg/cards"
)

func TestItIsWar(t *testing.T) {
	// Test when one player is out of cards
	notEnoughCards := deck.Deck{deck.Card{Name: "Jack of Spades", Rank: 11, Suit: "Spades"}}

	itIsWar(notEnoughCards, notEnoughCards)

	// test a single round of war
	playerCards := deck.Deck{deck.Card{Name: "Jack of Spades", Rank: 11, Suit: "Spades"}, deck.Card{Name: "Queen of Spades", Rank: 12, Suit: "Spades"}, deck.Card{Name: "King of Spades", Rank: 13, Suit: "Spades"}, deck.Card{Name: "Ace of Spades", Rank: 14, Suit: "Spades"}}
	botCards := deck.Deck{deck.Card{Name: "2 of Hearts", Rank: 2, Suit: "Hearts"}, deck.Card{Name: "3 of Hearts", Rank: 3, Suit: "Hearts"}, deck.Card{Name: "4 of Hearts", Rank: 4, Suit: "Hearts"}, deck.Card{Name: "5 of Hearts", Rank: 5, Suit: "Hearts"}}

	actualPlayerCards, actualBotCards := itIsWar(playerCards, botCards)

	expectedPlayerCards := deck.Deck{deck.Card{Name: "Jack of Spades", Rank: 11, Suit: "Spades"}, deck.Card{Name: "Queen of Spades", Rank: 12, Suit: "Spades"}, deck.Card{Name: "King of Spades", Rank: 13, Suit: "Spades"}, deck.Card{Name: "Ace of Spades", Rank: 14, Suit: "Spades"}, deck.Card{Name: "2 of Hearts", Rank: 2, Suit: "Hearts"}, deck.Card{Name: "3 of Hearts", Rank: 3, Suit: "Hearts"}, deck.Card{Name: "4 of Hearts", Rank: 4, Suit: "Hearts"}, deck.Card{Name: "5 of Hearts", Rank: 5, Suit: "Hearts"}}
	expectedBotCards := deck.Deck{}

	// Check if the player and bot cards match the expected cards
	if !reflect.DeepEqual(actualPlayerCards, expectedPlayerCards) {
		t.Errorf("Player cards were %v, expected %v", actualPlayerCards, expectedPlayerCards)
	}

	if !reflect.DeepEqual(actualBotCards, expectedBotCards) {
		t.Errorf("Bot cards were %v, expected %v", actualBotCards, expectedBotCards)
	}

	// test multiple rounds of war

	playerCards = deck.Deck{deck.Card{Name: "Jack of Spades", Rank: 11, Suit: "Spades"}, deck.Card{Name: "Queen of Spades", Rank: 12, Suit: "Spades"}, deck.Card{Name: "King of Spades", Rank: 13, Suit: "Spades"}, deck.Card{Name: "Ace of Spades", Rank: 14, Suit: "Spades"}}
	botCards = deck.Deck{deck.Card{Name: "2 of Hearts", Rank: 2, Suit: "Hearts"}, deck.Card{Name: "3 of Hearts", Rank: 3, Suit: "Hearts"}, deck.Card{Name: "4 of Hearts", Rank: 4, Suit: "Hearts"}, deck.Card{Name: "5 of Hearts", Rank: 14, Suit: "Hearts"}}

	actualPlayerCards, actualBotCards = itIsWar(playerCards, botCards)

	expectedPlayerCards = deck.Deck{deck.Card{Name: "Jack of Spades", Rank: 11, Suit: "Spades"}, deck.Card{Name: "Queen of Spades", Rank: 12, Suit: "Spades"}, deck.Card{Name: "King of Spades", Rank: 13, Suit: "Spades"}, deck.Card{Name: "Ace of Spades", Rank: 14, Suit: "Spades"}, deck.Card{Name: "2 of Hearts", Rank: 2, Suit: "Hearts"}, deck.Card{Name: "3 of Hearts", Rank: 3, Suit: "Hearts"}, deck.Card{Name: "4 of Hearts", Rank: 4, Suit: "Hearts"}, deck.Card{Name: "5 of Hearts", Rank: 14, Suit: "Hearts"}}

	expectedBotCards = deck.Deck{}

	// Check if the player and bot cards match the expected cards
	if !reflect.DeepEqual(actualPlayerCards, expectedPlayerCards) {
		t.Errorf("Player cards were %v, expected %v", actualPlayerCards, expectedPlayerCards)
	}

	if !reflect.DeepEqual(actualBotCards, expectedBotCards) {
		t.Errorf("Bot cards were %v, expected %v", actualBotCards, expectedBotCards)
	}

	

}
