package gopg

import (
	"github.com/go-pg/pg/v10"
	"github.com/raihansuwanto/go-boilerplate/internal/app/ent"
	"github.com/raihansuwanto/go-boilerplate/internal/app/helper/db"
)

type ProductRepoPG struct {
	db.GenericRepository[ent.Product]
}

func NewProductRepoPG(pgdb *pg.DB) *ProductRepoPG {
	return &ProductRepoPG{
		GenericRepository: db.NewGenericRepository[ent.Product](pgdb),
	}
}
