package main

func main() {
	cards := newDeckFromFile("bad file")
	cards.print()
}
