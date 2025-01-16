package ent

import "time"

type Category struct {
	ID        string    `pg:",pk,default:uuid_generate_v4(),type:uuid"`
	Name      string    `pg:",notnull"`
	CreatedAt time.Time `pg:",default:now(),notnull"`
	UpdatedAt time.Time `pg:",default:now(),notnull"`
}
