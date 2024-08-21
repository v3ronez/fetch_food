package main

import (
	"fmt"
	"log/slog"

	HttpService "github.com/v3ronez/fetch_food/pkg"
)

func main() {
	// TODO: move this to a another service
	url := "https://challenges.coode.sh/food/data/json/index.txt"
	httpClient := HttpService.New()
	response, err := httpClient.Get(url)
	if err != nil {
		slog.Error("Somethings happen: ", err)
	}
	fmt.Println(string(response))
}
