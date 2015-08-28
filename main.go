package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/games", CreateGame).Methods("POST")
    router.HandleFunc("/games", ListGames).Methods("GET")
    log.Fatal(http.ListenAndServe(":8080", router))
}

func CreateGame(response http.ResponseWriter, request *http.Request) {
    // generate new id
    // add to list of known ids
    // return id in response header
    fmt.Fprintf(response, "Done")
}

func ListGames(response http.ResponseWriter, request *http.Request) {
    fmt.Fprintf(response, "Done")
}
