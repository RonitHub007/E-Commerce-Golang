package dto

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserSignuP struct {
	UserLogin
	Phone string `json:"phone"`
}
