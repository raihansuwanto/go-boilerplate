package pgddl

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/raihansuwanto/go-boilerplate/internal/package/postgre"
)

type Modeler struct {
	db orm.DB
}

// compile-time type checking, in order to ensure that Modeler implement ddl.Modeler interface.
var _ postgre.Modeler = (*Modeler)(nil)

func NewModeler(db *pg.DB) *Modeler {
	return &Modeler{db: db}
}

func (m Modeler) Model(model interface{}) postgre.Model {
	return newDDLModel(m.db.Model(model))
}
