package entities

type Purchase struct {
	ID         string `db:"order_id" json:"orderId"`
	UserId     string `db:"user_id" json:"userId"`
	MerchantId string `db:"merchant_id" json:"merchantId"`
	ItemId     string `db:"item_id" json:"itemId"`
}

type NearbyMerchant struct {
	Merchant struct {
		ID               string   `db:"merchant_id" json:"merchantId"`
		Name             string   `db:"name" json:"name"`
		MerchantCategory string   `db:"merchant_category" json:"merchantCategory"`
		ImageUrl         string   `db:"image_url" json:"imageUrl"`
		Location         Location `json:"location"`
		CreatedAt        string   `db:"created_at" json:"createdAt"`
	} `json:"merchant"`
	Items []struct {
		ID              string `json:"itemId"`
		Name            string `json:"name"`
		ProductCategory string `json:"productCategory"`
		Price           int    `json:"price"`
		ImageUrl        string `json:"imageUrl"`
		CreatedAt       string `json:"createdAt"`
	} `json:"items"`
}
