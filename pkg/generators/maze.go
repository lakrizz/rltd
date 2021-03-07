package generators

import (
	"math/rand"
)

const (
	TILE_EMPTY = iota
	TILE_START
	TILE_END
	TILE_PATH
	TILE_OFFENSIVE
	TILE_SUPPORT
	TILE_DEFENSIVE

	Up = 1 << iota
	Down
	Left
	Right
)

// Point on the maze
type Point struct {
	X, Y int
}

type Tile struct {
	Type int
}

// Directions is the set of all the directions
var Directions = []int{Up, Down, Left, Right}
var Opposite = map[int]int{Up: Down, Down: Up, Left: Right, Right: Left}

type Maze struct {
	Width  int
	Height int

	Directions [][]int
	Tiles      [][]*Tile
	Start      *Point
	Goal       *Point
	Solved     bool
}

var (
	dx = map[int]int{Up: -1, Down: 1, Left: 0, Right: 0}
	dy = map[int]int{Up: 0, Down: 0, Left: -1, Right: 1}
)

func GenerateMaze(width, height int, seed int64) *Maze {
	rand.Seed(seed)
	mm := &Maze{}

	var directions [][]int
	for x := 0; x < height; x++ {
		directions = append(directions, make([]int, width))
	}

	mm.Directions = directions
	mm.Width = width
	mm.Height = height
	mm.generateStartAndEnd()
	mm.SetTiles()
	mm.Generate()

	// fmt.Println(fmt.Sprintf("Directions[%v][%v]", len(mm.Directions), len(mm.Directions[0])))
	// fmt.Println(fmt.Sprintf("Tiles[%v][%v]", len(mm.Tiles), len(mm.Tiles[0])))

	mm.Solve()
	mm.Tiles[mm.Start.X][mm.Start.Y].Type = TILE_START
	mm.Tiles[mm.Goal.X][mm.Goal.Y].Type = TILE_END
	return mm
}

func (m *Maze) SetTiles() {
	m.Tiles = make([][]*Tile, m.Height)
	for i := 0; i < m.Height; i++ {
		m.Tiles[i] = make([]*Tile, m.Width)
		for j := 0; j < m.Width; j++ {
			m.Tiles[i][j] = &Tile{TILE_EMPTY}
		}
	}

}

// @TODO: i'm totally sure that those switch cases could be optimized, code wise
func (m *Maze) generateStartAndEnd() {

	c := rand.Intn(1) // 0 = start left/ goal right, 1 = vice versa
	sx, sy := 0, rand.Intn(m.Height-1)
	gx, gy := m.Width-1, rand.Intn(m.Height-1)

	if c == 0 {
		m.Start = &Point{sy, sx}
		m.Goal = &Point{gy, gx}
	} else {
		m.Start = &Point{gy, gx}
		m.Goal = &Point{sy, sx}
	}
}

// Contains judges whether the argument point is inside the maze or not
func (maze *Maze) Contains(point *Point) bool {
	return 0 <= point.X && point.X < maze.Height && 0 <= point.Y && point.Y < maze.Width
}

// Neighbors gathers the nearest undecided points
func (maze *Maze) Neighbors(point *Point) (neighbors []int) {
	for _, direction := range Directions {
		next := point.Advance(direction)
		if maze.Contains(next) && maze.Directions[next.X][next.Y] == 0 {
			neighbors = append(neighbors, direction)
		}
	}
	return neighbors
}

// Connected judges whether the two points is connected by a path on the maze
func (maze *Maze) Connected(point *Point, target *Point) bool {
	dir := maze.Directions[point.X][point.Y]
	for _, direction := range Directions {
		if dir&direction != 0 {
			next := point.Advance(direction)
			if next.X == target.X && next.Y == target.Y {
				return true
			}
		}
	}
	return false
}

// Next advances the maze path randomly and returns the new point
func (maze *Maze) Next(point *Point) *Point {
	neighbors := maze.Neighbors(point)
	if len(neighbors) == 0 {
		return nil
	}
	direction := neighbors[rand.Int()%len(neighbors)]
	maze.Directions[point.X][point.Y] |= direction
	next := point.Advance(direction)
	maze.Directions[next.X][next.Y] |= Opposite[direction]
	return next
}

// Generate the maze
func (maze *Maze) Generate() {
	point := maze.Start
	stack := []*Point{point}
	for len(stack) > 0 {
		for {
			point = maze.Next(point)
			if point == nil {
				break
			}
			stack = append(stack, point)
		}
		i := rand.Int() % ((len(stack) + 1) / 2)
		point = stack[i]
		stack = append(stack[:i], stack[i+1:]...)
	}
}

// Equal judges the equality of the two points
func (point *Point) Equal(target *Point) bool {
	return point.X == target.X && point.Y == target.Y
}

// Advance the point forward by the argument direction
func (point *Point) Advance(direction int) *Point {
	return &Point{point.X + dx[direction], point.Y + dy[direction]}
}

// Solve the maze
func (maze *Maze) Solve() {
	if maze.Solved {
		return
	}
	point := maze.Start
	stack := []*Point{point}
	solution := []*Point{point}
	visited := 1 << 12
	// Repeat until we reach the goal
	for !point.Equal(maze.Goal) {
		maze.Directions[point.X][point.Y] |= visited
		for _, direction := range Directions {
			// Push the nearest points to the stack if not been visited yet
			if maze.Directions[point.X][point.Y]&direction == direction {
				next := point.Advance(direction)
				if maze.Directions[next.X][next.Y]&visited == 0 {
					stack = append(stack, next)
				}
			}
		}
		// Pop the stack
		point = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// We have reached to a dead end so we pop the solution
		for last := solution[len(solution)-1]; !maze.Connected(point, last); {
			solution = solution[:len(solution)-1]
			last = solution[len(solution)-1]
		}
		solution = append(solution, point)
	}
	maze.Solved = true

	for _, v := range solution {
		maze.Tiles[v.X][v.Y].Type = TILE_PATH
	}
}
