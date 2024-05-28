package entities

type Purchase struct {
	ID            string `db:"id" json:"id"`
	UserId        string `db:"user_id" json:"user_id"`
	MerchantId    string `db:"merchant_id" json:"merchant_id"`
	ProductItemId string `db:"product_item_id" json:"product_item_id"`
}

type NearbyMerchant struct {
	Id       string         `db:"id" json:"id"`
	Merchant Merchant       `json:"merchant"`
	Items    []MerchantItem `json:"items"`
}
