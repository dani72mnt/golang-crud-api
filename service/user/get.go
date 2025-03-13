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

	// use pre-allocation and make func
	// userResponses := make([]dto.UserRes, 0, len(users))
	// for _, user := range users { // dont make copies with it is not reasnoable
	// 	userResponses = append(userResponses, dto.UserRes{
	// 		Id:     user.Id,
	// 		Name:   user.Name,
	// 		Family: user.Family,
	// 		Email:  user.Email,
	// 	})
	// }

	userResponses := make([]dto.UserRes, 0, len(users))
	for i := range users {
		userResponses = append(userResponses, dto.UserRes{
			Id:     users[i].Id,
			Name:   users[i].Name,
			Family: users[i].Family,
			Email:  users[i].Email,
		})
	}

	return userResponses, nil
}
