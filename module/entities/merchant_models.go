package entities

type Merchant struct {
	ID               string   `db:"merchant_id" json:"merchantId"`
	Name             string   `db:"name" json:"name"`
	MerchantCategory string   `db:"merchant_category" json:"merchantCategory"`
	ImageUrl         string   `db:"image_url" json:"imageUrl"`
	Location         Location `json:"location"`
	LocationLat      float64  `db:"location_lat" json:"-"`
	LocationLong     float64  `db:"location_long" json:"-"`
	CreatedAt        string   `db:"created_at" json:"createdAt"`
}

type Location struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"long"`
}

type MerchantItem struct {
	ID              string `db:"item_id" json:"itemId"`
	MerchantId      string `db:"merchant_id" json:"-"`
	Name            string `db:"name" json:"name"`
	ProductCategory string `db:"product_category" json:"productCategory"`
	Price           int    `db:"price" json:"price"`
	ImageUrl        string `db:"image_url" json:"imageUrl"`
	CreatedAt       string `db:"created_at" json:"createdAt"`
}
