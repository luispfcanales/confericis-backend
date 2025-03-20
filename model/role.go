package model

import "time"

type Roles struct {
	ID          string
	Name        string
	Status      bool
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}
