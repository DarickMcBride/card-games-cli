package card

import "testing"

func TestNewCard(t *testing.T) {
	name := "Ace of Spades"
	rank := 14

	c := NewCard(name, rank)

	if c.Name != name {
		t.Errorf("Expected card name to be %s, but got %s", name, c.Name)
	}

	if c.Rank != rank {
		t.Errorf("Expected card rank to be %d, but got %d", rank, c.Rank)
	}
}
