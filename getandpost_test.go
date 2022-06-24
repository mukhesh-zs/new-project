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
		input         string
		expstatuscode int
		expoutput     book
	}{
		{"valid id", "1", http.StatusOK, book{"2states", author{"chetan", "bhagat", "23-04-1980", "max"}, "arihant", "23-10-2020"}},
		{"invalid id", "-1", http.StatusBadRequest, book{"2states", author{"chetan", "bhagat", "23-04-1980", "max"}, "arihant", "23-10-2020"}},
	}
	for _, tc := range testcase {
		req := httptest.NewRequest("GET", "/book"+tc.input, nil)
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
		{"getting data of all books", "/book", []book{{"2states", "chetan", "arihant", "23-10-2020"}, {"3states", "mukhesh", "GKpublications", "03-7-2000"}}},
	}
	for _, tc := range testcases {
		req := httptest.NewRequest("GET", "/book", nil)
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
		{"invalid input", book{"", author{"mukhesh", "mekala", "24-10-2000", "zero"}, "GKpublications", "03-7-2015"}, http.StatusCreated},
	}
	for _, tc := range testcases {
		data, _ := json.Marshal(tc.input)
		req := httptest.NewRequest("POST", "/book", bytes.NewBuffer(data))
		w := httptest.NewRecorder()
		PostBook(w, req)
		res := w.Result()
		assert.Equal(t, tc.expstatuscode, res.StatusCode)
	}

}
