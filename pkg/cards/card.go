package deck

type Card struct {
	Name string
	Rank int
	Suit string
}

// newCard creates a new card with the given suit and value.
func NewCard(name string, rank int, suit string) Card {
	return Card{Name: name, Rank: rank, Suit: suit}
}
