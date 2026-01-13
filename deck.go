package main

import (
	"fmt"
	"os"
	"strings"
)

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func NewDeck() deck{
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
  
    for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}  
	
	return cards
}

func deal(d deck, handSize int) (deck, deck){
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}


func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func readDeckFile (filename string) deck {
      value, err := os.ReadFile(filename)
       
	  if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	  }

	  s := strings.Split(string(value), ", ")

      return deck(s)
}