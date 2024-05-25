package entities

type Merchant struct {
	ID               string    `db:"id"`
	Name             string    `db:"name"`
	MerchantCategory string    `db:"merchant_category"`
	ImageUrl         string    `db:"image_url"`
	Location         []float64 `db:"location"`
	CreatedAt        string    `db:"created_at"`
}

type Location struct {
	Latitude  float64 `db:"lat"`
	Longitude float64 `db:"long"`
}

type MerchantItem struct {
	ID              string `db:"id"`
	MerchantID      string `db:"merchant_id"`
	Name            string `db:"name"`
	ProductCategory string `db:"product_category"`
	Price           int    `db:"price"`
	ImageUrl        string `db:"image_url"`
	CreatedAt       string `db:"created_at"`
}
