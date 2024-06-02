package entities

type Estimate struct {
	EstimateId            string  `db:"estimate_id" json:"estimateId"`
	UserId                string  `db:"user_id" json:"userId"`
	UserLat               float64 `db:"user_lat" json:"userLat"`
	UserLon               float64 `db:"user_lon" json:"userLon"`
	TotalPrice            int     `db:"total_price" json:"totalPrice"`
	EstimatedDeliveryTime int     `db:"estimated_delivery_time" json:"estimatedDeliveryTime"`
	CreatedAt             string  `db:"created_at" json:"createdAt"`
}

type NearbyMerchant struct {
	Merchant struct {
		ID               string   `db:"merchant_id" json:"merchantId"`
		Name             string   `db:"name" json:"name"`
		MerchantCategory string   `db:"merchant_category" json:"merchantCategory"`
		ImageUrl         string   `db:"image_url" json:"imageUrl"`
		Location         Location `json:"location"`
		Distance         float64  `json:"-"`
		CreatedAt        string   `db:"created_at" json:"createdAt"`
	} `json:"merchant"`
	Items []Item `json:"items"`
}

type Item struct {
	ID              string `json:"itemId"`
	Name            string `json:"name"`
	ProductCategory string `json:"productCategory"`
	Price           int    `json:"price"`
	ImageUrl        string `json:"imageUrl"`
	CreatedAt       string `json:"createdAt"`
}

type OrderItems struct {
	OrderItemId    string `db:"order_item_id" json:"orderItemId"`
	OrderId        string `db:"order_id" json:"orderId"`
	MerchantId     string `db:"merchant_id" json:"merchantId"`
	MerchantItemId string `db:"merchant_item_id" json:"merchantItemId"`
	Price          int    `db:"price" json:"price"`
	Quantity       int    `db:"quantity" json:"quantity"`
	CreatedAt      string `db:"created_at" json:"createdAt"`
}
