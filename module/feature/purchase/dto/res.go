package dto

type TempItems struct {
	ID              string `json:"item_id"`
	Name            string `json:"name"`
	ProductCategory string `json:"product_category"`
	Price           int    `json:"price"`
	ImageUrl        string `json:"image_url"`
	CreatedAt       string `json:"created_at"`
}
