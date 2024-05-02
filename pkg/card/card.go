package card

type Card struct {
	Name string
	Rank int
}

// newCard creates a new card with the given suit and value.
func NewCard(name string, rank int) Card {
	return Card{Name: name, Rank: rank}
}
