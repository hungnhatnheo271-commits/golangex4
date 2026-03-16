package main

import (
	"encoding/json"
	"net/http"
)

type Game struct {
	Name string `json:"name"`
}

type Response struct {
	Results []Game `json:"results"`
}

func fetchGames(ch chan []Game) {

	resp, err := http.Get("https://api.rawg.io/api/games")
	if err != nil {
		ch <- nil
		return
	}

	defer resp.Body.Close()

	var result Response

	json.NewDecoder(resp.Body).Decode(&result)

	ch <- result.Results
}