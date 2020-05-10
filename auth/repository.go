package auth

import (
	"github.com/lakkinzimusic/horse_maze/models"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUser(username, password string) (*models.User, error)
}
