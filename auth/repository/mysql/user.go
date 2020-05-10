package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lakkinzimusic/horse_maze/models"
)

//UserRepository struct
type UserRepository struct {
	db *sql.DB
}

//NewUserRepository func
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

//GetUser func
func (r UserRepository) GetUser(username, password string) (*models.User, error) {
	row, err := r.db.Query("SELECT useraname FROM users WHERE username = ?", username)
	if err != nil {
		return nil, err
	}
	user := &models.User{}
	err = row.Scan(&user.Username)
	if err != nil {
		return nil, err
	}
	return user, err
}

//CreateUser func
func (r UserRepository) CreateUser(username, password string) error {
	_, err := r.db.Exec("INSERT INTO users VALUES(?, ?)", username, password)
	if err != nil {
		return err
	}
	return nil
}
