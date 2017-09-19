package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
type deck []string

// Create a new deck of cards
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Hearts", "Clubs", "Diamonds", "Spades"}
	cardValues := []string{
		"Ace", "Two", "Three", "Four", "Five",
		"Six", "Seven", "Eight", "Nine",
		"Ten", "Jack", "Queen", "King",
	}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Print current deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Deal a deck of cards with custom hand size
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Render current deck to string format
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// Save current deck to file (relative to file path)
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// Create a new deck from existing file
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

// Shuffle the current deck
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
