package gocell

import "math/rand"

type Cell struct {
	State int
}

type World struct {
	Lattice       *[][]Cell
	tmpLattice    *[][]Cell
	Width         int
	Height        int
	RuleFunc      RuleFunc
	NeighborsFunc NeighborsFunc
}

type RuleFunc func(self Cell, neighbors []Cell) Cell

type NeighborsFunc func(x, y int) []Cell

func NewWorld(width, height int) *World {
	lattice := make([][]Cell, height)
	tmplattice := make([][]Cell, height)
	for i, _ := range lattice {
		lattice[i] = make([]Cell, width)
		tmplattice[i] = make([]Cell, width)
	}
	return &World{
		Lattice:    &lattice,
		tmpLattice: &tmplattice,
		Width:      width,
		Height:     height,
	}
}

func (w *World) SetRuleFunc(fun RuleFunc) {
	w.RuleFunc = fun
}

func (w *World) SetNeighborsFunc(fun NeighborsFunc) {
	w.NeighborsFunc = fun
}

func (w *World) At(x, y int) Cell {
	nx := x % w.Width
	if nx < 0 {
		nx = nx + w.Width
	}
	ny := y % w.Height
	if ny < 0 {
		ny = ny + w.Height
	}
	return (*w.Lattice)[ny][nx]
}

func (w *World) Randomize(n int) {
	for y := 0; y < w.Height; y++ {
		for x := 0; x < w.Width; x++ {
			(*w.Lattice)[y][x].State = rand.Intn(n)
		}
	}
}

func (w *World) OneStep() {
	for y := 0; y < w.Height; y++ {
		for x := 0; x < w.Width; x++ {
			neighbors := w.NeighborsFunc(x, y)
			(*w.tmpLattice)[y][x] = w.RuleFunc(w.At(x, y), neighbors)
		}
	}
	w.Lattice, w.tmpLattice = w.tmpLattice, w.Lattice
}
