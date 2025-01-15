package ent

import "time"

type Store struct {
	ID          string    `pg:",pk,default:uuid_generate_v4(),type:uuid"`
	UserID      string    `pg:","`
	Name        string    `pg:",notnull"`
	Description string    `pg:","`
	CategoryID  string    `pg:","`
	CreatedAt   time.Time `pg:",default:now(),notnull"`
	UpdatedAt   time.Time `pg:",default:now(),notnull"`
}
