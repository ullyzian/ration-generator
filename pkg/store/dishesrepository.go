package store

import (
	"github.com/lib/pq"
	"github.com/ullyzian/ration-generator/pkg/models"
)

type DishesRepository struct {
	store *Store
}

func (r *DishesRepository) Create(d *models.Dish) (*models.Dish, error) {
	cmd := "INSERT INTO dishes(title, portion, calories, contraindication) VALUES($1, $2, $3, $4) RETURNING id"
	if err := r.store.db.QueryRow(cmd, d.Title, d.Portion, d.Calories, d.Contraindications).Scan(&d.Id); err != nil {
		return nil, err
	}
	return d, nil
}

func (r *DishesRepository) Delete(id int) (int, error) {
	cmd1 := "DELETE FROM programs_dishes WHERE dish_id=$1"
	if _, err := r.store.db.Exec(cmd1, id); err != nil {
		return 0, err
	}
	cmd2 := "DELETE FROM dishes WHERE id = $1"
	if _, err := r.store.db.Exec(cmd2, id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *DishesRepository) Edit(d *models.Dish) (*models.Dish, error) {
	cmd := "UPDATE dishes SET title=$1, calories=$2, portion=$3, contraindication=$4 WHERE id=$5"
	if _, err := r.store.db.Exec(cmd, d.Title, d.Calories, d.Portion, d.Contraindications, d.Id); err != nil {
		return nil, err
	}
	return d, nil
}

func (r *DishesRepository) GetAll() ([]models.Dish, error) {
	rows, err := r.store.db.Query("SELECT * FROM dishes ORDER BY id DESC")

	if err != nil {
		return nil, err
	}
	var dishes []models.Dish

	for rows.Next() {
		var dish models.Dish
		if err := rows.Scan(&dish.Id, &dish.Title, &dish.Portion, &dish.Calories, &dish.Contraindications); err != nil {
			return nil, err
		}
		dishes = append(dishes, dish)
	}
	return dishes, nil
}

func (r *DishesRepository) GetById(id int) (*models.Dish, error) {
	dish := models.Dish{}
	query := "SELECT id, title, portion, calories, contraindication FROM dishes WHERE id = $1"
	if err := r.store.db.QueryRow(query, id).Scan(&dish.Id, &dish.Title, &dish.Portion, &dish.Calories, &dish.Contraindications); err != nil {
		return nil, err
	}
	return &dish, nil
}

func (r *DishesRepository) GetByIds(ids []string) ([]models.Dish, error) {
	rows, err := r.store.db.Query("SELECT * FROM dishes WHERE id = any($1)", pq.Array(ids))

	if err != nil {
		return nil, err
	}
	var dishes []models.Dish

	for rows.Next() {
		var dish models.Dish
		if err := rows.Scan(&dish.Id, &dish.Title, &dish.Portion, &dish.Calories, &dish.Contraindications); err != nil {
			return nil, err
		}
		dishes = append(dishes, dish)
	}
	return dishes, nil
}
