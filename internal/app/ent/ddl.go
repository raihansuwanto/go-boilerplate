package ent

import (
	"github.com/raihansuwanto/go-boilerplate/internal/package/postgre"
)

func MustCreateSchemaIfNotExist(modeler postgre.Modeler) {
	entities := []interface{}{
		(*Category)(nil),
		(*Product)(nil),
		(*User)(nil),
		(*Store)(nil),
	}

	for _, entity := range entities {
		err := modeler.Model(entity).CreateTable(&postgre.TableCreationOptions{
			IfNotExists:   true,
			Temp:          false,
			FKConstraints: false,
		})
		if err != nil {
			panic(err)
		}
	}
}
