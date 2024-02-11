package internal

import (
	"net/http"
	"strings"

	"github.com/bthkn/basilisk/scryfall"
)

func CardHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	name := q.Get("name")
	code := q.Get("code")
	number := q.Get("number")
	lang := q.Get("lang")

	card, err := scryfall.GetCard(code, number, lang)
	if err == nil {
		http.Redirect(w, r, card.Images.Large, http.StatusFound)
		return
	}

	cardList, err := scryfall.SearchCard(name, lang)
	if err == nil {
		http.Redirect(w, r, cardList.Data[0].Images.Large, http.StatusFound)
		return
	}

	http.Redirect(w, r, getScryfallFallbackUrl(code, number), http.StatusFound)
}

func MakeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		endpoint := strings.Split(r.URL.Path, "/")[1]
		switch endpoint {
		case "card":
			fn(w, r)
		default:
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}
	}
}

func getScryfallFallbackUrl(code string, num string) string {
	return "https://api.scryfall.com/cards/" + code + "/" + num + "?format=image"
}
