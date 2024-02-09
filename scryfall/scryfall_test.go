package scryfall

import (
	"testing"
)

func TestBuildSearchQuery(t *testing.T) {
	type CardQuery struct {
		name  string
		query string
	}

	testCases := []CardQuery{
		{"Vizier of Tumbling Sands", "%21%22Vizier+of+Tumbling+Sands%22+lang%3Aru"},
		{"+2 Mace", "%21%22%2B2+Mace%22+lang%3Aru"},
	}

	for _, cq := range testCases {
		result := buildSearchQuery(cq.name, "ru")

		if result != cq.query {
			t.Fatalf(`search query for "%s" is %q, wanted: %#q`, cq.name, result, cq.query)
		}
	}
}
