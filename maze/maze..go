package maze

type Maze struct {
	Height    int
	Width     int
	DeadCells []Point
}

//Point struct
type Point struct {
	X int
	Y int
}
