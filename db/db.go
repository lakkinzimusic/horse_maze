package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//InitDB func
func InitDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/horse_maze")
	if err != nil {
		fmt.Println(err)
	}
	return db
	// defer Database.Close() //разобраться, зачем закрывать db
}