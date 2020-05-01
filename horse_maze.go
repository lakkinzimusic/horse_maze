package main

import (
	"database/sql"
	"my_projects/horse_maze/maze"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
	// "rsc.io/quote"
)

func main() {
	db := InitDB()
	// quote.Hello()
	deads := []maze.Point{{2, 4}, {4, 5}}
	maze := maze.Maze{Height: 10, Width: 10, DeadCells: deads}
	// searcher.StartSearch(maze, db)
	Draw(db, maze.Width, maze.Height, maze.DeadCells)
}

//Draw func
func Draw(db *sql.DB, mazeWidth, mazeHeight int, deadCells []maze.Point) {
	WriteFile("\n\n")
	cells := GetBranch(1, db)
	cellWidth := utf8.RuneCountInString(strconv.Itoa(cells[len(cells)-1].Order)) + 1 //длина последнего элемента

	indentWidth := utf8.RuneCountInString(strconv.Itoa(mazeHeight + 1))
	drawHeader(mazeWidth, cellWidth)
	drawPartition(mazeWidth, cellWidth)
	countRow := 0
	lineNumberString := 1
	for i, cell := range cells {
		thisIndentWidth := utf8.RuneCountInString(strconv.Itoa(lineNumberString))
		if i == 0 {
			WriteFile(strconv.Itoa(lineNumberString) + strings.Repeat(" ", indentWidth-thisIndentWidth))
		}
		if countRow == mazeWidth+1 && i != 0 {
			WriteFile("||\n")
			drawPartition(mazeWidth, cellWidth)
			WriteFile(strconv.Itoa(lineNumberString) + strings.Repeat(" ", indentWidth-thisIndentWidth))
			lineNumberString++
			countRow = 0
		}
		drawCell(cell, cellWidth, deadCells)
		if nextDead(deadCells, cell.X, cell.Y) {
			WriteFile(`||` + strings.Repeat(" ", cellWidth-1) + "*")
			countRow++
		}
		countRow++
	}
	WriteFile("||\n")
	drawPartition(mazeWidth, cellWidth)
}

func stringValueOf(i int) string {
	var foo = "ABCDEFGHIJKLMNOPQRSTUVWXYZ???????????????????"
	return string(foo[i])
}

func drawHeader(width, cellWidth int) {
	result := ""
	// пишем 1ю строку
	for i := 0; i <= width; i++ {
		result = result + strings.Repeat(" ", cellWidth+1) + stringValueOf(i)
	}
	WriteFile(" " + result + "\n")
}

func getLine(element string, lenth int) string {
	for i := 0; i < lenth; i++ {
		element = element + element
	}
	return element
}

func drawPartition(width, cellWidth int) {
	result := ""

	for j := 0; j <= width; j++ {
		result = result + `||` + strings.Repeat("=", cellWidth)
	}
	result = "  " + result + `||` + "\n"
	WriteFile(result)
}

func drawCell(cell Cell, cellWidth int, deadCells []maze.Point) {
	WriteFile(`||` + strings.Repeat(" ", cellWidth-utf8.RuneCountInString(strconv.Itoa(cell.Order))) + strconv.Itoa(cell.Order))
}

// TODO: ПРОВЕРИТЬ НА ПЕРЕНОС СТРОКИ
func nextDead(deadCells []maze.Point, X, Y int) bool {
	for _, dead := range deadCells {

		if dead.X == X+1 && dead.Y == Y {
			return true
		}

	}
	return false
}

//WriteFile func
func WriteFile(char string) {
	f, err := os.OpenFile("./map.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err = f.WriteString(char); err != nil {
		panic(err)
	}
}
