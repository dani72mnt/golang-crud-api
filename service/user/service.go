package userservice

import (
	"context"
	"khademi-practice/config"
	"khademi-practice/dto"
	"khademi-practice/entity"
	userrepository "khademi-practice/repository"
)

type userRepository interface {
	GetAll(ctx context.Context) ([]entity.User, error)
	Get(ctx context.Context, id int) (entity.User, error)
	Create(ctx context.Context, params dto.UserCreateReq) error
	Update(ctx context.Context, params dto.UserUpdateReq, id int) (entity.User, error)
	Delete(ctx context.Context, id int) error
}

type UserService struct {
	cfg  *config.Config
	repo userRepository
}

func New(cfg *config.Config, repo userrepository.UserRepository) UserService {
	return UserService{
		cfg:  cfg,
		repo: repo,
	}
}
