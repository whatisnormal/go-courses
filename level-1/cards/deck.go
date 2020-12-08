package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type 'deck'
// which is a slice of strings
// 'type' extends / lends behavior from type []string = slice of strings
type deck []string

/*
Anytime someones call´s newDeck, it returns a
value of type deck.
*/
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// When there is a variable we do not need to use it, then replace it with '_'. In this case 'i' or 'j'
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Any variable of type 'deck' has access to print method.

// d = reference to the instance of the current 'deck'
// deck = the type of the instance.
// d receiver is similar to 'this' on java or 'self' in python
// by convention, receiver has a one or two letter abreviation.

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

/*
	Receives a deck and handSize and returns 2 decks.
*/
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	//fmt.Println(d)
	return strings.Join(d, ",")
}

func (d deck) saveToFile(filename string) error {
	//0666 permission sets = 'everyone can read and write'
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return strings.Split(string(bs), ",")
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())

	r := rand.New(source)

	// Only iterates over the index.
	for i := range d {
		// Returns a random number between 0 and -1 the lenght of the slice.
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
