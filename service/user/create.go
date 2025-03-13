package userservice

import (
	"context"
	"khademi-practice/dto"
	hashutil "khademi-practice/pkg"
)

func (s UserService) Create(ctx context.Context, param dto.UserCreateReq) error {
	hashedPassword, err := hashutil.HashPassword(param.Password)
	if err != nil {
		return err // Do Not return any system or database error and convert it to a clear error without sensitive data
	}

	param.Password = hashedPassword

	_, err = s.repo.Create(ctx, param) // Do Not return any system or database error and convert it to a clear error without sensitive data

	return err
}
