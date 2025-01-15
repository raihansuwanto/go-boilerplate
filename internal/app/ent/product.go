package ent

import "time"

type Product struct {
	ID          string    `pg:",pk,default:uuid_generate_v4(),type:uuid"`
	Name        string    `pg:",notnull"`
	Description string    `pg:","`
	CategoryID  string    `pg:","`
	Price       float64   `pg:","`
	CreatedAt   time.Time `pg:",default:now(),notnull"`
	UpdatedAt   time.Time `pg:",default:now(),notnull"`
}
