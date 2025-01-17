package web

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/raihansuwanto/go-boilerplate/app/product/usecase"
	"github.com/raihansuwanto/go-boilerplate/app/product/usecase/dto"
	"github.com/raihansuwanto/go-boilerplate/package/errors"
)

func CategoryCreator(handler usecase.Category) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		request := dto.CategoryCreatorRequest{}

		if err := render.DecodeJSON(r.Body, &request); err != nil {
			errors.RenderError(r, w, errors.NewBadRequestError())
			return
		}

		result, err := handler.Create(ctx, &request)
		if err != nil {
			errors.RenderError(r, w, err)
			return
		}
		render.Status(r, http.StatusOK)
		render.JSON(w, r, result)
	}
}

func CategoryLoader(handler usecase.Category) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()

		request := dto.CategoryLoaderRequest{
			ID: chi.URLParam(r, "id"),
		}

		result, err := handler.GetDetail(ctx, &request)
		if err != nil {
			errors.RenderError(r, w, err)
			return
		}
		render.Status(r, http.StatusOK)
		render.JSON(w, r, result)
	}
}

func ProductCreator(handler usecase.Product) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		request := dto.ProductCreatorRequest{}

		if err := render.DecodeJSON(r.Body, &request); err != nil {
			errors.RenderError(r, w, errors.NewBadRequestError())
			return
		}

		result, err := handler.Create(ctx, &request)
		if err != nil {
			errors.RenderError(r, w, err)
			return
		}
		render.Status(r, http.StatusOK)
		render.JSON(w, r, result)
	}
}

func ProductLoader(handler usecase.Product) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()

		request := dto.ProductLoaderRequest{
			ID: chi.URLParam(r, "id"),
		}

		result, err := handler.GetDetail(ctx, &request)
		if err != nil {
			errors.RenderError(r, w, err)
			return
		}
		render.Status(r, http.StatusOK)
		render.JSON(w, r, result)
	}
}
