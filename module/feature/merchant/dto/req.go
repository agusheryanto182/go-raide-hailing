package dto

type ReqCreateMerchant struct {
	Name             string `json:"name" validate:"required,min=5,max=30"`
	MerchantCategory string `json:"merchantCategory" validate:"required,oneof=SmallRestaurant MediumRestaurant LargeRestaurant MerchandiseRestaurant BoothKiosk ConvenienceStore"`

	Location struct {
		Latitude  *float64 `json:"lat" validate:"required"`
		Longitude *float64 `json:"long" validate:"required"`
	} `json:"location" validate:"required"`

	ImageUrl string `json:"imageUrl" validate:"required,imageUrl"`
}

type ReqGetMerchantByFilters struct {
	MerchantId       string `json:"merchantId" db:"id"`
	Name             string `json:"name" db:"name"`
	MerchantCategory string `json:"merchantCategory" db:"merchant_category"`
	CreatedAt        string `json:"createdAt" db:"created_at"`
	Limit            int    `json:"limit"`
	Offset           int    `json:"offset"`
}

type ReqCreateMerchantItem struct {
	MerchantId      string `json:"merchantId" validate:"required"`
	Name            string `json:"name" validate:"required,min=2,max=30"`
	ProductCategory string `json:"productCategory" validate:"required,oneof=Beverage Food Snack Condiments Additions"`
	Price           int    `json:"price" validate:"required,min=1"`
	ImageUrl        string `json:"imageUrl" validate:"required,imageUrl"`
}

type ReqGetMerchantItemsByFilters struct {
	ItemId          string `json:"itemId" db:"id"`
	MerchantId      string `json:"merchantId" db:"merchant_id"`
	Name            string `json:"name" db:"name"`
	ProductCategory string `json:"productCategory" db:"product_category"`
	CreatedAt       string `json:"createdAt" db:"created_at"`
	Limit           int    `json:"limit"`
	Offset          int    `json:"offset"`
}
