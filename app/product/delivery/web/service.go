package web

import (
	"github.com/go-chi/chi"
	"github.com/raihansuwanto/go-boilerplate/app/product/usecase"
)

type ProductService struct {
	categoryHandler usecase.Category
	productHandler  usecase.Product
}

func NewProductService(categoryHandler usecase.Category, productHandler usecase.Product) *ProductService {
	return &ProductService{
		categoryHandler: categoryHandler,
		productHandler:  productHandler,
	}
}

func (p *ProductService) RegisterRoutesTo(router chi.Router) error {

	router.Route("/api/v1", func(r chi.Router) {

		r.Route("/category", func(r chi.Router) {
			r.Post("/", CategoryCreator(p.categoryHandler))
			r.Get("/{id}", CategoryLoader(p.categoryHandler))
		})

		r.Route("/product", func(r chi.Router) {
			r.Post("/", ProductCreator(p.productHandler))
			r.Get("/{id}", ProductLoader(p.productHandler))
		})

	})

	return nil
}
