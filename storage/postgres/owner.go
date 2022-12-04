package postgres

import (
	"database/sql"
	"fmt"
	"github.com/MuhammadyusufAdhamov/booking/storage/repo"
	"github.com/jmoiron/sqlx"
)

type ownerRepo struct {
	db *sqlx.DB
}

func NewOwner(db *sqlx.DB) repo.OwnerStorageI {
	return &ownerRepo{
		db: db,
	}
}

func (ur *ownerRepo) Create(owner *repo.Owner) (*repo.Owner, error) {
	query := `
		INSERT INTO owners(
			first_name,
			last_name,
			email,
			phone_number,
			username,
			password
		) VALUES($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at
	`

	row := ur.db.QueryRow(
		query,
		owner.FirstName,
		owner.LastName,
		owner.Email,
		owner.PhoneNumber,
		owner.Username,
		owner.Password,
	)

	err := row.Scan(&owner.ID, &owner.CreatedAt)
	if err != nil {
		return nil, err
	}

	return owner, nil
}

func (ur *ownerRepo) Get(id int64) (*repo.Owner, error) {
	var result repo.Owner

	query := `
		SELECT
			id,
			first_name,
			last_name,
			email,
			phone_number,
			username,
			password,
			created_at
		FROM owners
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
		&result.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (ur *ownerRepo) GetAll(params *repo.GetAllOwnersParams) (*repo.GetAllOwnersResult, error) {
	result := repo.GetAllOwnersResult{
		Owners: make([]*repo.Owner, 0),
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
			created_at
		FROM owners
		` + filter + `
		ORDER BY created_at desc
		` + limit

	rows, err := ur.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var u repo.Owner

		err := rows.Scan(
			&u.ID,
			&u.FirstName,
			&u.LastName,
			&u.Email,
			&u.PhoneNumber,
			&u.Username,
			&u.Password,
			&u.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		result.Owners = append(result.Owners, &u)
	}

	queryCount := `SELECT count(1) FROM owners ` + filter
	err = ur.db.QueryRow(queryCount).Scan(&result.Count)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (ur *ownerRepo) Update(user *repo.Owner) (*repo.Owner, error) {
	query := `update owners set 
			first_name=$1,
			last_name=$2,
			email=$3,
			phone_number=$4,
			username=$5,
			password=$6
		where id=$7
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
		user.ID,
	).Scan(&user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *ownerRepo) Delete(id int64) error {
	query := `delete from owners where id=$1 returning id`

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
