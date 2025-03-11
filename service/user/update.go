package userservice

import (
	"context"
	"khademi-practice/dto"
)

func (s UserService) Update(ctx context.Context, param dto.UserUpdateReq, id int) (dto.UserRes, error) {
	user, err := s.repo.Update(ctx, param, id)

	if err != nil {
		return dto.UserRes{}, err
	}

	return dto.UserRes{
		Id:     user.Id,
		Name:   user.Name,
		Family: user.Family,
		Email:  user.Email,
	}, nil
}
