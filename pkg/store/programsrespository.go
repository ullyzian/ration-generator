package store

import (
	"github.com/ullyzian/ration-generator/pkg/models"
)

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
		r.store.db.QueryRow(cmd2, p.Id, d.Id)
	}
	return p, nil
}

func (r *ProgramsRepository) AddDishes(p *models.Program, dishes []models.Dish) (*models.Program, error) {
	cmd := "INSERT INTO programs_dishes(program_id, dish_id) VALUES($1, $2)"
	for _, d := range dishes {
		if _, err := r.store.db.Exec(cmd, p.Id, d.Id); err != nil {
			return nil, err
		}
		p.Dishes = append(p.Dishes, d)
	}
	return p, nil
}

func (r *ProgramsRepository) Delete(id int) (int, error) {
	cmd1 := "DELETE FROM programs_dishes WHERE program_id=$1"
	if _, err := r.store.db.Exec(cmd1, id); err != nil {
		return 0, err
	}
	cmd2 := "DELETE FROM programs WHERE id = $1"
	if _, err := r.store.db.Exec(cmd2, id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *ProgramsRepository) Edit(p *models.Program) (*models.Program, error) {
	cmd := "UPDATE programs SET name=$1 WHERE id=$2"
	if _, err := r.store.db.Exec(cmd, p.Name); err != nil {
		return nil, err
	}
	return p, nil
}

func (r *ProgramsRepository) GetAll() ([]models.Program, error) {
	cmd := "SELECT p.id AS program_id, p.name AS program_name, d.id AS dish_id, d.title AS dish_title, d.portion AS dish_portion, d.calories AS dish_calories FROM programs p JOIN programs_dishes pd ON p.id = pd.program_id JOIN dishes d ON d.id = pd.dish_id"
	rows, err := r.store.db.Query(cmd)
	if err != nil {
		return nil, err
	}
	pMap := make(map[int]models.Program)

	for rows.Next() {
		p := models.Program{}
		d := models.Dish{}
		if err := rows.Scan(&p.Id, &p.Name, &d.Id, &d.Title, &d.Portion, &d.Calories); err != nil {
			return nil, err
		}
		if val, ok := pMap[p.Id]; ok {
			val.Dishes = append(val.Dishes, d)
			pMap[p.Id] = val
		} else {
			p.Dishes = append(p.Dishes, d)
			pMap[p.Id] = p
		}
	}
	var programs []models.Program
	for _, value := range pMap {
		programs = append(programs, value)
	}

	return programs, nil
}

//func (r *ProgramsRepository) GetById() (*models.Program, error) {
//	cmd := "SELECT p.id AS program_id, p.name AS program_name, d.id AS dish_id, d.title AS dish_title, d.portion AS dish_portion, d.calories AS dish_calories FROM programs p JOIN programs_dishes pd ON p.id = pd.program_id JOIN dishes d ON d.id = pd.dish_id WHERE program_id=$1"
//	rows, err := r.store.db.Query(cmd)
//	if err != nil {
//		return nil, err
//	}
//	p := models.Program{}
//
//}
