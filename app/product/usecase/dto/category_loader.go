package dto

import (
	"github.com/raihansuwanto/go-boilerplate/app/ent"
)

type CategoryLoaderRequest struct {
	ID string `json:"id"`
}

func (e *CategoryLoaderRequest) RequestToEntity() ent.Category {
	return ent.Category{
		ID: e.ID,
	}
}

type CategoryLoaderResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
