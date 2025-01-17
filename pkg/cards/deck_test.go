package deck

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
	"testing"
)

func TestDeal(t *testing.T) {
	d := NewDeck()
	handSize := 5

	hand, remaining := d.Deal(handSize)

	if len(hand) != handSize {
		t.Errorf("Expected hand size of %d, but got %d", handSize, len(hand))
	}

	if len(remaining) != len(d)-handSize {
		t.Errorf("Expected remaining deck size of %d, but got %d", len(d)-handSize, len(remaining))
	}
}

func TestNewDeck(t *testing.T) {
	d := NewDeck()

	expectedSize := 52
	if len(d) != expectedSize {
		t.Errorf("Expected deck size of %d, but got %d", expectedSize, len(d))
	}

	expectedFirstCard := "Two of Spades"
	if d[0].Name != expectedFirstCard {
		t.Errorf("Expected first card to be %s, but got %s", expectedFirstCard, d[0].Name)
	}

	expectedLastCard := "Ace of Clubs"
	if d[len(d)-1].Name != expectedLastCard {
		t.Errorf("Expected last card to be %s, but got %s", expectedLastCard, d[len(d)-1].Name)
	}
}

func TestShuffle(t *testing.T) {
	d := NewDeck()
	originalDeck := make(Deck, len(d))
	copy(originalDeck, d)

	d.Shuffle()

	if len(d) != len(originalDeck) {
		t.Errorf("Expected deck size to remain the same after shuffling, but got %d", len(d))
	}

	if reflect.DeepEqual(d, originalDeck) {
		t.Errorf("Expected deck to be shuffled, but it remained the same")
	}
}

func TestPrint(t *testing.T) {
	d := NewDeck()

	// Redirect stdout to capture the output
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	d.Print()

	// Reset stdout
	w.Close()
	os.Stdout = old

	expectedOutput := "0 Two of Spades 2\n1 Three of Spades 3\n2 Four of Spades 4\n3 Five of Spades 5\n4 Six of Spades 6\n5 Seven of Spades 7\n6 Eight of Spades 8\n7 Nine of Spades 9\n8 Ten of Spades 10\n9 Jack of Spades 11\n10 Queen of Spades 12\n11 King of Spades 13\n12 Ace of Spades 14\n13 Two of Hearts 2\n14 Three of Hearts 3\n15 Four of Hearts 4\n16 Five of Hearts 5\n17 Six of Hearts 6\n18 Seven of Hearts 7\n19 Eight of Hearts 8\n20 Nine of Hearts 9\n21 Ten of Hearts 10\n22 Jack of Hearts 11\n23 Queen of Hearts 12\n24 King of Hearts 13\n25 Ace of Hearts 14\n26 Two of Diamonds 2\n27 Three of Diamonds 3\n28 Four of Diamonds 4\n29 Five of Diamonds 5\n30 Six of Diamonds 6\n31 Seven of Diamonds 7\n32 Eight of Diamonds 8\n33 Nine of Diamonds 9\n34 Ten of Diamonds 10\n35 Jack of Diamonds 11\n36 Queen of Diamonds 12\n37 King of Diamonds 13\n38 Ace of Diamonds 14\n39 Two of Clubs 2\n40 Three of Clubs 3\n41 Four of Clubs 4\n42 Five of Clubs 5\n43 Six of Clubs 6\n44 Seven of Clubs 7\n45 Eight of Clubs 8\n46 Nine of Clubs 9\n47 Ten of Clubs 10\n48 Jack of Clubs 11\n49 Queen of Clubs 12\n50 King of Clubs 13\n51 Ace of Clubs 14\n"

	out, _ := io.ReadAll(r)
	if string(out) != expectedOutput {
		t.Errorf("Expected output:\n%s\n\nBut got:\n%s", expectedOutput, string(out))
	}
}

func TestWriteToFile(t *testing.T) {
	d := Deck{
		Card{Name: "Two of Spades", Rank: 2},
		Card{Name: "Three of Spades", Rank: 3},
		Card{Name: "Four of Spades", Rank: 4},
		// Add more cards here if needed
	}

	filename := "test_deck.json"

	err := d.WriteToFile(filename)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Read the file contents
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedBytes, err := d.toJson()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(fileBytes, expectedBytes) {
		t.Errorf("File contents do not match expected bytes")
	}

	// Clean up the test file
	err = os.Remove(filename)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestReadDeckFromFile(t *testing.T) {
	filename := "test_deck.json"

	// Create a test deck
	testDeck := Deck{
		Card{Name: "Two of Spades", Rank: 2},
		Card{Name: "Three of Spades", Rank: 3},
		Card{Name: "Four of Spades", Rank: 4},
		// Add more cards here if needed
	}

	// Convert the test deck to JSON
	testDeckJSON, err := json.Marshal(testDeck)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Write the test deck JSON to a file
	err = os.WriteFile(filename, testDeckJSON, 0644)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Call the function being tested
	deck, err := ReadDeckFromFile(filename)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Compare the returned deck with the test deck
	if !reflect.DeepEqual(deck, testDeck) {
		t.Errorf("Returned deck does not match the test deck")
	}

	// Clean up the test file
	err = os.Remove(filename)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestToJson(t *testing.T) {
	d := Deck{
		Card{Name: "Two of Spades", Rank: 2},
		Card{Name: "Three of Spades", Rank: 3},
		Card{Name: "Four of Spades", Rank: 4},
		// Add more cards here if needed
	}

	expectedBytes, err := json.Marshal(d)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	bytes, err := d.toJson()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(bytes, expectedBytes) {
		t.Errorf("Returned bytes do not match the expected bytes")
	}
}

func TestDraw(t *testing.T) {
	d := Deck{
		Card{Name: "Two of Spades", Rank: 2, Suit: "Spades"},
		Card{Name: "Three of Spades", Rank: 3, Suit: "Spades"},
		Card{Name: "Four of Spades", Rank: 4, Suit: "Spades"},
	}
	expectedDeckSize := 2

	expectedCard := Card{Name: "Two of Spades", Rank: 2, Suit: "Spades"}
	actualCard, err := d.Draw()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if actualCard != expectedCard {
		t.Errorf("Expected card %v, but got %v", expectedCard, actualCard)
	}


	if len(d) != expectedDeckSize {
		t.Errorf("Expected deck size of %d, but got %d", expectedDeckSize, len(d))
	}

	d.Draw()

	// Draw last card
	expectedCard = Card{Name: "Four of Spades", Rank: 4, Suit: "Spades"}
	actualCard, err = d.Draw()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if actualCard != expectedCard {
		t.Errorf("Expected card %v, but got %v", expectedCard, actualCard)
	}
	fmt.Println(d)

	expectedDeckSize = len(d)
	if len(d) != expectedDeckSize {
		t.Errorf("Expected deck size of %d, but got %d", expectedDeckSize, len(d))
	}

	// Draw on empty deck
	_, err = d.Draw()
	if err == nil {
		t.Errorf("Expected error when drawing from an empty deck, but got nil")
	}
}

func TestAddToBottom(t *testing.T) {
	d := Deck{
		Card{Name: "Two of Spades", Rank: 2},
		Card{Name: "Three of Spades", Rank: 3},
		Card{Name: "Four of Spades", Rank: 4},
	}

	card := Card{Name: "Five of Spades", Rank: 5}
	expectedDeckSize := len(d) + 1

	d.AddToBottom(card)

	if len(d) != expectedDeckSize {
		t.Errorf("Expected deck size of %d, but got %d", expectedDeckSize, len(d))
	}

	if d[len(d)-1] != card {
		t.Errorf("Expected last card to be %v, but got %v", card, d[len(d)-1])
	}
}


func TestAddToBottomMultipleCards(t *testing.T) {
	d := Deck{
		Card{Name: "Two of Spades", Rank: 2},
		Card{Name: "Three of Spades", Rank: 3},
		Card{Name: "Four of Spades", Rank: 4},
	}

	cards := []Card{
		{Name: "Five of Spades", Rank: 5},
		{Name: "Six of Spades", Rank: 6},
		{Name: "Seven of Spades", Rank: 7},
	}
	expectedDeckSize := len(d) + len(cards)

	d.AddToBottom(cards...)

	if len(d) != expectedDeckSize {
		t.Errorf("Expected deck size of %d, but got %d", expectedDeckSize, len(d))
	}

	for i, card := range cards {
		if d[len(d)-len(cards)+i] != card {
			t.Errorf("Expected card %v at index %d, but got %v", card, len(d)-len(cards)+i, d[len(d)-len(cards)+i])
		}
	}
}

func TestAddToTop(t *testing.T) {
	d := Deck{
		Card{Name: "Two of Spades", Rank: 2},
		Card{Name: "Three of Spades", Rank: 3},
		Card{Name: "Four of Spades", Rank: 4},
	}

	card := Card{Name: "Ace of Spades", Rank: 14}
	expectedDeckSize := len(d) + 1

	d.AddToTop(card)

	if len(d) != expectedDeckSize {
		t.Errorf("Expected deck size of %d, but got %d", expectedDeckSize, len(d))
	}

	if d[0] != card {
		t.Errorf("Expected first card to be %v, but got %v", card, d[0])
	}
}

func TestAddToTopMultipleCards(t *testing.T) {
	d := Deck{
		Card{Name: "Two of Spades", Rank: 2},
		Card{Name: "Three of Spades", Rank: 3},
		Card{Name: "Four of Spades", Rank: 4},
	}

	cards := []Card{
		{Name: "Ace of Spades", Rank: 14},
		{Name: "King of Spades", Rank: 13},
		{Name: "Queen of Spades", Rank: 12},
	}
	expectedDeckSize := len(d) + len(cards)

	d.AddToTop(cards...)

	if len(d) != expectedDeckSize {
		t.Errorf("Expected deck size of %d, but got %d", expectedDeckSize, len(d))
	}

	for i, card := range cards {
		if d[i] != card {
			t.Errorf("Expected card %v at index %d, but got %v", card, i, d[i])
		}
	}
}