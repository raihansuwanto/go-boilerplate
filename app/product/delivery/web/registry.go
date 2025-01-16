package web

import "github.com/raihansuwanto/go-boilerplate/runner/web"

type ProductRegistry struct {
	ProductWebRegistry *ProductService
}

func NewProductRegistry(productWebRegistry *ProductService) *ProductRegistry {
	return &ProductRegistry{
		ProductWebRegistry: productWebRegistry,
	}
}

func (r *ProductRegistry) AppendToWebRegistries(webRegistries *[]web.WebModuleRegistry) {
	*webRegistries = append(*webRegistries, (r.ProductWebRegistry))
}
