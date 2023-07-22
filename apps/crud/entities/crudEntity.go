package entities

import "time"

type CrudEntity struct {
	ID        uint   `json:"id" gorm:"primary_key"` // gorm:"primary_key" tells gorm that this field is the primary key
	Title     string `json:"title"`
	CreatedAt time.Time
}

type CrudSerializer struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}

func ResponseCrud(crud CrudEntity) CrudSerializer {
	return CrudSerializer{
		ID:        crud.ID,
		Title:     crud.Title,
		CreatedAt: crud.CreatedAt,
	}
}
