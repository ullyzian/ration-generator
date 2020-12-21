package models

import (
	"fmt"
)

type Dish struct {
	Id       int
	Title    string
	Portion  int
	Calories int
	Contraindications string
}

func (dish *Dish) String() string {
	return fmt.Sprintf("Dish %s", dish.Title)
}
