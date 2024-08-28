package pkg

import (
	"fmt"
	"log/slog"
)

const BASE_URL string = "https://challenges.coode.sh/food/data/json/"

type FoodService struct {
	httpClient HttpService
}

func NewFoodService(client HttpService) *FoodService {
	return &FoodService{
		httpClient: client,
	}
}
func (f *FoodService) CheckForNewFiles() {
	url := fmt.Sprintf("%s%s", BASE_URL, "index.txt")

	response, err := f.httpClient.Get(url)
	if err != nil {
		slog.Error("Somethings happen: ", err)
	}
	files := []string{string(response)}
	for _, filePath := range files {
		f.downLoadNewFiles(filePath)
	}
}

func (f *FoodService) downLoadNewFiles(urlPath string) {
	fmt.Println(urlPath)
}
