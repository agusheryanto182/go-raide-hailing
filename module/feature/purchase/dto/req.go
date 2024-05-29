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

type ReqPostEstimate struct {
	UserLocation struct {
		Lat  float64 `json:"lat" validate:"required"`
		Long float64 `json:"long" validate:"required"`
	} `json:"userLocation"`
	Orders []struct {
		MerchantId      string `json:"merchantId" validate:"required"`
		IsStartingPoint bool   `json:"isStartingPoint" validate:"required"`
		Items           []struct {
			ItemId   string `json:"itemId" validate:"required"`
			Quantity int    `json:"quantity" validate:"required"`
		}
	}
	UserId string
}

type ItemParams struct {
	ItemIDParams []string
	MerchantId   []string
	Quantity     []int
}
