package deck

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"os"
	"github.com/DarickMcBride/card-games-cli/pkg/card"
)

// Create a new type of 'Deck'
// slice of strings
type Deck []Card

// NewDeck creates and returns a new Deck of cards.
func NewDeck() Deck {
	cards := Deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	ranks := []string{"Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}

	// create a Deck of cards
	for _, suit := range cardSuits {
		for i, rank := range ranks {
			cards = append(cards, NewCard(rank+" of "+suit, i+2))
		}
	}

	return cards
}

// deal splits the deck into two parts, returning a hand of cards and the remaining deck.
func (d Deck) deal(handSize int) (Deck, Deck) {
	return d[:handSize], d[handSize:]
}

// shuffle randomly reorders the cards in the Deck.
func (d Deck) shuffle() {

	rand.Shuffle(len(d), func(i, j int) {
		(d)[i], (d)[j] = (d)[j], (d)[i]
	})

}

func (d Deck) print() {
	for i, card := range d {
		fmt.Println(i, card.Name, card.Rank)
	}
}

// toJson converts the deck to JSON format.
// It returns a byte slice containing the JSON representation of the deck and an error, if any.
func (d Deck) toJson() ([]byte, error) {
	bytes, err := json.Marshal(d)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// writeToFile writes the deck to a file with the specified filename.
// It converts the deck to JSON format and writes it to the file.
// Returns an error if there was a problem writing the file.
func (d Deck) writeToFile(filename string) error {
	bytes, err := d.toJson()
	if err != nil {
		return err
	}

	return os.WriteFile(filename, bytes, 0666)
}

// readDeckFromFile reads a deck of cards from a file and returns it.
// It takes a filename as input and returns a Deck and an error.
// If the file cannot be read or the data cannot be unmarshaled into a Deck,
// an error is returned.
func readDeckFromFile(filename string) (Deck, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var deck Deck
	err = json.Unmarshal(data, &deck)
	if err != nil {
		return nil, err
	}

	return deck, nil
}
