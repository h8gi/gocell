package gocam

import fifth "github.com/h8gi/fifth/lib"

type Cell int

type Neighborhood [20]Cell

type RuleTable map[Neighborhood]Cell

type CAM struct {
	forth   *fifth.Interpreter
	Table   RuleTable
	program String
}

func (cam *CAM) FillRuleTable() error {

}
