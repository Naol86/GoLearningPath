package usecase

import (
	"context"
	"time"

	"github.com/naol86/learn-go/CleanArchitecture/domain"
	"github.com/naol86/learn-go/CleanArchitecture/internal/tokenutil"
)

type loginUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

// CreateAccessToken implements domain.LoginUsecase.
func (l *loginUsecase) CreateAccessToken(user *domain.User, secret string, exp int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, exp)
}

// CreateRefreshToken implements domain.LoginUsecase.
func (l *loginUsecase) CreateRefreshToken(user *domain.User, secret string, exp int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, exp)

}

// GetUserByEmail implements domain.LoginUsecase.
func (l *loginUsecase) GetUserByEmail(c context.Context, email string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, l.contextTimeout)
	defer cancel()
	return l.userRepository.GetByEmail(ctx, email)
}

func NewLoginUseCase(ur domain.UserRepository, timeout time.Duration) domain.LoginUsecase {
	return &loginUsecase{
		userRepository: ur,
		contextTimeout: timeout,
	}
}
