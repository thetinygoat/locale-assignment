package postgres

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/thetinygoat/localeai/pkg/entities/ride"
)

// Repository implements ride.Repository
type Repository struct {
	conn *pgx.Conn
}

// New instantiates a new postgres repository
func New(conn *pgx.Conn) ride.Repository {
	return Repository{conn: conn}
}

// CreateRide creates a ride
func (r Repository) CreateRide(ride *ride.Ride) error {
	ctx := context.Background()
	stmt, err := r.conn.Prepare(
		ctx,
		"create",
		`INSERT INTO 
		rides(
			id,
			user_id,
			vehicle_model_id,
			package_id,
			travel_type_id,
			from_area_id,
			to_area_id,
			from_city_id,
			to_city_id,
			from_date,
			to_date,
			online_booking,
			mobile_site_booking,
			booking_created,
			from_lat,
			from_long,
			to_lat,
			to_long,
			car_cancellation,
			job_id,
			created_at,
			updated_at) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22)`)

	if err != nil {
		return err
	}
	_, err = r.conn.Exec(ctx, stmt.Name,
		ride.ID,
		ride.UserID,
		ride.VehicleModelID,
		ride.PackageID,
		ride.TravelTypeID,
		ride.FromAreaID,
		ride.ToAreaID,
		ride.FromCityID,
		ride.ToCityID,
		ride.FromDate,
		ride.ToDate,
		ride.OnlineBooking,
		ride.MobileSiteBooking,
		ride.BookingCreated,
		ride.FromLat,
		ride.FromLong,
		ride.ToLat,
		ride.ToLong,
		ride.CarCancellation,
		ride.JobID,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return err
	}
	return nil
}
