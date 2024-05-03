package game

import (
	"reflect"
	"testing"

	deck "github.com/DarickMcBride/card-games-cli/pkg/cards"
)

func TestItIsWar(t *testing.T) {
	// Create a deck for player and bot
	playerCards := deck.NewDeck()
	botCards := deck.NewDeck()

	// Add cards to player and bot decks
	playerCards.AddToTop(deck.Card{Rank: 2, Suit: "Spades"}, deck.Card{Rank: 3, Suit: "Spades"}, deck.Card{Rank: 4, Suit: "Spades"}, deck.Card{Rank: 5, Suit: "Spades"})

	botCards.AddToTop(deck.Card{Rank: 6, Suit: "Hearts"}, deck.Card{Rank: 7, Suit: "Hearts"}, deck.Card{Rank: 8, Suit: "Hearts"}, deck.Card{Rank: 9, Suit: "Hearts"})

	// Call the function under test
	playerCards, botCards = itIsWar(playerCards, botCards)

	// Assert the expected outcome
	expectedPlayerCards := deck.NewDeck()
	expectedPlayerCards.AddToTop(deck.Card{Rank: 2, Suit: "Spades"}, deck.Card{Rank: 3, Suit: "Spades"}, deck.Card{Rank: 4, Suit: "Spades"}, deck.Card{Rank: 5, Suit: "Spades"})

	expectedBotCards := deck.NewDeck()

	if !reflect.DeepEqual(playerCards, expectedPlayerCards) {
		t.Errorf("Expected player cards: %v, but got: %v", expectedPlayerCards, playerCards)
	}

	if !reflect.DeepEqual(botCards, expectedBotCards) {
		t.Errorf("Expected bot cards: %v, but got: %v", expectedBotCards, botCards)
	}
}
