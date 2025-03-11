package userservice

import (
	"context"
)

func (s UserService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
