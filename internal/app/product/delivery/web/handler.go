package web

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/raihansuwanto/go-boilerplate/internal/app/product/usecase"
	"github.com/raihansuwanto/go-boilerplate/internal/app/product/usecase/dto"
)

func CategoryCreator(handler usecase.Category) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		request := dto.CategoryCreatorRequest{}

		if err := render.DecodeJSON(r.Body, &request); err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err)
			return
		}

		result, err := handler.Create(ctx, &request)
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, err)
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

		result, err := handler.Load(ctx, &request)
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, err)
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
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, err)
			return
		}

		result, err := handler.Create(ctx, &request)
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, err)
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

		result, err := handler.Load(ctx, &request)
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, err)
			return
		}
		render.Status(r, http.StatusOK)
		render.JSON(w, r, result)
	}
}
