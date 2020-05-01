package searcher

import (
	"database/sql"
	"fmt"
	"my_projects/horse_maze/maze"
	"os"
)

// ThisMaze var
var ThisMaze = maze.Maze{}

// WinnerBranchLen var
var WinnerBranchLen int

// MaxLenBranch int
var MaxLenBranch = 0

// Branch type
type Branch []maze.Point

// TestBranches type
type TestBranches []Branch

// WrongBranches type
type WrongBranches []Branch

var testBranches = TestBranches{}
var startBranch = Branch{}

func isDead(deadCells []maze.Point, point maze.Point) bool {
	for _, dead := range deadCells {
		if dead.X == point.X && dead.Y == point.Y {
			return true
		}
	}
	return false
}

//StartSearch func
func StartSearch(thisMaze maze.Maze, db *sql.DB) {
	ThisMaze = thisMaze
	WinnerBranchLen = (ThisMaze.Height+1)*(ThisMaze.Width+1) - len(ThisMaze.DeadCells)
	for x := 0; x <= ThisMaze.Width; x++ {
		for y := 0; y <= ThisMaze.Height; y++ {
			startPoint := maze.Point{X: x, Y: y}
			if isDead(ThisMaze.DeadCells, startPoint) {
				continue
			}
			fmt.Printf("startPoint %+v\n", startPoint)
			startBranch = append(startBranch, startPoint)
			testBranches = append(testBranches, startBranch) //добавляем одну тестовую ветку (1й ход)
			//цикл, пока ветвление ветки себя не исчерпает
			for len(testBranches) != 0 {
				// _ = resources.GetAvailableMemory()
				lastBranch := testBranches[len(testBranches)-1] //вырезаем последнюю ветку из массива веток

				testBranches = testBranches[:len(testBranches)-1]
				// fmt.Print("testBranches %+v\n", testBranches)
				availableTurns := SearchAvailableTurns(lastBranch) //ищем доступные ходы для ветки

				availableTurns = fullVarnsdorfFilter(lastBranch, availableTurns) //реализуем правило Варнсдорфа
				// fmt.Println(availableTurns)
				availableBranches := initiateAvailableBranches(availableTurns, lastBranch, db) //создаём доступные
				testBranches = append(testBranches, availableBranches...)
			}

			startBranch = startBranch[:0]
			testBranches = testBranches[:0]
		}
	}
}

var minVariants = 8
var variantsCount = 0

func fullVarnsdorfFilter(branch Branch, turns []maze.Point) (setTurns []maze.Point) {

	for _, turn := range turns {
		newBranch := append(branch, turn)
		variantsCount = len(SearchAvailableTurns(newBranch))
		if variantsCount < minVariants {
			minVariants = variantsCount
		}
	}
	for _, turn := range turns {
		newBranch2 := append(branch, turn)
		if len(SearchAvailableTurns(newBranch2)) == minVariants {
			setTurns = append(setTurns, turn)
		}
	}
	minVariants = 8
	return setTurns
}

func initiateAvailableBranches(turns []maze.Point, branch Branch, db *sql.DB) (branches []Branch) {
	for _, turn := range turns {
		newBranch := append(branch, turn)
		if len(newBranch) == WinnerBranchLen {
			// fmt.Println("Winner: ", newBranch)
			WriteFile(newBranch, db)
		}
		// if len(newBranch) > MaxLenBranch {
		// 	MaxLenBranch = len(newBranch)
		// 	// fmt.Println("New record: ", MaxLenBranch)
		// }
		branches = append(branches, newBranch)
	}
	return branches
}

//SearchAvailableTurns func
func SearchAvailableTurns(branch Branch) (setTurns []maze.Point) {
	allTurns := []maze.Point{{X: 1, Y: 2}, {X: 2, Y: 1}, {X: 2, Y: -1}, {X: 1, Y: -2}, {X: -1, Y: -2}, {X: -2, Y: -1}, {X: -2, Y: 1}, {X: -1, Y: 2}}
	for _, turn := range allTurns {
		successTurn, success := tryMove(branch, turn)
		if success {
			setTurns = append(setTurns, successTurn)
		}
	}
	return setTurns
}

func tryMove(branch Branch, turn maze.Point) (maze.Point, bool) {
	lastTurn := branch[len(branch)-1]

	x := lastTurn.X + turn.X
	y := lastTurn.Y + turn.Y

	tryingPoint := maze.Point{X: x, Y: y}

	if checkCorrectingTurn(tryingPoint) {
		if checkExistTurn(branch, tryingPoint) {
			return tryingPoint, true
		}
	}
	return tryingPoint, false
}

func checkExistTurn(branch Branch, tryingPoint maze.Point) bool {
	for _, turn := range branch {
		if turn.X == tryingPoint.X && turn.Y == tryingPoint.Y {
			return false
		}
	}
	return true
}

func checkCorrectingTurn(tryingPoint maze.Point) bool {
	XCorrect := (tryingPoint.X <= ThisMaze.Width) && (tryingPoint.X >= 0)
	YCorrect := (tryingPoint.Y <= ThisMaze.Height) && (tryingPoint.Y >= 0)
	if XCorrect && YCorrect && !isDead(ThisMaze.DeadCells, tryingPoint) {
		return true
	}
	return false
}

//WriteFile func
func WriteFile(branch Branch, db *sql.DB) {
	f, err := os.OpenFile("./log.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	str2 := fmt.Sprintf("%d\n", len(branch))
	str := fmt.Sprintf("%+v\n", branch)

	if _, err = f.WriteString(str2); err != nil {
		panic(err)
	}
	if _, err = f.WriteString(str); err != nil {
		panic(err)
	}
	result, err := db.Exec("INSERT INTO winner_branches (lenth) VALUES (?)", len(branch))
	if err != nil {
		panic(err)
	}
	for i, turn := range branch {
		branchID, _ := result.LastInsertId()
		_, err := db.Exec("INSERT INTO horse_maze.turns (branch_id, `order`, x_coordinate, y_coordinate) VALUES (?, ?, ?, ?)", branchID, i, turn.X, turn.Y)
		if err != nil {
			panic(err)
		}
	}
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
