//go:build wireinject
// +build wireinject

package app

import (
	"context"

	"github.com/google/wire"
	"github.com/raihansuwanto/go-boilerplate/app/product"
	"github.com/raihansuwanto/go-boilerplate/package/config"
	"github.com/raihansuwanto/go-boilerplate/package/postgre"
	"github.com/raihansuwanto/go-boilerplate/package/postgre/pgddl"
)

var ModuleSet = wire.NewSet(

	makeWebServiceRegistries,
	newApp,
	pgddl.NewModeler,
	wire.Bind(new(postgre.Modeler), new(*pgddl.Modeler)),

	config.MakeConfig,

	postgre.MakePostgreDBClient,

	product.ModuleSet,
)

func ProvideApp(ctx context.Context) (*App, error) {
	panic(wire.Build(ModuleSet))
}
