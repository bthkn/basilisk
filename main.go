package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/bthkn/basilisk/internal"
	"github.com/bthkn/basilisk/scryfall"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func serve() {
	// http.HandleFunc("/", handler)
	// log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	args := os.Args[1:]
	fmt.Println(args)

	card, err := GetCard("afr", "1", "en")
	card, err := scryfall.GetCard("afr", "1", "en")
	if err != nil {
		// handle error
		log.Fatal(err)
	}

	fmt.Println(card.Name)

	cards, err := SearchCard("Vizier of Tumbling Sands", "en")
	cards, err := scryfall.SearchCard("Vizier of Tumbling Sands", "en")
	if err != nil {
		// handle error
		log.Fatal(err)
	}

	fmt.Println(cards.Data[0].Name)
}
