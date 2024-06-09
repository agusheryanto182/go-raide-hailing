package helper

import (
	"math"

	"github.com/agusheryanto182/go-raide-hailing/module/feature/purchase/dto"
	"github.com/agusheryanto182/go-raide-hailing/utils/customErr"
)

func Haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371 // Earth radius in kilometers
	dLat := (lat2 - lat1) * math.Pi / 180.0
	dLon := (lon2 - lon1) * math.Pi / 180.0
	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1*math.Pi/180.0)*math.Cos(lat2*math.Pi/180.0)*
			math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return R * c
}

func CalculateTotalPriceAndDeliveryTime(merchant []*dto.ResEstimateMerchant, items []*dto.ResEstimateItem, userLat, userLong float64) (int, int, error) {
	totalPrice := 0
	currDistance := 0.0

	for _, val := range merchant {
		// Calculate distance from user location to merchant
		distance := Haversine(userLat, userLong, val.LocationLat, val.LocationLong)

		maxDistance := math.Sqrt(3) // 1.7320508075688772km
		if distance > maxDistance {
			return 0, 0, customErr.NewBadRequestError("merchant is too far away")
		}

		if distance > currDistance {
			currDistance = distance
		}
	}

	for _, val := range items {
		totalPrice += val.TotalPrice
	}

	// Calculate delivery time based on total distance and speed
	speed := 40.0                                    // km/h
	deliveryTime := int((currDistance / speed) * 60) // in minutes

	return totalPrice, deliveryTime, nil
}
