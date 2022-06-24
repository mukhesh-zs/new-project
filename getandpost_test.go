package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetOne(t *testing.T) {
	testcase := []struct {
		desc          string
		inputurl      string
		expstatuscode int
		expoutput     book
	}{
		{"valid id", "books/{1}", http.StatusOK, book{"2states", author{"chetan", "bhagat", "23-04-1980", "max"}, "arihant", "23-10-2020"}},
		{"invalid id", "{-1}", http.StatusBadRequest, book{}},
	}
	for _, tc := range testcase {
		req := httptest.NewRequest("GET", "/books"+tc.inputurl, nil)
		w := httptest.NewRecorder()
		GetOne(w, req)
		res := w.Result()
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		var b book
		json.Unmarshal(body, &b)

		assert.Equal(t, tc.expoutput, b)

	}

}
func TestGetAll(t *testing.T) {
	testcases := []struct {
		desc       string
		url        string
		expetedout []book
	}{
		{"getting data of all books", "/books", []book{{"2states", author{"chetan", "bhagat", "30-12-1980", "neon"}, "arihant", "23-10-2020"},
			{"3states", author{"mukhesh", "chandra", "24-10-2000", "rolex"}, "GKpublications", "03-7-2000"}},
		},
		{"getting data of all books", "/books?title=2states", []book{{"2states", author{"chetan", "bhagat", "30-12-1980", "neon"}, "arihant", "23-10-2020"}}},
	}
	for _, tc := range testcases {
		req := httptest.NewRequest("GET", "/books", nil)
		w := httptest.NewRecorder()
		GetAll(w, req)
		res := w.Result()
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		var b book
		json.Unmarshal(body, &b)

		assert.Equal(t, tc.expetedout, b)

	}

}

func TestPostBook(t *testing.T) {
	testcases := []struct {
		desc          string
		input         book
		expstatuscode int
	}{
		{"details posted", book{"3states", author{"mukhesh", "mekala", "24-10-2000", "zero"}, "GKpublications", "03-7-2000"}, http.StatusCreated},
		{"invalid input title", book{"", author{"mukhesh", "mekala", "24-10-2000", "zero"}, "GKpublications", "03-7-2015"}, http.StatusBadRequest},
	}
	for _, tc := range testcases {
		data, _ := json.Marshal(tc.input)
		req := httptest.NewRequest("POST", "/books", bytes.NewBuffer(data))
		w := httptest.NewRecorder()
		PostBook(w, req)
		res := w.Result()
		assert.Equal(t, tc.expstatuscode, res.StatusCode)
	}

}

func TestPostAuthor(t *testing.T) {
	testcases := []struct {
		desc          string
		input         author
		expstatuscode int
	}{
		{"valid author", author{"kishore", "kumar", "16-06-1906", "rolex"}, http.StatusCreated},
		{"author name not given", author{"", "kumar", "16-06-1906", "rolex"}, http.StatusBadRequest},
	}
	for _, tc := range testcases {
		data, _ := json.Marshal(tc.input)
		req := httptest.NewRequest("POST", "/books/author", bytes.NewBuffer(data))
		w := httptest.NewRecorder()
		PostAutor(w, req)
		res := w.Result()
		assert.Equal(t, tc.expstatuscode, res.StatusCode)
	}
}
func TestPutBook(t *testing.T) {
	testcases := []struct {
		desc          string
		inputurl      string
		input         book
		expstatuscode int
	}{
		{"updating details", "/book/{1}", book{"it", author{"abdul", "kalam", "25-01-1950", "ak"}, "rkpublishers", "03-04-1999"}, http.StatusCreated},
	}
	for _, tc := range testcases {
		body, _ := json.Marshal(tc.input)
		req := httptest.NewRequest("PUT", "books/"+tc.inputurl, bytes.NewBuffer(body))
		w := httptest.NewRecorder()
		PutBook(w, req)
		res := w.Result()
		defer res.Body.Close()
		assert.Equal(t, tc.expstatuscode, res.StatusCode)
	}
}
func TestPutAuthor(t *testing.T) {
	testcases := []struct {
		desc          string
		inputurl      string
		input         book
		expstatuscode int
	}{
		{"updating details", "/author/{1}", book{"it", author{"abdul", "kalam", "25-01-1950", "ak"}, "rkpublishers", "03-04-1999"}, http.StatusCreated},
	}
	for _, tc := range testcases {
		body, _ := json.Marshal(tc.input)
		req := httptest.NewRequest("PUT", "books/"+tc.inputurl, bytes.NewBuffer(body))
		w := httptest.NewRecorder()
		PutAuthor(w, req)
		res := w.Result()
		defer res.Body.Close()
		assert.Equal(t, tc.expstatuscode, res.StatusCode)
	}
}
