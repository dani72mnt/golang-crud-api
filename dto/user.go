package dto

type UserCreateReq struct {
	Name     string `json:"name"`
	Family   string `json:"family"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserUpdateReq struct {
	Name   string `json:"name"`
	Family string `json:"family"`
	Email  string `json:"email"`
}

type UserRes struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Family string `json:"family"`
	Email  string `json:"email"`
}

type UserIDURI struct {
	ID int `json:"id" uri:"id"`
}
