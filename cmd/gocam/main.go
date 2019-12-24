package gocam

import (
	"math"
	"strconv"
	"strings"

	fifth "github.com/h8gi/fifth/lib"
)

type Cell int

type Neighborhood [20]Cell

type RuleTable map[Neighborhood]Cell

type RuleFunc func(Neighborhood) Cell

type CAM struct {
	forth     *fifth.Interpreter
	RuleTable RuleTable
	RuleFunc  RuleFunc
	program   String
	Plane1    *[][]Cell
	Plane2    *[][]Cell
	Width     int
	Height    int
}

func NewCAM(width, height int) *CAM {
	plane1 := make([][]Cell, height)
	plane2 := make([][]Cell, height)
	for i, _ := range plane1 {
		plane1[i] = make([]Cell, width)
		plane2[i] = make([]Cell, width)
	}

	forth := fifth.NewInterpreter()
	return &CAM{
		Plane1: plane1,
		Plane2: plane2,
		forth:  forth,
	}
}

func (cam *CAM) FillRuleTable(stateNum int, neighborsNum int) error {
	n := int(math.Pow(float64(stateNum), float64(neighborsNum)))
	for i := 0; i < n; i++ {
		var nb Neighborhood
		str := strconv.FormatInt(i, stateNum)
		strs := strings.Split(str, "")
		for j, s := range strs {
			x, err := strconv.Atoi(s)
			if err != nil {
				return err
			}
			nb[j] = x
		}
		cam.RuleTable[nb] = cam.RuleFunc(nb)
	}
}

func (cam *CAM) CompileForth(program string) error {

}
