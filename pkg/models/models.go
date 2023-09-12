package models

import "time"

type User struct {
	Id        int64     `json:"id" bson:"id"`
	Name      string    `json:"name" bson:"name"`
	Mobile    string    `json:"mobile" bson:"mobile"`
	Latitude  float64   `json:"latitude" bson:"latitude"`
	Longitude float64   ` json:"longitude" bson:"longitude"`
	CreatedAt time.Time `json:"created_at" bson:"created_At"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}
type Product struct {
	ProductID               int64     `json:"product_id" bson:"product_id"`
	ProductName             string    `json:"product_name" validate:"required" bson:"product_name"`
	ProductDescription      string    `json:"product_description" bson:"product_description" validate:"required"`
	ProductImages           []string  `json:"product_images" bson:"product_images" validate:"required"`
	ProductPrice            int       `json:"product_price" bson:"product_price" validate:"required"`
	CompressedProductImages []string  `json:"compressed_product_images" bson:"compressed_product_images"`
	Created_At              time.Time `json:"created_at" bson:"created_at"`
	Updated_At              time.Time `json:"updated_at" bson:"updated_at"`
}
