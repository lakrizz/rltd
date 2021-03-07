package generators

import (
	"errors"
	"fmt"
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
)

type Maze struct {
	Width          int
	Height         int
	Tiles          [][]int
	StartX, StartY int
	EndX, EndY     int
}

func GenerateMaze(width, height int, seed int64) (*Maze, error) {
	rand.Seed(seed)
	mm := &Maze{}
	mm.Width = width
	mm.Height = height
	m := make([][]int, height)
	for y := 0; y < height; y++ {
		m[y] = make([]int, width)
		for x := 0; x < width; x++ {
			m[y][x] = TILE_EMPTY
		}
	}
	mm.Tiles = m
	mm.generateStart()
	mm.generateEnd()
	mm.generatePath()
	return mm, nil
}

// @TODO: i'm totally sure that those switch cases could be optimized, code wise
func (m *Maze) generateStart() {
	// [0,1,2,3] = [top, bottom, left, right] corner
	c := rand.Intn(3)
	x, y := 0, 0
	switch c {
	case 0:
		x, y = rand.Intn(m.Width-1), 0
		break
	case 1:
		x, y = rand.Intn(m.Width-1), m.Height-1
		break
	case 2:
		x, y = 0, rand.Intn(m.Height-1)
		break
	case 3:
		x, y = m.Width-1, rand.Intn(m.Height-1)
		break
	}
	m.Tiles[y][x] = TILE_START
	m.StartX = x
	m.StartY = y
}

func (m *Maze) generateEnd() {
	good := false
	for !good {
		// [0,1,2,3] = [top, bottom, left, right] corner
		c := rand.Intn(3)
		x, y := 0, 0
		switch c {
		case 0:
			x, y = rand.Intn(m.Width-1), 0
			break
		case 1:
			x, y = rand.Intn(m.Width-1), m.Height-1
			break
		case 2:
			x, y = 0, rand.Intn(m.Height-1)
			break
		case 3:
			x, y = 8, rand.Intn(m.Height-1)
			break
		}
		if m.Tiles[y][x] != TILE_START {
			m.Tiles[y][x] = TILE_END
			good = true
		}
	}
}

func (m *Maze) generatePath() {
	// now this is where the fun begins, eh? let's start the start
	stepsX := make([]int, 0) // backtracking
	stepsY := make([]int, 0) // backtracking
	sx, sy := m.StartX, m.StartY
	if sx == 0 {
		sx++
	} else if sx == m.Width {
		sx--
	}

	if sy == 0 {
		sy++
	} else if sy == m.Height {
		sy--
	}

	m.Tiles[sy][sx] = TILE_PATH

	done := false

	for !done {
		// [0,1,2,3] = [top,bottom,left,right]
		// stepsX = append(stepsX, sx)
		// stepsY = append(stepsY, sy)
		c := rand.Intn(3)
		switch c {
		case 0:
			if sy > 1 && m.Tiles[sy-1][sx] != TILE_PATH { // we can go here
				sy--
			}
			break
		case 1:
			if sy < m.Height-1 && m.Tiles[sy+1][sx] != TILE_PATH {
				sy++
			}
			break
		case 2:
			if sx > 1 && m.Tiles[sy][sx-1] != TILE_PATH {
				sx--
			}
			break
		case 3:
			if sx < m.Width && m.Tiles[sy][sx+1] != TILE_PATH {
				sx++
			}
			break
		}
		if m.Tiles[sy][sx] == TILE_END {
			done = true
		} else if m.Tiles[sy][sx] == TILE_EMPTY && in(m.countSurrounding(sx, sy, TILE_PATH), []int{0, 1}) {
			m.Tiles[sy][sx] = TILE_PATH
			stepsX = append(stepsX, sx)
			stepsY = append(stepsY, sy)
		} else {
			if len(stepsX) <= 0 || len(stepsY) <= 0 {
				for _, v := range m.Tiles {
					fmt.Println(v)
				}
				panic(errors.New("yeah, big mistake while finding a path"))
			} else {
				sx = stepsX[len(stepsX)-1]
				sy = stepsY[len(stepsY)-1]
				stepsX = stepsX[:len(stepsX)-1]
				stepsY = stepsY[:len(stepsY)-1]
			}
		}
	}
}

func (m *Maze) countSurrounding(x, y, tile int) int {
	// lets find out if none of the surrounding tiles is the one given in the arguments
	sur := 0
	for yy := max(y-1, 0); yy <= min(y+1, m.Height-1); yy++ {
		for xx := max(x-1, 0); xx <= min(x+1, m.Width-1); xx++ {
			if yy != y && xx != x && m.Tiles[yy][xx] == tile {
				sur++
			}
		}
	}
	return sur
}

func in(needle int, haystack []int) bool {
	for _, v := range haystack {
		if needle == v {
			return true
		}
	}
	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
