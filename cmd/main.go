package main

import "fmt"

func main() {

	//read from file
	deck, err := readDeckFromFile("my_cards.json")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	deck.print()

}
