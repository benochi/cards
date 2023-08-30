package main

import "fmt"

func main() {
	cards := []string{newCard(), newCard(), "Ace of Diamonds"}
	cards = append(cards, "New Element")
	//iterate over slice
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}