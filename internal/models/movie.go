package models

type Movie struct {
	Id    uint8  `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`
}
