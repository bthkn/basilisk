package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bthkn/basilisk/scryfall"
)

func main() {
	args := os.Args[1:]
	fmt.Println(args)

	card, err := scryfall.GetCard("afr", "1", "en")
	if err != nil {
		// handle error
		log.Fatal(err)
	}

	fmt.Println(card.Name)

	cards, err := scryfall.SearchCard("Vizier of Tumbling Sands", "en")
	if err != nil {
		// handle error
		log.Fatal(err)
	}

	fmt.Println(cards.Data[0].Name)
}
