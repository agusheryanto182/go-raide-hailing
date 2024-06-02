package dto

import (
	"github.com/agusheryanto182/go-raide-hailing/module/entities"
)

type TempItems struct {
	ID              string `json:"item_id"`
	MerchantId      string `json:"merchant_id"`
	Name            string `json:"name"`
	ProductCategory string `json:"product_category"`
	Price           int    `json:"price"`
	ImageUrl        string `json:"image_url"`
	CreatedAt       string `json:"created_at"`
}

type ResPostEstimate struct {
	TotalPrice                     int    `db:"total_price" json:"totalPrice"`
	EstimatedDeliveryTimeInMinutes int    `db:"estimated_delivery_time_in_minutes" json:"estimatedDeliveryTimeInMinutes"`
	CalculatedEstimateId           string `db:"estimate_id" json:"calculatedEstimateId"`
}

type ResEstimateMerchant struct {
	MerchantId   string  `db:"merchant_id" json:"merchantId"`
	LocationLat  float64 `db:"location_lat" json:"locationLat"`
	LocationLong float64 `db:"location_long" json:"locationLong"`
}

type ResEstimateItem struct {
	ItemId     string `db:"item_id" json:"itemId"`
	MerchantId string `db:"merchant_id" json:"merchantId"`
	TotalPrice int    `db:"total_price" json:"totalPrice"`
}

type ResGetOrders struct {
	OrderId string `db:"order_id" json:"orderId"`
	Orders  []*Orders
}

type Orders struct {
	Merchant *entities.Merchant `json:"merchants"`
	Items    []*Items           `json:"items"`
}

type Items struct {
	ID              string `db:"item_id" json:"itemId"`
	Name            string `db:"name" json:"name"`
	ProductCategory string `db:"product_category" json:"productCategory"`
	Price           int    `db:"price" json:"price"`
	Quantity        int    `json:"quantity"`
	ImageUrl        string `db:"image_url" json:"imageUrl"`
	CreatedAt       string `db:"created_at" json:"createdAt"`
}

type TempMerchant struct {
	ID               string  `db:"merchant_id" json:"merchant_id"`
	Name             string  `db:"name" json:"name"`
	MerchantCategory string  `db:"merchant_category" json:"merchant_category"`
	ImageUrl         string  `db:"image_url" json:"image_url"`
	LocationLat      float64 `db:"location_lat" json:"location_lat"`
	LocationLong     float64 `db:"location_long" json:"location_long"`
	CreatedAt        string  `db:"created_at" json:"created_at"`
}

type TempOrders struct {
	OrderId    string `db:"order_id" json:"order_id"`
	EstimateId string `db:"estimate_id" json:"estimate_id"`
	UserId     string `db:"user_id" json:"user_id"`
}

type TempOrderItems struct {
	OrderItemId    string `db:"order_item_id" json:"order_item_id"`
	EstimateId     string `db:"estimate_id" json:"estimate_id"`
	OrderId        string `db:"order_id" json:"order_id"`
	MerchantId     string `db:"merchant_id" json:"merchant_id"`
	MerchantItemId string `db:"merchant_item_id" json:"merchant_item_id"`
	Price          int    `db:"price" json:"price"`
	Quantity       int    `db:"quantity" json:"quantity"`
	CreatedAt      string `db:"created_at" json:"created_at"`
}
