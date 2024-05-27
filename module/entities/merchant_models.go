package entities

type Merchant struct {
	ID               string    `db:"id" json:"id"`
	Name             string    `db:"name" json:"name"`
	MerchantCategory string    `db:"merchant_category" json:"merchant_category"`
	ImageUrl         string    `db:"image_url" json:"image_url"`
	Location         []float64 `db:"location" json:"location"`
	CreatedAt        string    `db:"created_at" json:"created_at"`
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
