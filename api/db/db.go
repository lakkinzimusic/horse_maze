package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//InitDB func
func InitDB(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) *sql.DB {
	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
	db, err := sql.Open(Dbdriver, DBURL)
	// db, err := sql.Open("mysql", "steven:password@tcp(fullstack-mysql:3306)/horse_maze?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	return db
	// defer Database.Close() //разобраться, зачем закрывать db
}
