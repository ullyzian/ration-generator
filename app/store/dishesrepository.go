package store

import (
	"github.com/lib/pq"
	"github.com/ullyzian/ration-generator/app/models"
)

type DishesRepository struct {
	store *Store
}

func (r *DishesRepository) Create(d *models.Dish) (*models.Dish, error) {
	cmd := "INSERT INTO dishes(title, portion, calories) VALUES($1, $2, $3) RETURNING id"
	if err := r.store.db.QueryRow(cmd, d.Title, d.Portion, d.Calories).Scan(&d.Id); err != nil {
		return nil, err
	}
	return d, nil
}

func (r *DishesRepository) GetAll() ([]models.Dish, error) {
	rows, err := r.store.db.Query("SELECT * FROM dishes")

	if err != nil {
		return nil, err
	}
	var dishes []models.Dish

	for rows.Next() {
		var dish models.Dish
		if err := rows.Scan(&dish.Id, &dish.Title, &dish.Portion, &dish.Calories); err != nil {
			return nil, err
		}
		dishes = append(dishes, dish)
	}
	return dishes, nil
}

func (r *DishesRepository) GetByIds(ids []string) ([]models.Dish, error) {
	rows, err := r.store.db.Query("SELECT * FROM dishes WHERE id = any($1)", pq.Array(ids))

	if err != nil {
		return nil, err
	}
	var dishes []models.Dish

	for rows.Next() {
		var dish models.Dish
		if err := rows.Scan(&dish.Id, &dish.Title, &dish.Portion, &dish.Calories); err != nil {
			return nil, err
		}
		dishes = append(dishes, dish)
	}
	return dishes, nil
}
