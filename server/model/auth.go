package model

type SignUpRequest struct {
	FirstName string `json: firstName`
	LastName string  `json: lastName`
	Email string     `json: email`
	Password string  `json: password`
}

type UserResponse struct {
	BaseModel
	Name string `json: name`
	Email string `json: email`
}

type AuthResponse struct {
	Token string `json: token`
	User UserResponse `json: userResponse`
}