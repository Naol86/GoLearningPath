package usecase

import (
	"context"
	"time"

	"github.com/naol86/learn-go/CleanArchitecture/domain"
	"github.com/naol86/learn-go/CleanArchitecture/internal/tokenutil"
)

type SignupUseCase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

// Create implements domain.SignupUseCase.
func (s *SignupUseCase) Create(c context.Context, user *domain.User) error {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout)
	defer cancel()
	return s.userRepository.Create(ctx, user)
}

// CreateAccessToken implements domain.SignupUseCase.
func (s *SignupUseCase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

// CreateRefreshToken implements domain.SignupUseCase.
func (s *SignupUseCase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}

// GetUserByEmail implements domain.SignupUseCase.
func (s *SignupUseCase) GetUserByEmail(c context.Context, email string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, s.contextTimeout);
	defer cancel()
	return s.userRepository.GetByEmail(ctx, email)
}

func NewSignupUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.SignupUseCase {
	return &SignupUseCase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}
