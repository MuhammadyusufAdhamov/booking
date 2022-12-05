package postgres

import (
	"database/sql"
	"fmt"

	"github.com/MuhammadyusufAdhamov/booking/storage/repo"
	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) repo.UserStorageI {
	return &userRepo{
		db: db,
	}
}

func (ur *userRepo) Create(user *repo.User) (*repo.User, error) {
	query := `
		INSERT INTO users(
			first_name,
			last_name,
			email,
			phone_number,
			username,
			password,
		    type
		) VALUES($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, created_at
	`

	row := ur.db.QueryRow(
		query,
		user.FirstName,
		user.LastName,
		user.Email,
		user.PhoneNumber,
		user.Username,
		user.Password,
		user.Type,
	)

	err := row.Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *userRepo) Get(id int64) (*repo.User, error) {
	var result repo.User

	query := `
		SELECT
			id,
			first_name,
			last_name,
			email,
			phone_number,
			username,
			password,
			type,
			created_at
		FROM users
		WHERE id=$1
	`

	row := ur.db.QueryRow(query, id)
	err := row.Scan(
		&result.ID,
		&result.FirstName,
		&result.LastName,
		&result.Email,
		&result.PhoneNumber,
		&result.Username,
		&result.Password,
		&result.Type,
		&result.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (ur *userRepo) GetAll(params *repo.GetAllUsersParams) (*repo.GetAllUsersResult, error) {
	result := repo.GetAllUsersResult{
		Users: make([]*repo.User, 0),
	}

	offset := (params.Page - 1) * params.Limit

	limit := fmt.Sprintf(" LIMIT %d OFFSET %d ", params.Limit, offset)

	filter := ""
	if params.Search != "" {
		str := "%" + params.Search + "%"
		filter += fmt.Sprintf(`
			WHERE first_name ILIKE '%s' OR last_name ILIKE '%s' OR email ILIKE '%s' 
				OR username ILIKE '%s' OR phone_number ILIKE '%s'`,
			str, str, str, str, str,
		)
	}

	query := `
		SELECT
			id,
			first_name,
			last_name,
			email,
			phone_number,
			username,
			password,
			type,
			created_at
		FROM users
		` + filter + `
		ORDER BY created_at desc
		` + limit

	rows, err := ur.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var u repo.User

		err := rows.Scan(
			&u.ID,
			&u.FirstName,
			&u.LastName,
			&u.Email,
			&u.PhoneNumber,
			&u.Username,
			&u.Password,
			&u.Type,
			&u.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		result.Users = append(result.Users, &u)
	}

	queryCount := `SELECT count(1) FROM users ` + filter
	err = ur.db.QueryRow(queryCount).Scan(&result.Count)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (ur *userRepo) Update(user *repo.User) (*repo.User, error) {
	query := `update users set 
			first_name=$1,
			last_name=$2,
			email=$3,
			phone_number=$4,
			username=$5,
			password=$6,
			type=$7
		where id=$8
		returning created_at
		`

	err := ur.db.QueryRow(
		query,
		user.FirstName,
		user.LastName,
		user.Email,
		user.PhoneNumber,
		user.Username,
		user.Password,
		user.Type,
		user.ID,
	).Scan(&user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *userRepo) Delete(id int64) error {
	query := `delete from users where id=$1`

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

func (ur *userRepo) GetByEmail(email string) (*repo.User, error) {
	var result repo.User

	query := `
		SELECT
			id,
			first_name,
			last_name,
			email,
			phone_number,
			username,
			password,
			type,
			created_at
		FROM users
		WHERE email=$1
	`

	row := ur.db.QueryRow(query, email)
	err := row.Scan(
		&result.ID,
		&result.FirstName,
		&result.LastName,
		&result.Email,
		&result.PhoneNumber,
		&result.Username,
		&result.Password,
		&result.Type,
		&result.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (ur *userRepo) UpdatePassword(req *repo.UpdatePassword) error {
	query := `UPDATE users SET password=$1 WHERE id=$2`

	_, err := ur.db.Exec(query, req.Password, req.UserID)
	if err != nil {
		return err
	}

	return nil
}
