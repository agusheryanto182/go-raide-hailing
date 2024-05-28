package dto

type ReqNearbyMerchants struct {
	Lat              float64
	Lon              float64
	MerchantId       string `json:"merchantId"`
	Name             string `json:"name"`
	MerchantCategory string `json:"merchantCategory"`
	Limit            int    `json:"limit"`
	Offset           int    `json:"offset"`
}
