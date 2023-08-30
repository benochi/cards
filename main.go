package main

func main() {
	cards := deck{newCard(), newCard(), "Ace of Diamonds"}
	cards = append(cards, "New Element")
	//iterate over slice
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
