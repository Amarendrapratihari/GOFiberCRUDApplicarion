package models

type Book struct {
	Id          int    `json:"id" gorm:"primaryKey"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Price       int    `json:"price"`
	Description string `json:"description"`
}
