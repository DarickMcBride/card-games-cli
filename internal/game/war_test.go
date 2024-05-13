package game

import (
	"reflect"
	"testing"

	deck "github.com/DarickMcBride/card-games-cli/pkg/cards"
)

func TestItIsWarOutOfCards(t *testing.T) {
	// Test when one player is out of cards
	playerCards := deck.Deck{deck.Card{Name: "Jack of Spades", Rank: 11, Suit: "Spades"}}

	botCards := deck.Deck{deck.Card{Name: "10 of Spades", Rank: 10, Suit: "Spades"}}

	_, _, _, _, isGameOver := itIsWar(playerCards, botCards, []deck.Card{})

	if isGameOver == false {
		t.Errorf("Expected isGameOver to be true, but got false")
	}

}

func TestItIsWar(t *testing.T) {
	var err error

	// Test a single round of war
	playerCards := deck.Deck{deck.Card{Name: "10 of Spades", Rank: 10, Suit: "Spades"}, deck.Card{Name: "Jack of Spades", Rank: 11, Suit: "Spades"}, deck.Card{Name: "Queen of Spades", Rank: 12, Suit: "Spades"}, deck.Card{Name: "King of Spades", Rank: 13, Suit: "Spades"}, deck.Card{Name: "Ace of Spades", Rank: 14, Suit: "Spades"}}

	botCards := deck.Deck{deck.Card{Name: "2 of Hearts", Rank: 2, Suit: "Hearts"}, deck.Card{Name: "3 of Hearts", Rank: 3, Suit: "Hearts"}, deck.Card{Name: "4 of Hearts", Rank: 4, Suit: "Hearts"}, deck.Card{Name: "5 of Hearts", Rank: 5, Suit: "Hearts"}, deck.Card{Name: "6 of Hearts", Rank: 6, Suit: "Hearts"}}

	playerCards, botCards, wonCards, playerWins, isGameOver := itIsWar(playerCards, botCards, []deck.Card{{Name: "9 of Spades", Rank: 9, Suit: "Spades"}, {Name: "9 of Hearts", Rank: 9, Suit: "Hearts"}})

	expectedPlayerCards := deck.Deck{deck.Card{Name: "Ace of Spades", Rank: 14, Suit: "Spades"}}

	expectedBotCards := deck.Deck{deck.Card{Name: "6 of Hearts", Rank: 6, Suit: "Hearts"}}

	expectedWonCards := []deck.Card{{Name: "9 of Spades", Rank: 9, Suit: "Spades"}, {Name: "9 of Hearts", Rank: 9, Suit: "Hearts"}, {Name: "10 of Spades", Rank: 10, Suit: "Spades"}, {Name: "Jack of Spades", Rank: 11, Suit: "Spades"}, {Name: "Queen of Spades", Rank: 12, Suit: "Spades"}, {Name: "King of Spades", Rank: 13, Suit: "Spades"}, {Name: "2 of Hearts", Rank: 2, Suit: "Hearts"}, {Name: "3 of Hearts", Rank: 3, Suit: "Hearts"}, {Name: "4 of Hearts", Rank: 4, Suit: "Hearts"}, {Name: "5 of Hearts", Rank: 5, Suit: "Hearts"}}

	expectedPlayerWins := true

	// Check for errors
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if isGameOver == true {
		t.Errorf("Expected isGameOver to be false, but got true")
	}

	if !reflect.DeepEqual(playerCards, expectedPlayerCards) {
		t.Errorf("Player cards were %v, expected %v", playerCards, expectedPlayerCards)
	}
	if !reflect.DeepEqual(botCards, expectedBotCards) {
		t.Errorf("Bot cards were %v, expected %v", botCards, expectedBotCards)
	}
	if !reflect.DeepEqual(wonCards, expectedWonCards) {
		t.Errorf("Won cards were %v, expected %v", wonCards, expectedWonCards)
	}
	if playerWins != expectedPlayerWins {
		t.Errorf("Expected playerWins to be %v, but got %v", expectedPlayerWins, playerWins)
	}

}
