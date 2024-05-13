package game

import (
	"fmt"
	"os"

	cards "github.com/DarickMcBride/card-games-cli/pkg/cards"
	deck "github.com/DarickMcBride/card-games-cli/pkg/cards"
)

func War() {
	fmt.Println("Welcome to War!")
	fmt.Println("Press 'Enter' to begin the game")
	fmt.Scanln()
	// initialize a new deck of cards and shuffle it
	deck := deck.NewDeck()
	deck.Shuffle()

	// deal the cards to the players splitting the deck in half
	playerCards, botCards := deck.Deal(26)

	// play the game till one of the players runs out of cards
	for len(playerCards) > 0 && len(botCards) > 0 {
		fmt.Println("Player has", len(playerCards), "cards left")
		fmt.Println("Bot has", len(botCards), "cards left")

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

		wonCards := []cards.Card{playerCard, botCard}

		isGameOver := false

		playerWins := false

		// compare the ranks of the cards
		if playerCard.Rank > botCard.Rank {
			fmt.Println("Player takes card!")
			playerWins = true

		} else if botCard.Rank > playerCard.Rank {
			fmt.Println("Bot takes card!")
			playerWins = false

		} else {

			playerCards, botCards, wonCards, playerWins, isGameOver = itIsWar(playerCards, botCards, wonCards)

		}

		if playerWins {
			playerCards = append(playerCards, wonCards...)
			playerCards.Shuffle()
		} else {
			botCards = append(botCards, wonCards...)
			botCards.Shuffle()
		}

		//add the won cards to the winner's deck

		fmt.Println("Press 'Enter' to continue")
		fmt.Scanln()

		// if one of the players runs out of cards, print the winner
		if len(playerCards) == 0 || isGameOver {
			fmt.Println("Bot wins!")
			gameOver()

		} else if len(botCards) == 0 || isGameOver {
			fmt.Println("Player wins!")
			gameOver()

		}

	}

}

func itIsWar(playerCards deck.Deck, botCards deck.Deck, wonCards []cards.Card) (deck.Deck, deck.Deck, []cards.Card, bool, bool) {
	isGameOver := false
	playerWins := false

	for {
		fmt.Println("It is War!")
		fmt.Println("Press 'Enter' to continue")
		fmt.Scanln()
		// check if the players have enough cards to play a war
		if len(playerCards) < 4 {
			fmt.Println("Bot wins!")
			isGameOver = true
			break
		} else if len(botCards) < 4 {
			fmt.Println("Player wins!")
			isGameOver = true
			break
		}

		// draw 4 cards from each player's deck
		playerWarCards, pc := playerCards.Deal(4)
		botWarCards, bc := botCards.Deal(4)

		// compare the last card of the war
		playerLastCard := playerWarCards[len(playerWarCards)-1]
		botLastCard := botWarCards[len(botWarCards)-1]

		fmt.Println("Player drew", playerLastCard)
		fmt.Println("Bot drew", botLastCard)

		playerCards = pc
		botCards = bc

		wonCards = append(wonCards, playerWarCards...)

		wonCards = append(wonCards, botWarCards...)

		if playerLastCard.Rank > botLastCard.Rank {
			fmt.Println("Player wins the war!")
			playerWins = true
			break

		} else if botLastCard.Rank > playerLastCard.Rank {
			fmt.Println("Bot wins the war!")
			playerWins = false
			break

		} else {
			fmt.Println("Another war!")
			fmt.Println("Press 'Enter' to continue")
			fmt.Scanln()
			// if there is another w
		}

	}

	return playerCards, botCards, wonCards, playerWins, isGameOver

}

func gameOver() {
	fmt.Println("Game Over")
	fmt.Println("Play again? (y/n)")
	var input string
	fmt.Scanln(&input)
	if input == "y" {
		War()
	}
	os.Exit(0)

}

func printScores(playerScore, botScore int) {
	fmt.Println("Player score:", playerScore)
	fmt.Println("Bot score:", botScore)
}
