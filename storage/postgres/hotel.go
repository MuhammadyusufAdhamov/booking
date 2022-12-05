package postgres

import (
	"database/sql"
	"fmt"

	"github.com/MuhammadyusufAdhamov/booking/storage/repo"
	"github.com/jmoiron/sqlx"
)

type hotelRepo struct {
	db *sqlx.DB
}

func NewHotel(db *sqlx.DB) repo.HotelStorageI {
	return &hotelRepo{
		db: db,
	}
}

func (ur *hotelRepo) Create(hotel *repo.Hotel) (*repo.Hotel, error) {
	query := `
		INSERT INTO hotels(
			user_id,
			hotel_name,
			hotel_location,
			hotel_image_url,
			number_of_rooms
		) VALUES($1, $2, $3, $4, $5)
		RETURNING id, created_at
	`

	row := ur.db.QueryRow(
		query,
		hotel.UserID,
		hotel.HotelName,
		hotel.HotelLocation,
		hotel.HotelImageUrl,
		hotel.NumberOfRooms,
	)

	err := row.Scan(&hotel.ID, &hotel.CreatedAt)
	if err != nil {
		return nil, err
	}

	return hotel, nil
}

func (ur *hotelRepo) Get(id int64) (*repo.Hotel, error) {
	var result repo.Hotel

	query := `
		SELECT
			id,
			user_id,
			hotel_name,
			hotel_location,
			hotel_image_url,
			number_of_rooms,
			created_at
		FROM hotels
		WHERE id=$1
	`

	row := ur.db.QueryRow(query, id)
	err := row.Scan(
		&result.ID,
		&result.UserID,
		&result.HotelName,
		&result.HotelLocation,
		&result.HotelImageUrl,
		&result.NumberOfRooms,
		&result.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (ur *hotelRepo) GetAll(params *repo.GetAllHotelsParams) (*repo.GetAllHotelsResult, error) {
	result := repo.GetAllHotelsResult{
		Hotels: make([]*repo.Hotel, 0),
	}

	offset := (params.Page - 1) * params.Limit

	limit := fmt.Sprintf(" LIMIT %d OFFSET %d ", params.Limit, offset)

	filter := ""
	if params.Search != "" {
		str := "%" + params.Search + "%"
		filter += fmt.Sprintf(`
			WHERE hotel_name ILIKE '%s' OR hotel_rating ILIKE '%s' OR hotel_location ILIKE '%s' `,
			str, str, str,
		)
	}

	query := `
		SELECT
			id,
			user_id,
			hotel_name,
			hotel_location,
			hotel_image_url,
			number_of_rooms,
			created_at
		FROM hotels
		` + filter + `
		ORDER BY created_at desc
		` + limit

	rows, err := ur.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var h repo.Hotel

		err := rows.Scan(
			&h.ID,
			&h.UserID,
			&h.HotelName,
			&h.HotelLocation,
			&h.HotelImageUrl,
			&h.NumberOfRooms,
			&h.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		result.Hotels = append(result.Hotels, &h)
	}

	queryCount := `SELECT count(1) FROM hotels ` + filter
	err = ur.db.QueryRow(queryCount).Scan(&result.Count)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (ur *hotelRepo) Update(hotel *repo.Hotel) (*repo.Hotel, error) {
	query := `update hotels set 
			user_id=$1,
			hotel_name=$2,
			hotel_location=$3,
			hotel_image_url=$4,
			number_of_rooms=$5
		where id=$6
		returning created_at
		`

	err := ur.db.QueryRow(
		query,
		hotel.UserID,
		hotel.HotelName,
		hotel.HotelLocation,
		hotel.HotelImageUrl,
		hotel.NumberOfRooms,
		hotel.ID,
	).Scan(&hotel.CreatedAt)
	if err != nil {
		return nil, err
	}

	return hotel, nil
}

func (ur *hotelRepo) Delete(id int64) error {
	query := `delete from hotels where id=$1 returning id`

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
