package main

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

type Cell struct {
	Order int
	X     int
	Y     int
}

//GetBranch func
func GetBranch(branch_id int, db *sql.DB) (cells []Cell) {
	rows, err := db.Query("SELECT t.order, t.x_coordinate, t.y_coordinate FROM winner_branches wb JOIN turns t ON wb.id = t.branch_id WHERE wb.id = (?) ORDER BY t.y_coordinate, t.x_coordinate ", branch_id)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		c := Cell{}
		err := rows.Scan(&c.Order, &c.X, &c.Y)
		if err != nil {
			panic(err)
		}
		cells = append(cells, c)
	}
	return cells
}
