// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type AuthOps struct {
	Login    interface{} `json:"login"`
	Register interface{} `json:"register"`
}

type NewAddress struct {
	Location   string `json:"location"`
	City       string `json:"city"`
	Phone      string `json:"phone"`
	PostalCode int    `json:"postalCode"`
	User       string `json:"user"`
}

type NewCart struct {
	Qty     int    `json:"qty"`
	Product string `json:"product"`
	User    string `json:"user"`
}

type NewCategory struct {
	Name string `json:"name"`
}

type NewProduct struct {
	Name        string  `json:"name"`
	Images      *string `json:"images"`
	Description string  `json:"description"`
	Price       int     `json:"price"`
	Discount    float64 `json:"discount"`
	Stock       int     `json:"stock"`
	Metadata    string  `json:"metadata"`
	Category    string  `json:"category"`
}

type NewProductImage struct {
	Image   string `json:"image"`
	Product string `json:"product"`
}

type NewShop struct {
	Name     string `json:"name"`
	NameSlug string `json:"nameSlug"`
	Address  string `json:"address"`
	User     string `json:"user"`
	Phone    string `json:"phone"`
}

type NewUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Role     string `json:"role"`
}

type UpdateAddress struct {
	Location   string `json:"location"`
	City       string `json:"city"`
	Phone      string `json:"phone"`
	PostalCode int    `json:"postalCode"`
}

type UpdateCart struct {
	Qty int `json:"qty"`
}

type UpdateShop struct {
	Name        string    `json:"name"`
	NameSlug    string    `json:"nameSlug"`
	Profile     string    `json:"profile"`
	Slogan      string    `json:"slogan"`
	Description string    `json:"description"`
	OpenHour    time.Time `json:"openHour"`
	CloseHour   time.Time `json:"closeHour"`
	IsOpen      bool      `json:"isOpen"`
}
