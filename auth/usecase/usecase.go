package usecase

import (
	"time"

	"github.com/lakkinzimusic/horse_maze/models"

	"github.com/lakkinzimusic/horse_maze/auth"
)

type AuthUseCase struct {
	userRepo auth.UserRepository
}

//NewAuthUseCase func
func NewAuthUseCase(userRepo auth.UserRepository, hashSalt string, signingKey []byte, tokenTTLSeconds time.Duration) *AuthUseCase {
	return &AuthUseCase{userRepo: userRepo}
}

//SignUp func
func (a *AuthUseCase) SignUp(username, password string) error {
	user := &models.User{
		Username: username,
		Password: password,
	}

	return a.userRepo.CreateUser(user)
}

//SignIn func
func (a *AuthUseCase) SignIn(username, password string) (string, error) {
	user, _ := a.userRepo.GetUser(username, password)

	return user.Username, nil
}
