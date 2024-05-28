package dto

type ReqNearbyMerchants struct {
	UserLat          float64 `db:"user_lat"`
	UserLong         float64 `db:"user_long"`
	MerchantId       string  `json:"merchantId"`
	Name             string  `json:"name"`
	MerchantCategory string  `json:"merchantCategory" validate:"oneof=SmallRestaurant MediumRestaurant LargeRestaurant MerchandiseRestaurant BoothKiosk ConvenienceStore"`
	Limit            int     `json:"limit"`
	Offset           int     `json:"offset"`
}
