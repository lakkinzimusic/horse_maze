package drawer

import (
	game "my_projects/horse_maze/game"
	maze "my_projects/horse_maze/game"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

//Drawer struct
type Drawer struct {
}

func drawHeader(width, cellWidth int) {
	result := ""
	for i := 0; i <= width; i++ {
		result = result + strings.Repeat(" ", cellWidth+1) + stringValueOf(i)
	}
	WriteFile(" " + result + "\n")
}

//PaintMaze func
func (d *Drawer) PaintMaze(maze game.Maze, branch game.Branch) {
	WriteFile("\n\n")
	// indentWidth := utf8.RuneCountInString(strconv.Itoa(maze.Height + 1))
	cellWidth := utf8.RuneCountInString(strconv.Itoa(branch.Moves[len(branch.Moves)-1].Order)) + 1
	drawHeader(maze.Width, cellWidth)
	drawPartition(maze.Width, cellWidth)
	countRow := 0
	lineNumberString := 2
	countRows := 0
	for i := 0; i <= maze.Width+1; i++ {
		// thisIndentWidth := utf8.RuneCountInString(strconv.Itoa(lineNumberString))
		if i == 0 {
			// WriteFile(strconv.Itoa(lineNumberString) + strings.Repeat(" ", indentWidth-thisIndentWidth))
			WriteFile(strconv.Itoa(lineNumberString-1) + " ")
		}
		if countRow == maze.Width+1 && i != 0 {
			WriteFile("||\n")
			drawPartition(maze.Width, cellWidth)
			if countRows < maze.Height {
				i = 0
			} else {
				break
			}
			// WriteFile(strconv.Itoa(lineNumberString) + strings.Repeat(" ", indentWidth-thisIndentWidth))
			WriteFile(strconv.Itoa(lineNumberString) + " ")
			lineNumberString++
			countRow = 0
			countRows++

		}
		for _, move := range branch.Moves {
			if move.Point.X == i && move.Point.Y == lineNumberString {
				drawCell(move, cellWidth)
				countRow++
				continue
			}
		}

		// if isDead(maze.DeadPoints, cell.Point.X, cell.Point.Y) {
		// 	WriteFile(`||` + strings.Repeat(" ", cellWidth-1) + "*")
		// 	countRow++
		// 	continue
		// }

		WriteFile(`||` + strings.Repeat(" ", cellWidth-1) + " ")
		countRow++
	}

	// WriteFile("||\n")
	// drawPartition(maze.Width, cellWidth)

}

func stringValueOf(i int) string {
	var foo = "ABCDEFGHIJKLMNOPQRSTUVWXYZ???????????????????"
	return string(foo[i])
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

func drawCell(cell game.Move, cellWidth int) {
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
	f, err := os.OpenFile("./home.html", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err = f.WriteString(char); err != nil {
		panic(err)
	}
}
