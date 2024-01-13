package main

import (
	"fmt"
)

func main() {
	//var card = "Ace of cards"
	card := getCard()
	fmt.Println(card)

	//slice
	cards := newDeck()
	dealCard, remainingCards := deal(cards, 5)

	dealCard.print()
	remainingCards.print()

	writeToFile(cards)

	//play pointer
	a := 1
	var b *int

	if b == nil {
		fmt.Println("b is nil", b)
		b = &a
		fmt.Println("b is ", *b)
	}
}

func getCard() string {
	return "Ace-A"
}
