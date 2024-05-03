package game

import (
	"fmt"
	"os"

	deck "github.com/DarickMcBride/card-games-cli/pkg/cards"
)

func War() {
	// initialize a new deck of cards and shuffle it
	deck := deck.NewDeck()
	deck.Shuffle()

	// deal the cards to the players splitting the deck in half
	playerCards, botCards := deck.Deal(26)

	playerScore := 0
	botScore := 0

	// play the game till one of the players runs out of cards
	for len(playerCards) > 0 && len(botCards) > 0 {
		printScores(playerScore, botScore)
		// draw a card from each player's deck
		playerCard, err := playerCards.Draw()
		if err != nil {
			fmt.Println(err)
		}

		botCard, err := botCards.Draw()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Player drew", playerCard)
		fmt.Println("Bot drew", botCard)
		fmt.Println("Player has", len(playerCards), "cards left")
		fmt.Println("Bot has", len(botCards), "cards left")

		// compare the ranks of the cards
		if playerCard.Rank > botCard.Rank {
			fmt.Println("Player takes card!")
			playerScore++
			playerCards.AddToBottom(playerCard)
			playerCards.AddToBottom(botCard)
		} else if botCard.Rank > playerCard.Rank {
			fmt.Println("Bot takes card!")
			botScore++
			botCards.AddToBottom(botCard)
			botCards.AddToBottom(playerCard)
		} else {
			playerCards, botCards = itIsWar(playerCards, botCards)
		}

		// if one of the players runs out of cards, print the winner
		if len(playerCards) == 0 {
			fmt.Println("Bot wins!")

		} else if len(botCards) == 0 {
			fmt.Println("Player wins!")

		}
	}
	fmt.Println("Game Over")
	os.Exit(0)

}

//draw 3

// war time
// itIsWar simulates a war between the player and the bot in a card game.
// It takes two decks of cards, playerCards and botCards, as input and returns the updated decks after the war.
// If either player does not have enough cards to play a war (less than 4 cards), the original decks are returned.
// During the war, each player draws 4 cards from their deck.
// The last card of each player's war cards is compared, and the player with the higher rank wins the war.
// If there is a tie, another war is initiated recursively until a winner is determined.
// The function prints the results of each war to the console.
// The updated playerCards and botCards are returned after all wars are resolved.
func itIsWar(playerCards, botCards deck.Deck) (deck.Deck, deck.Deck) {
	// check if the players have enough cards to play a war
	if len(playerCards) < 4 {
		return playerCards, botCards
	} else if len(botCards) < 4 {
		return playerCards, botCards
	}

	// draw 4 cards from each player's deck
	playerWarCards, playerRemaining := playerCards.Deal(4)
	botWarCards, botRemaining := botCards.Deal(4)

	// compare the last card of the war
	playerLastCard := playerWarCards[len(playerWarCards)-1]
	botLastCard := botWarCards[len(botWarCards)-1]

	fmt.Println("Player drew", playerLastCard)
	fmt.Println("Bot drew", botLastCard)

	if playerLastCard.Rank > botLastCard.Rank {
		fmt.Println("Player wins the war!")
		playerCards.AddToBottom(playerWarCards...)
		playerCards.AddToBottom(botWarCards...)
	} else if botLastCard.Rank > playerLastCard.Rank {
		fmt.Println("Bot wins the war!")
		botCards.AddToBottom(botWarCards...)
		botCards.AddToBottom(playerWarCards...)
	} else {
		// if there is another war
		playerCards, botCards = itIsWar(playerRemaining, botRemaining)
	}

	return playerCards, botCards

}

func printScores(playerScore, botScore int) {
	fmt.Println("Player score:", playerScore)
	fmt.Println("Bot score:", botScore)
}
