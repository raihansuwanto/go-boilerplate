package dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/raihansuwanto/go-boilerplate/app/ent"
)

type ProductCreatorRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CategoryID  string  `json:"category_id"`
}

func (e *ProductCreatorRequest) Validate() error {
	return validation.ValidateStruct(e,
		validation.Field(&e.Name, validation.Required, validation.Length(3, 50)),
		validation.Field(&e.Description, validation.Required, validation.Length(3, 250)),
		validation.Field(&e.Price, validation.Required),
		validation.Field(&e.CategoryID, validation.Required),
	)
}

func (e *ProductCreatorRequest) RequestToEntity() ent.Product {
	return ent.Product{
		Name:        e.Name,
		Description: e.Description,
		Price:       e.Price,
		CategoryID:  e.CategoryID,
	}
}

type ProductCreatorResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
