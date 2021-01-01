package ride

import (
	"github.com/google/uuid"
)

// Ride holds ride data
type Ride struct {
	JobID             uuid.UUID `json:"job_id"`
	ID                int64     `json:"id" `
	UserID            int64     `json:"user_id" `
	VehicleModelID    int64     `json:"vehicle_model_id" `
	PackageID         int64     `json:"package_id" `
	TravelTypeID      int64     `json:"travel_type_id" `
	FromAreaID        int64     `json:"from_area_id" `
	ToAreaID          int64     `json:"to_area_id" `
	FromCityID        int64     `json:"from_city_id" `
	ToCityID          int64     `json:"to_city_id" `
	FromDate          string    `json:"from_date" `
	ToDate            string    `json:"to_date" `
	OnlineBooking     int       `json:"online_booking" `
	MobileSiteBooking int       `json:"mobile_site_booking"`
	BookingCreated    string    `json:"booking_created"`
	FromLat           float64   `json:"from_lat" `
	FromLong          float64   `json:"from_long"`
	ToLat             float64   `json:"to_lat"`
	ToLong            float64   `json:"to_long"`
	CarCancellation   int       `json:"car_cancellation"`
}

// Repository abstracts away persistance
type Repository interface {
	CreateRide(*Ride) error
}

// Service abstracts away business logic
type Service interface {
	Dump(string)
}
