//go:build wireinject
// +build wireinject

package app

import (
	"context"

	"github.com/google/wire"
	"github.com/raihansuwanto/go-boilerplate/internal/app/product"
	"github.com/raihansuwanto/go-boilerplate/internal/package/config"
	"github.com/raihansuwanto/go-boilerplate/internal/package/postgre"
	"github.com/raihansuwanto/go-boilerplate/internal/package/postgre/pgddl"
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
