package postgres

import (
	"database/sql"
	"fmt"
	"github.com/MuhammadyusufAdhamov/booking/storage/repo"
	"github.com/jmoiron/sqlx"
)

type bookingRepo struct {
	db *sqlx.DB
}

func NewBooking(db *sqlx.DB) repo.BookingsStorageI {
	return &bookingRepo{
		db: db,
	}
}

func (ur *bookingRepo) Create(booking *repo.Booking) (*repo.Booking, error) {
	query := `
		INSERT INTO bookings(
			 room_id,
			 user_id,
		     hotel_id,
		     from_date,
		     to_date,
		     price
		) VALUES($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at
	`

	row := ur.db.QueryRow(
		query,
		booking.RoomId,
		booking.UserId,
		booking.HotelId,
		booking.FromDate,
		booking.ToDate,
		booking.Price,
	)
	err := row.Scan(&booking.ID, &booking.CreatedAt)
	if err != nil {
		return nil, err
	}

	return booking, nil
}

func (ur *bookingRepo) Get(id int64) (*repo.Booking, error) {
	var result repo.Booking

	query := `
		SELECT
			id,
			room_id,
			user_id,
			hotel_id,
			from_date,
			to_date,
			price,
			created_at
		FROM bookings
		WHERE id=$1
	`

	row := ur.db.QueryRow(query, id)
	err := row.Scan(
		&result.ID,
		&result.RoomId,
		&result.UserId,
		&result.HotelId,
		&result.FromDate,
		&result.ToDate,
		&result.Price,
		&result.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (ur *bookingRepo) GetAll(params *repo.GetAllBookingsParams) (*repo.GetAllBookingResult, error) {
	result := repo.GetAllBookingResult{
		Bookings: make([]*repo.Booking, 0),
	}

	offset := (params.Page - 1) * params.Limit

	limit := fmt.Sprintf(" LIMIT %d OFFSET %d ", params.Limit, offset)

	filter := ""
	if params.Search != "" {
		str := "%" + params.Search + "%"
		filter += fmt.Sprintf(`
			WHERE stay ILIKE '%s' `,
			str,
		)
	}

	query := `
		SELECT
			id,
			room_id,
			user_id,
			hotel_id,
			from_date,
			to_date,
			price,
			created_at
		FROM bookings
		` + filter + `
		ORDER BY created_at desc
		` + limit

	rows, err := ur.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var u repo.Booking

		err := rows.Scan(
			&u.ID,
			&u.RoomId,
			&u.UserId,
			&u.HotelId,
			&u.FromDate,
			&u.ToDate,
			&u.Price,
			&u.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		result.Bookings = append(result.Bookings, &u)
	}

	queryCount := `SELECT count(1) FROM bookings ` + filter
	err = ur.db.QueryRow(queryCount).Scan(&result.Count)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (ur *bookingRepo) Update(booking *repo.Booking) (*repo.Booking, error) {
	query := `update bookings set 
			room_id=$1,
			user_id=$2,
			hotel_id=$3,
			from_date=$4,
			to_date=$5,
			price=$6
		where id=$7
		returning created_at
		`

	err := ur.db.QueryRow(
		query,
		booking.RoomId,
		booking.UserId,
		booking.HotelId,
		booking.FromDate,
		booking.ToDate,
		booking.Price,
		booking.ID,
	).Scan(&booking.CreatedAt)
	if err != nil {
		return nil, err
	}

	return booking, nil
}

func (ur *bookingRepo) Delete(id int64) error {
	query := `delete from bookings where id=$1 returning id`

	result, err := ur.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
