package main

import (
	"fmt"
	"os"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardsSuite := deck{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardsValues := deck{"Ace", "Two", "Three", "Four"}

	for _, suite := range cardsSuite {
		for _, value := range cardsValues {
			cards = append(cards, suite+" of "+value)
		}
	}

	return cards
}

func (d deck) print() {
	for i, vd := range d {
		fmt.Println(i, vd)
	}
}

func (d deck) toString() string {
	result := ""
	for _, de := range d {
		result += de
	}
	return result
}

func deal(decks deck, handSize int) (deck, deck) {
	return decks[:handSize], decks[handSize:]
}

func writeToFile(decks deck) (bool, error) {
	filePath := "/Users/srinivaskarre/scalekit/code/go-play/decks.txt"
	_, err := os.ReadFile(filePath)

	if err != nil {
		_, err := os.Create(filePath)
		validateError(err)
		err = os.WriteFile(filePath, []byte(decks.toString()), 0644)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

func validateError(err error) {
	if err != nil {
		panic(err)
	}
}
