package repo

import (
	"github.com/raihansuwanto/go-boilerplate/app/ent"
	"github.com/raihansuwanto/go-boilerplate/app/helper/db"
)

type CategoryRepo interface {
	db.GenericRepository[ent.Category]
}

type ProductRepo interface {
	db.GenericRepository[ent.Product]
}
