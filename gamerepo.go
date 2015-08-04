package main

var games []int

func init() {
    games = append(games, 0)
}

func Create() int {
    max := games[0]

    for _, value := range games {
        if (value > max) {
            max = value
        }
    }

    newGameId := max + 1

    games = append(games, newGameId)

    return newGameId
}

