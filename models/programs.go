package models

type Program struct {
	name string
	ingestions []Ingestion
}

type Ingestion struct {
	title string
	dishes []Dish
}
