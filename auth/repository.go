package auth

import (
	"github.com/lakkinzimusic/horse_maze/models"
)

//UserRepository interface
type UserRepository interface {
	CreateUser(username, password string) error
	GetUser(username, password string) (*models.User, error)
}
