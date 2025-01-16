package usecaseimpl

import (
	"context"

	"github.com/raihansuwanto/go-boilerplate/app/helper/db"
	"github.com/raihansuwanto/go-boilerplate/app/product/repo"
	"github.com/raihansuwanto/go-boilerplate/app/product/usecase/dto"
	"github.com/sirupsen/logrus"
)

type CategoryImpl struct {
	categoryRepo repo.CategoryRepo
}

func NewCategory(categoryRepo repo.CategoryRepo) *CategoryImpl {
	return &CategoryImpl{
		categoryRepo: categoryRepo,
	}
}

func (e *CategoryImpl) Create(ctx context.Context, req *dto.CategoryCreatorRequest) (*dto.CategoryCreatorResponse, error) {
	category := req.RequestToEntity()
	if err := e.categoryRepo.Create(ctx, &category); err != nil {
		logrus.WithError(err).Error("failed to create category")
		return nil, err
	}

	return &dto.CategoryCreatorResponse{
		ID:   category.ID,
		Name: category.Name,
	}, nil
}

func (e *CategoryImpl) Load(ctx context.Context, req *dto.CategoryLoaderRequest) (*dto.CategoryLoaderResponse, error) {
	category, err := e.categoryRepo.Load(ctx, db.Filter{Field: "id", Value: req.ID})
	if err != nil {
		logrus.WithError(err).Error("failed to load category")
		return nil, err
	}

	return &dto.CategoryLoaderResponse{
		ID:   category.ID,
		Name: category.Name,
	}, nil
}
