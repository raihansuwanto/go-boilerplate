package app

import (
	productweb "github.com/raihansuwanto/go-boilerplate/internal/app/product/delivery/web"
	"github.com/raihansuwanto/go-boilerplate/runner/web"
)

func makeWebServiceRegistries(
	productRegistry *productweb.ProductRegistry,

) []web.WebModuleRegistry {

	var webRegistries []web.WebModuleRegistry
	productRegistry.AppendToWebRegistries(&webRegistries)

	return webRegistries
}
