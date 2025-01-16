package pgddl

import (
	"github.com/go-pg/pg/v10/orm"
	"github.com/raihansuwanto/go-boilerplate/package/postgre"
)

type DDLModel struct {
	modelQuery *orm.Query
}

// compile-time type checking, in order to ensure that DDLModel implement ddl.Model interface.
var _ postgre.Model = (*DDLModel)(nil)

func newDDLModel(modelQuery *orm.Query) *DDLModel {
	return &DDLModel{modelQuery: modelQuery}
}

func (d DDLModel) CreateTable(options *postgre.TableCreationOptions) error {
	err := d.modelQuery.CreateTable(&orm.CreateTableOptions{
		Varchar:       255,
		Temp:          options.Temp,
		IfNotExists:   options.IfNotExists,
		FKConstraints: options.FKConstraints,
	})

	if err != nil {
		return err
	}

	return nil
}
