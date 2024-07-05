package dto

type LoginResponseData struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type SignupResponseData struct {
	Id    uint   `json:"id"`
	Email string `json:"email"`
}

type VerificationCodeData struct {
	VerificationCode int `json:"verification_code,omitempty"`
}
