package dto

type ReqCreateMerchant struct {
	Name             string `json:"name" validate:"required,min=5,max=30"`
	MerchantCategory string `json:"merchantCategory" validate:"required,oneof=SmallRestaurant MediumRestaurant LargeRestaurant MerchandiseRestaurant BoothKiosk ConvenienceStore"`

	Location struct {
		Latitude  float64 `json:"lat" validate:"required"`
		Longitude float64 `json:"long" validate:"required"`
	} `json:"location" validate:"required"`

	ImageUrl string `json:"imageUrl" validate:"required,imageUrl"`
}
