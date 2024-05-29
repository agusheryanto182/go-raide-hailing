package helper

import (
	"math"

	"github.com/agusheryanto182/go-raide-hailing/module/feature/purchase/dto"
	"github.com/agusheryanto182/go-raide-hailing/utils/customErr"
)

func CalculateDistance(lat1, lon1, lat2, lon2 float64) (float64, error) {
	const R = 6371 // Radius of the Earth in km
	dLat := (lat2 - lat1) * (math.Pi / 180)
	dLon := (lon2 - lon1) * (math.Pi / 180)

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1*(math.Pi/180))*math.Cos(lat2*(math.Pi/180))*
			math.Sin(dLon/2)*math.Sin(dLon/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distance := R * c
	if distance > 3 {
		return 0, customErr.NewBadRequestError("distance is too large")
	}

	return distance, nil
}

func CalculateTotalPriceAndDeliveryTime(merchant []*dto.ResEstimateMerchant, items []*dto.ResEstimateItem, userLat, userLong float64) (int, int, error) {
	totalPrice := 0
	currDistance := 0.0

	for _, val := range merchant {
		// Calculate distance from user location to merchant
		distance, err := CalculateDistance(userLat, userLong, val.LocationLat, val.LocationLong)
		if err != nil {
			return 0, 0, err
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
