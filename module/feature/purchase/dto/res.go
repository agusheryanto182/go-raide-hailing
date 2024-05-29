package dto

type TempItems struct {
	ID              string `json:"item_id"`
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
