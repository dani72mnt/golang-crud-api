package userhandler

import (
	userservice "khademi-practice/service/user"
)

type UserHandler struct {
	userSvc userservice.UserService
}

func New(svc userservice.UserService) UserHandler {
	return UserHandler{
		userSvc: svc,
	}
}
