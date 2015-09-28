package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"
)

var (
	server *httptest.Server
	gamesUrl string
)

func setup() {
	server = httptest.NewServer(GetRouter())
	gamesUrl = fmt.Sprintf("%s/games", server.URL)
}

func TestPostGameReturnsOK(t *testing.T) {
	setup()

	postString := ""

	reader := strings.NewReader(postString)

	request, err := http.NewRequest("POST", gamesUrl, reader)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 201 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}