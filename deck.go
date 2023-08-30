package main

import "fmt"

//create a new type of deck
//which is a slice of strings
type deck []string

//create function for instances to use as method. function has a receiver.
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
