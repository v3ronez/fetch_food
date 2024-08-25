package data

import "database/sql"

type Models struct {
	Food FoodModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Food: FoodModel{DB: db},
	}
}
