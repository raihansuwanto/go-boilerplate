package product

import (
	"github.com/google/wire"
	"github.com/raihansuwanto/go-boilerplate/internal/app/product/delivery/web"
	"github.com/raihansuwanto/go-boilerplate/internal/app/product/repo"
	"github.com/raihansuwanto/go-boilerplate/internal/app/product/repo/gopg"
	"github.com/raihansuwanto/go-boilerplate/internal/app/product/usecase"
	"github.com/raihansuwanto/go-boilerplate/internal/app/product/usecase/usecaseimpl"
)

var ModuleSet = wire.NewSet(
	repositories,
	usecases,
	webservice,
)

var repositories = wire.NewSet(
	gopg.NewCategoryRepoPG,
	wire.Bind(new(repo.CategoryRepo), new(*gopg.CategoryRepoPG)),
	gopg.NewProductRepoPG,
	wire.Bind(new(repo.ProductRepo), new(*gopg.ProductRepoPG)),
)

var usecases = wire.NewSet(
	usecaseimpl.NewCategory,
	wire.Bind(new(usecase.Category), new(*usecaseimpl.CategoryImpl)),
	usecaseimpl.NewProduct,
	wire.Bind(new(usecase.Product), new(*usecaseimpl.ProductImpl)),
)

var webservice = wire.NewSet(
	web.NewProductService,
	web.NewProductRegistry,

	// pgddl.NewModeler,
	// wire.Bind(new(ddl.Modeler), new(*pgddl.Modeler)),
)
