package usecase

import (
	"github.com/lakkinzimusic/horse_maze/api/auth"
)

type AuthUseCase struct {
	userRepo auth.UserRepository
}

//NewAuthUseCase func
func NewAuthUseCase(userRepo auth.UserRepository) *AuthUseCase {
	return &AuthUseCase{userRepo: userRepo}
}

//SignUp func
func (a *AuthUseCase) SignUp(username, password string) error {

	return a.userRepo.CreateUser(username, password)
}

//SignIn func
func (a *AuthUseCase) SignIn(username, password string) (string, error) {
	user, _ := a.userRepo.GetUser(username, password)

	return user.Username, nil
}
