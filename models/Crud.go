package models

type Post struct {
	TDN         int    `json:"tdn" gorm:"primary_key"`
	TrainName   string `json:"train_name"`
	TrainType   string `json:"train_type"`
	TrainNumber string `json:"train_number"`
}

type Create struct {
	TDN         int    `json:"tdn" binding:"required"`
	TrainName   string `json:"train_name" binding:"required"`
	TrainType   string `json:"train_type" binding:"required"`
	TrainNumber string `json:"train_number" binding:"required"`
}
type Update struct {
	TrainName   string `json:"train_name"`
	TrainType   string `json:"train_type"`
	TrainNumber string `json:"train_number"`
}
