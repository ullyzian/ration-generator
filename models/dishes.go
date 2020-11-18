package models

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
)

type Dish struct {
	Id       int
	Title    string
	Portion  int
	Calories int
}

func (dish *Dish) String() string {
	return fmt.Sprintf("Dish %s", dish.Title)
}

func GetDishes(conn *sql.DB) ([]Dish, error) {
	rows, err := conn.Query("SELECT * FROM dishes")

	if err != nil {
		return nil, err
	}
	var dishes []Dish

	for rows.Next() {
		var dish Dish
		if err := rows.Scan(&dish.Id, &dish.Title, &dish.Portion, &dish.Calories); err != nil {
			return nil, err
		}
		dishes = append(dishes, dish)
	}
	return dishes, nil
}

func CreateDish(conn *sql.DB, title string, portion int, calories int) (*Dish, error) {
	cmd := "INSERT INTO dishes(title, portion, calories) VALUES($1, $2, $3) RETURNING id"
	var id int
	err := conn.QueryRow(cmd, title, portion, calories).Scan(&id)
	if err != nil {
		return nil, err
	}
	return &Dish{Id: int(id), Title: title, Portion: portion, Calories: calories}, nil
}
