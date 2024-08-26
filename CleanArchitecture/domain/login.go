package domain

import "context"

type LoginRequest struct {
	Email string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	User UserResponse `json:"user"`
}

type LoginUsecase interface {
	GetUserByEmail(c context.Context ,email string) (User, error)
	CreateAccessToken(user *User, secret string, exp int) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, exp int) (refreshToken string, err error)
}