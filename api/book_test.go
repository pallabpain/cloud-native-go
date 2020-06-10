package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookToJSON(t *testing.T) {
	book := Book{Title: "My Book", Author: "Pallab Pain", ISBN: "0123456789"}
	json := book.ToJSON()
	assert.Equal(t, `{"title":"My Book","author":"Pallab Pain","isbn":"0123456789"}`, string(json), "Book JSON marshalling incorrect")
}

func TestBookFromJSON(t *testing.T) {
	json := []byte(`{"title":"My Book","author":"Pallab Pain","isbn":"0123456789"}`)
	actualBook := FromJSON(json)
	expectedBook := Book{Title: "My Book", Author: "Pallab Pain", ISBN: "0123456789"}
	assert.Equal(t, expectedBook, actualBook, "Book unmarshalling incorrect")
}
