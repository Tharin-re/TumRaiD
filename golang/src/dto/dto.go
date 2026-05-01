package dto

type RegisterUserPassReqBody struct {
	Username string
	Password string
}

type LoginData struct {
	Username string
	Password string
	Token string
}