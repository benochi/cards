package main

func main() {
	cards := newDeck()

	//iterate over slice
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
