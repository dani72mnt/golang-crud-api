package userrepository

import (
	"context"
	"khademi-practice/config/models"
	"khademi-practice/dto"
	"khademi-practice/entity"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

const timeout = 3

type UserRepository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (r UserRepository) GetAllOrm(ctx context.Context) ([]entity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	users, err := models.Users().All(ctx, r.db)
	if err != nil {
		return nil, err
	}

	var result []entity.User
	for _, u := range users {
		result = append(result, entity.User{
			Id:     u.ID,
			Name:   u.Name,
			Family: u.Family,
			Email:  u.Email,
		})
	}

	return result, nil
}

func (r UserRepository) GetOrm(ctx context.Context, id int) (entity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout*time.Second)
	defer cancel()

	userModel, err := models.Users(qm.Where("id=?", id)).One(ctx, r.db)
	if err != nil {
		return entity.User{}, err
	}

	user := entity.User{
		Id:     userModel.ID,
		Name:   userModel.Name,
		Family: userModel.Family,
		Email:  userModel.Email,
	}

	return user, nil
}

func (r UserRepository) GetAll(ctx context.Context) ([]entity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `SELECT id, name, family, email FROM users` // offset
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	offset := 10 // per_page=10
	users := make([]entity.User, 0, offset)

	for rows.Next() {
		var user entity.User
		if err := rows.Scan(&user.Id, &user.Name, &user.Family, &user.Email); err != nil {
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
	ctx, cancel := context.WithTimeout(ctx, timeout*time.Second)
	defer cancel()

	query := `SELECT id, name, family, email FROM users WHERE id = $1`
	var user entity.User

	if err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.Id,
		&user.Name,
		&user.Family,
		&user.Email,
	); err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (r UserRepository) Create(ctx context.Context, params dto.UserCreateReq) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `INSERT INTO users (name, family, email, password) VALUES ($1, $2, $3, $4) RETURNING id`

	var lastInserId int

	if err := r.db.QueryRowContext(ctx, query, params.Name, params.Family, params.Email, params.Password).
		Scan(&lastInserId); err != nil {
		return 0, err
	}

	return lastInserId, nil
}

func (r UserRepository) Update(ctx context.Context, params dto.UserUpdateReq, id int) (entity.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	query := `UPDATE users SET name = $1, family = $2, email = $3 WHERE id = $5`

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
