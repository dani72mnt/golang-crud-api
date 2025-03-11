package userservice

import (
	"context"
	"khademi-practice/dto"
)

func (s UserService) Get(ctx context.Context, id int) (dto.UserRes, error) {
	user, err := s.repo.GetOrm(ctx, id)
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

func (s UserService) GetAll(ctx context.Context) ([]dto.UserRes, error) {
	users, err := s.repo.GetAllOrm(ctx)
	if err != nil {
		return nil, err
	}

	var userResponses []dto.UserRes
	for _, user := range users {
		userResponses = append(userResponses, dto.UserRes{
			Id:     user.Id,
			Name:   user.Name,
			Family: user.Family,
			Email:  user.Email,
		})
	}

	return userResponses, nil
}
