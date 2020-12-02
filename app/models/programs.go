package models

import (
	"fmt"
)

type Program struct {
	Id     int
	Name   string
	Dishes []Dish
}

func (program *Program) String() string {
	return fmt.Sprintf("Programs %s", program.Name)
}
