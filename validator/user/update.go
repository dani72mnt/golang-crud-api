package uservalidator

import (
	"khademi-practice/dto"

	validator "github.com/rezakhademix/govalidator/v2"
)

func (v UserValidator) ValidateUpdateReq(vd validator.Validator, req dto.UserUpdateReq) (bool, map[string]string) {

	vd.RequiredString(req.Name, "name", "name is required").
		MinString(req.Name, 2, "name", "name must be at least 2 characters long").
		MaxString(req.Name, 50, "name", "name must be at most 50 characters long")

	vd.RequiredString(req.Family, "family", "family is required").
		MinString(req.Family, 2, "family", "family must be at least 2 characters long").
		MaxString(req.Family, 50, "family", "family must be at most 50 characters long")

	vd.RequiredString(req.Email, "email", "email is required").
		Email(req.Email, "email", "invalid email format")

	return vd.IsPassed(), vd.Errors()
}
