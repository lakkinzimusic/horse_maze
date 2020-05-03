package handlers

import (
	"html/template"
	draw "my_projects/horse_maze/drawer"
	game "my_projects/horse_maze/game"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

var branch = game.Branch{}
var count = 0
var maze = game.Maze{7, 7}

//Home func
func Home(w http.ResponseWriter, r *http.Request) {
	clearFile()
	vars := mux.Vars(r)
	letter, _ := strconv.Atoi(vars["letter"])
	number, _ := strconv.Atoi(vars["number"])
	letter = letter - 1
	number = number - 1
	if !branch.MoveExist(letter, number) {
		move := game.Move{Order: count, Point: game.Point{X: letter, Y: number}}
		branch.Moves = append(branch.Moves, move)

		count = count + 1

	}
	drawer := draw.Drawer{}
	drawer.PaintMaze(maze, branch)
	tmpl := template.Must(template.ParseFiles("./home.html"))
	tmpl.Execute(w, nil)
}

func clearFile() {
	configFile, _ := os.OpenFile("./home.html", os.O_RDWR, 0666)
	defer configFile.Close()
	//some actions happen here
	configFile.Truncate(0)

}
