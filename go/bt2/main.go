package main

import (
	"fmt"
	"time"
)

func main() {
	gameChannel := make(chan []Game)

	go fetchGames(gameChannel)

	select {

	case games := <-gameChannel:

		fmt.Println("Top Games:")

		for i := 0; i < 5; i++ {
			fmt.Println(games[i].Name)
		}

	case <-time.After(5 * time.Second):

		fmt.Println("Request timeout")
	}
}
