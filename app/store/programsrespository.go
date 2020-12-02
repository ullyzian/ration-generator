package store

import "github.com/ullyzian/ration-generator/app/models"

type ProgramsRepository struct {
	store *Store
}

func (r *ProgramsRepository) Create(p *models.Program) (*models.Program, error) {
	cmd := "INSERT INTO programs(name) VALUES($1) RETURNING id"
	if err := r.store.db.QueryRow(cmd, p.Name).Scan(&p.Id); err != nil {
		return nil, err
	}
	cmd2 := "INSERT INTO programs_dishes(program_id, dish_id) VALUES ($1, $2)"
	for _, d := range p.Dishes {
		if err := r.store.db.QueryRow(cmd2, p.Id, d.Id).Scan(); err != nil {
			return nil, err
		}
	}
	return p, nil
}

func (r *ProgramsRepository) AddDishes(p *models.Program, dishes []models.Dish) (*models.Program, error) {
	cmd := "INSERT INTO programs_dishes(program_id, dish_id) VALUES($1, $2)"
	for _, d := range dishes {
		if err := r.store.db.QueryRow(cmd, p.Id, d.Id).Scan(); err != nil {
			return nil, err
		}
		p.Dishes = append(p.Dishes, d)
	}
	return p, nil
}

//func (r *ProgramsRepository) GetAll() ([]models.Program, error) {
//	cmd := "SELECT * FROM programs p JOIN programs_dishes pd ON p.id = pd.program_id JOIN dishes d ON d.id = pd.dish_id"
//	rows, err := r.store.db.Query(cmd)
//	if err != nil {
//		return nil, err
//	}
//	var programs []models.Program
//
//	for rows.Next() {
//		var program models.Program
//
//	}
//	return nil, nil
//}
