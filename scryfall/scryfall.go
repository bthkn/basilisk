package scryfall

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type ScryfallResponse struct {
	Object  string `json:"object"`
	Status  int    `json:"status"`
	Details string `json:"details"`
}

type Card struct {
	ID       string    `json:"id"`
	OracleID string    `json:"oracle_id"`
	Name     string    `json:"name"`
	Lang     string    `json:"lang"`
	Images   ImageURIs `json:"image_uris"`
}

type CardList struct {
	TotalCards int    `json:"total_cards"`
	Data       []Card `json:"data"`
}

type ImageURIs struct {
	Small  string `json:"small"`
	Normal string `json:"normal"`
	Large  string `json:"large"`
	Png    string `json:"png"`
}

func GetCard(set string, num string, lang string) (Card, error) {
	url := "https://api.scryfall.com/cards/" + set + "/" + num + "/" + lang

	data, err := sendAPIRequest(url)
	if err != nil {
		return Card{}, err
	}

	var card Card
	jsonErr := json.Unmarshal(data, &card)
	if jsonErr != nil {
		return Card{}, jsonErr
	}

	return card, nil
}

func SearchCard(query string, lang string) (CardList, error) {
	searchQuery := buildSearchQuery(query, lang)
	url := "https://api.scryfall.com/cards/search?q=" + searchQuery
	data, err := sendAPIRequest(url)
	if err != nil {
		return CardList{}, err
	}

	var cardList CardList
	jsonErr := json.Unmarshal(data, &cardList)
	if jsonErr != nil {
		return CardList{}, jsonErr
	}

	return cardList, nil
}

func sendAPIRequest(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var res ScryfallResponse
	jsonErr := json.Unmarshal(body, &res)
	if jsonErr != nil {
		return nil, jsonErr
	}

	if res.Object == "error" {
		return nil, errors.New(res.Details)
	}

	return body, nil
}

func buildSearchQuery(name string, lang string) string {
	var quotedName []string

	nameParts := strings.Split(name, " ")
	for _, part := range nameParts {
		quotedName = append(quotedName, url.QueryEscape(part))
	}

	query := url.QueryEscape("!\"") +
		strings.Join(quotedName, "+") +
		url.QueryEscape("\"") + "+" +
		url.QueryEscape("lang:"+lang)

	return query
}
