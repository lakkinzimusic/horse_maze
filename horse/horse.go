package horse

import "my_projects/horse_maze/maze"

//Mover inteface
type Mover interface {
	Turn()
}

//Horse struct
type Horse struct {
	StartPoint maze.Point
}
