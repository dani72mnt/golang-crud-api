package userrepository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"khademi-practice/dto"
	"khademi-practice/entity"
	"time"
)

type UserRepository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (r UserRepository) GetAll(ctx context.Context) ([]entity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `SELECT id, name, family, email FROM users`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entity.User
	for rows.Next() {
		var user entity.User
		err := rows.Scan(&user.Id, &user.Name, &user.Family, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r UserRepository) Get(ctx context.Context, id int) (entity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `SELECT id, name, family, email FROM users WHERE id = $1`
	var user entity.User

	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.Id,
		&user.Name,
		&user.Family,
		&user.Email,
	)

	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (r UserRepository) Create(ctx context.Context, params dto.UserCreateReq) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `INSERT INTO users (name, family, email, password) VALUES ($1, $2, $3, $4)`

	_, err := r.db.ExecContext(ctx, query, params.Name, params.Family, params.Email, params.Password)
	if err != nil {
		return err
	}

	return nil
}

func (r UserRepository) Update(ctx context.Context, params dto.UserUpdateReq, id int) (entity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `UPDATE users SET name = $1, family = $2, email = $3 WHERE id = $4`

	_, err := r.db.ExecContext(ctx, query, params.Name, params.Family, params.Email, id)
	if err != nil {
		return entity.User{}, err
	}

	var updatedUser entity.User
	selectQuery := `SELECT id, name, family, email FROM users WHERE id = $1`

	err = r.db.QueryRowContext(ctx, selectQuery, id).Scan(
		&updatedUser.Id,
		&updatedUser.Name,
		&updatedUser.Family,
		&updatedUser.Email,
	)

	if err != nil {
		return entity.User{}, err
	}

	return updatedUser, nil
}

func (r UserRepository) Delete(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `DELETE FROM users WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, id)

	if err != nil {
		return err
	}

	return nil
}
