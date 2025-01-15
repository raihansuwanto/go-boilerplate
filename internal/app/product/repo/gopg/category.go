package gopg

import (
	"github.com/go-pg/pg/v10"
	"github.com/raihansuwanto/go-boilerplate/internal/app/ent"
	"github.com/raihansuwanto/go-boilerplate/internal/app/helper/db"
)

type CategoryRepoPG struct {
	db.GenericRepository[ent.Category]
}

func NewCategoryRepoPG(pgdb *pg.DB) *CategoryRepoPG {
	return &CategoryRepoPG{
		GenericRepository: db.NewGenericRepository[ent.Category](pgdb),
	}
}
