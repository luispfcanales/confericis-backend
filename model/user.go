package model

import "time"

type User struct {
	ID        string
	Email     string
	Password  string
	Name      string
	Role      *Role  // Relación con Role
	RoleID    string // ID del rol
	CreatedAt time.Time
	UpdatedAt time.Time
}
