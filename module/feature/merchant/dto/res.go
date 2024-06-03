package dto

import "github.com/agusheryanto182/go-raide-hailing/module/entities"

type Meta struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

type ResGetMerchant struct {
	ID               string            `db:"id" json:"merchantId"`
	Name             string            `db:"name" json:"name"`
	MerchantCategory string            `db:"merchant_category" json:"merchantCategory"`
	ImageUrl         string            `db:"image_url" json:"imageUrl"`
	Location         entities.Location `db:"location" json:"location"`
	CreatedAt        string            `db:"created_at" json:"createdAt"`
}
