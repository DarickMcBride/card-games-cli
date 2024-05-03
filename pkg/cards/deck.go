package deck

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand/v2"
	"os"
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
			cards = append(cards, NewCard(rank+" of "+suit, i+2, suit))
		}
	}

	return cards
}

// deal splits the deck into two parts, returning a hand of cards and the remaining deck.
func (d Deck) Deal(handSize int) (Deck, Deck) {
	return d[:handSize], d[handSize:]
}

// shuffle randomly reorders the cards in the Deck.
func (d Deck) Shuffle() {

	rand.Shuffle(len(d), func(i, j int) {
		(d)[i], (d)[j] = (d)[j], (d)[i]
	})

}

// Print prints the cards in the deck.
func (d Deck) Print() {
	for i, card := range d {
		fmt.Println(i, card.Name, card.Rank)
	}
}

// Draw removes and returns the top card from the deck.
func (d *Deck) Draw() (Card, error) {
	if len(*d) < 1 {
		return Card{}, errors.New("cannot draw card: deck is empty")
	}

	topCard := (*d)[0]
	if len(*d) > 0 {
		*d = (*d)[1:]
	}

	return topCard, nil
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
func (d Deck) WriteToFile(filename string) error {
	bytes, err := d.toJson()
	if err != nil {
		return err
	}

	return os.WriteFile(filename, bytes, 0666)
}


// AddToBottom adds one or more cards to the bottom of the deck.
// The cards are appended to the existing deck in the order they are provided.
// The deck is modified in-place.
func (d *Deck) AddToBottom(card ...Card) {
	*d = append(*d, card...)
}

// readDeckFromFile reads a deck of cards from a file and returns it.
// It takes a filename as input and returns a Deck and an error.
// If the file cannot be read or the data cannot be unmarshaled into a Deck,
// an error is returned.
func ReadDeckFromFile(filename string) (Deck, error) {
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
