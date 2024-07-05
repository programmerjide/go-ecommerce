package dto

type UserLoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserSignupRequestDto struct {
	UserLoginDto
	Phone string `json:"phone"`
}

type VerificationCodeInput struct {
	Code int `json:"code"`
}
