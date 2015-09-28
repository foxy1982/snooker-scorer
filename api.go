package main

import (
    "fmt"
    "net/http"
    
    "github.com/gorilla/mux"
)

func GetRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/games", CreateGame).Methods("POST")
    router.HandleFunc("/games", ListGames).Methods("GET")
    return router
}

func CreateGame(response http.ResponseWriter, request *http.Request) {
    gameId := Create()
    // generate new id
    // add to list of known ids
    // return id in response header
    response.WriteHeader(http.StatusCreated)
    fmt.Fprintf(response, fmt.Sprintf("%d", gameId))
}

func ListGames(response http.ResponseWriter, request *http.Request) {
    fmt.Fprintf(response, "Done")
}
