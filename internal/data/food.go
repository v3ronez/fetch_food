package data

import (
	"database/sql"
	"time"
)

// s/\<./\u&/g --> capitalize all words selecteds
type Food struct {
	Code             int32
	Status           string
	Imported_t       time.Time
	Url              string
	Creator          string
	Created_t        int32
	Last_modified_t  time.Time
	Product_name     string
	Quantity         string
	Brands           string
	Categories       string
	Labels           string
	Cities           string
	Purchase_places  string
	Stores           string
	Ingredients_text string
	Traces           string
	Serving_size     string
	Serving_quantity float32
	Nutriscore_score int32
	Nutriscore_grade string
	Main_category    string
	Image_url        string
}

type FoodModel struct {
	DB *sql.DB
}

func (f *FoodModel) Get() (Food, error) {
	return Food{}, nil
}
