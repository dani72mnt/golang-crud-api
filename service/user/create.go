package userservice

import (
	"context"
	"khademi-practice/dto"
	hashutil "khademi-practice/pkg"
)

func (s UserService) Create(ctx context.Context, param dto.UserCreateReq) error {
	hashedPassword, err := hashutil.HashPassword(param.Password)
	if err != nil {
		return err
	}

	param.Password = hashedPassword

	return s.repo.Create(ctx, param)
}
