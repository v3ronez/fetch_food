package data

import (
	"database/sql"
	"time"
)

// s/\<./\u&/g --> capitalize all words selecteds
type Food struct {
	Code            int32     `json:"code"`
	Status          string    `json:"status"`
	ImportedAt      time.Time `json:"imported_at"`
	Url             string    `json:"url"`
	Creator         string    `json:"creator"`
	CreatedAt       int32     `json:"created_at"`
	LastModifiedAt  time.Time `json:"last_modified_at"`
	ProductName     string    `json:"product_name"`
	Quantity        string    `json:"quantity"`
	Brands          string    `json:"brands"`
	Categories      string    `json:"categories"`
	Labels          string    `json:"labels"`
	Cities          string    `json:"cities"`
	PurchasePlaces  string    `json:"purchase_places"`
	Stores          string    `json:"stores"`
	IngredientsText string    `json:"ingredients_text"`
	Traces          string    `json:"traces"`
	ServingSize     string    `json:"serving_size"`
	ServingQuantity float32   `json:"serving_quantity"`
	NutriscoreScore int32     `json:"nutriscore_score"`
	NutriscoreGrade string    `json:"nutriscore_grade"`
	MainCategory    string    `json:"main_category"`
	ImageuRrl       string    `json:"image_url"`
}

type FoodModel struct {
	DB *sql.DB
}

func (f *FoodModel) Get() (Food, error) {
	return Food{}, nil
}
