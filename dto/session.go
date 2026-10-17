package dto

type JWTRequest struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}

type JWTResponse struct {
	JWTRequest  `json:",inline"`
	AccessToken string `json:"accessToken"`
}
