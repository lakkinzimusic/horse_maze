package main

import (
	"fmt"
	"my_projects/horse_maze/maze"
	"my_projects/horse_maze/searcher"
)

func main() {
	// startPoint := maze.Point{X: 0, Y: 0}
	maze := maze.Maze{Height: 7, Width: 7}
	succsessBranch := searcher.StartSearch(maze)
	fmt.Println(succsessBranch)
}
