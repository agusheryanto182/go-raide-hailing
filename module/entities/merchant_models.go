package entities

type Merchant struct {
	ID               string  `db:"id"`
	Name             string  `db:"name"`
	MerchantCategory string  `db:"merchant_category" `
	ImageUrl         string  `db:"image_url"`
	LocationLat      float64 `db:"location_lat"`
	LocationLong     float64 `db:"location_long"`
	CreatedAt        string  `db:"created_at"`
}

type Location struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"long"`
}

type MerchantItem struct {
	ID              string `db:"id" json:"itemId"`
	MerchantId      string `db:"merchant_id" json:"-"`
	Name            string `db:"name" json:"name"`
	ProductCategory string `db:"product_category" json:"productCategory"`
	Price           int    `db:"price" json:"price"`
	ImageUrl        string `db:"image_url" json:"imageUrl"`
	CreatedAt       string `db:"created_at" json:"createdAt"`
}
