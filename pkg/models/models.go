package models

import "time"

type User struct {
	Id        int       `json:"id" bson:"id"`
	Name      string    `json:"name" bson:"name"`
	Mobile    string    `json:"mobile" bson:"mobile"`
	Latitude  float64   `json:"latitude" bson:"latitude"`
	Longitude float64   ` json:"longitude" bson:"longitude"`
	CreatedAt time.Time `json:"created_at" bson:"created_At"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}
type Product struct {
	ProductID               int       `json:"product_id"`
	ProductName             string    `json:"product_name"`
	ProductDescription      string    `json:"product_description"`
	ProductImages           []string  `json:"product_images"`
	ProductPrice            int       `json:"product_price"`
	CompressedProductImages []string  `json:"compressed_product_images"`
	CreatedAt               time.Time `json:"created_at"`
	UpdatedAt               time.Time `json:"updated_at"`
}
