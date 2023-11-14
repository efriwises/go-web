package entities

import "time"

type ProductEntities struct {
	Id          uint
	Name        string
	Category    CategoryEntities
	Stock       int64
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}