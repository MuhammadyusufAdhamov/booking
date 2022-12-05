package postgres

import (
	"database/sql"
	"fmt"

	"github.com/MuhammadyusufAdhamov/booking/storage/repo"
	"github.com/jmoiron/sqlx"
)

type roomRepo struct {
	db *sqlx.DB
}

func NewRoom(db *sqlx.DB) repo.RoomsStorageI {
	return &roomRepo{
		db: db,
	}
}

func (ur *roomRepo) Create(room *repo.Room) (*repo.Room, error) {
	query := `
		INSERT INTO rooms(
			type,
			number_of_room,
			room_image_url,
			status,
		    hotel_id
		) VALUES($1, $2, $3, $4, $5)
		RETURNING id, created_at
	`

	row := ur.db.QueryRow(
		query,
		room.Type,
		room.NumberOfRoom,
		room.RoomImageUrl,
		room.Status,
		room.HotelId,
	)

	err := row.Scan(&room.ID, &room.CreatedAt)
	if err != nil {
		return nil, err
	}

	return room, nil
}

func (ur *roomRepo) Get(id int64) (*repo.Room, error) {
	var result repo.Room

	query := `
		SELECT
			id,
			type,
			number_of_room,
			room_image_url,
			status,
			hotel_id,
			created_at
		FROM rooms
		WHERE id=$1
	`

	row := ur.db.QueryRow(query, id)
	err := row.Scan(
		&result.ID,
		&result.Type,
		&result.NumberOfRoom,
		&result.RoomImageUrl,
		&result.Status,
		&result.HotelId,
		&result.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (ur *roomRepo) GetAll(params *repo.GetAllRoomsParams) (*repo.GetAllRoomsResult, error) {
	result := repo.GetAllRoomsResult{
		Rooms: make([]*repo.Room, 0),
	}

	offset := (params.Page - 1) * params.Limit

	limit := fmt.Sprintf(" LIMIT %d OFFSET %d ", params.Limit, offset)

	filter := ""
	if params.Search != "" {
		str := "%" + params.Search + "%"
		filter += fmt.Sprintf(`
			WHERE type ILIKE '%s' OR sleeps ILIKE '%s' OR status ILIKE '%s' `,
			str, str, str,
		)
	}

	query := `
		SELECT
			id,
			type,
			number_of_room,
			room_image_url,
			status,
			hotel_id,
			created_at
		FROM rooms
		` + filter + `
		ORDER BY created_at desc
		` + limit

	rows, err := ur.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var u repo.Room

		err := rows.Scan(
			&u.ID,
			&u.Type,
			&u.NumberOfRoom,
			&u.RoomImageUrl,
			&u.Status,
			&u.HotelId,
			&u.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		result.Rooms = append(result.Rooms, &u)
	}

	queryCount := `SELECT count(1) FROM rooms ` + filter
	err = ur.db.QueryRow(queryCount).Scan(&result.Count)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (ur *roomRepo) Update(room *repo.Room) (*repo.Room, error) {
	query := `update rooms set 
			type=$1,
			number_of_room=$2,
			room_image_url=$3,
			status=$4,
			hotel_id=$5
		where id=$6
		returning created_at
		`

	err := ur.db.QueryRow(
		query,
		room.Type,
		room.NumberOfRoom,
		room.RoomImageUrl,
		room.Status,
		room.HotelId,
		room.ID,
	).Scan(&room.CreatedAt)
	if err != nil {
		return nil, err
	}

	return room, nil
}

func (ur *roomRepo) Delete(id int64) error {
	query := `delete from rooms where id=$1 returning id`

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
