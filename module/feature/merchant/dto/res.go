package dto

import "github.com/agusheryanto182/go-raide-hailing/module/entities"

type Meta struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

type ResGetMerchant struct {
	ID               string            `db:"id" json:"id"`
	Name             string            `db:"name" json:"name"`
	MerchantCategory string            `db:"merchant_category" json:"merchant_category"`
	ImageUrl         string            `db:"image_url" json:"image_url"`
	Location         entities.Location `db:"location" json:"location"`
	CreatedAt        string            `db:"created_at" json:"created_at"`
}
