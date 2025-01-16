package dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
	"github.com/raihansuwanto/go-boilerplate/app/ent"
)

type CategoryCreatorRequest struct {
	Name string `json:"name"`
}

func (e *CategoryCreatorRequest) Validate() error {
	return validation.ValidateStruct(e,
		validation.Field(&e.Name, validation.Required),
	)
}

func (e *CategoryCreatorRequest) RequestToEntity() ent.Category {
	return ent.Category{
		ID:   uuid.New().String(),
		Name: e.Name,
	}
}

type CategoryCreatorResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
