package maze

type Maze struct {
	Height int
	Width  int
	// DeadPoints []Point
}

//Point struct
type Point struct {
	X int
	Y int
}

type Branch struct {
	Moves []Move
}

func (b *Branch) MoveExist(x, y int) bool {
	for _, move := range b.Moves {
		if move.Point.X == x && move.Point.Y == y {
			return true
		}
	}
	return false
}

type Move struct {
	Order int
	Point Point
}
