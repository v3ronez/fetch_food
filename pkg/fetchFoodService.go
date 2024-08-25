package pkg

//
// import (
// 	"fmt"
// )
//
// const BASE_URL string = "https://challenges.coode.sh/food/data/json/"
//
// type FetchFoodService struct {
// 	client *HttpClientService
// }
//
// func (f *FetchFoodService) New(clientHttp *HttpClientService) *FetchFoodService {
// 	return &FetchFoodService{
// 		client: clientHttp,
// 	}
// }
//
// func (f *FetchFoodService) FetchProducts() ([]Food, error) {
// 	urlFetchFiles := fmt.Sprintf("%s%s", BASE_URL, "index.txt")
// 	resp, err := f.client.Get(urlFetchFiles)
// 	if err != nil {
// 		return nil, err
// 	}
// }
