package gocell

import (
	"fmt"
	"testing"
	"time"
)

func TestLifeGame(t *testing.T) {
	var w = NewWorld(30, 30)
	w.SetNeighborsFunc(func(x, y int) []Cell {
		neighbors := []Cell{}
		for dx := -1; dx < 2; dx++ {
			for dy := -1; dy < 2; dy++ {
				if dx == 0 && dy == 0 {
					continue
				}
				neighbors = append(neighbors, w.At(x+dx, y+dy))
			}
		}
		return neighbors
	})
	w.SetRuleFunc(func(self Cell, neighbors []Cell) Cell {
		sum, next := 0, 0
		for _, c := range neighbors {
			sum += c.State
		}
		if self.State == 0 && sum == 3 {
			next = 1
		}
		if self.State == 1 && (sum == 3 || sum == 2) {
			next = 1
		}
		return Cell{next}
	})
	w.Randomize(2)
	fmt.Println(*w.Lattice)
	for i := 0; i < 50; i++ {
		time.Sleep(1 * time.Second)
		w.OneStep()
		fmt.Println(*w.Lattice)
	}
}
