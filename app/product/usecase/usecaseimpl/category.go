package usecaseimpl

import (
	"context"

	"github.com/raihansuwanto/go-boilerplate/app/helper/db"
	"github.com/raihansuwanto/go-boilerplate/app/product/repo"
	"github.com/raihansuwanto/go-boilerplate/app/product/usecase/dto"
	"github.com/raihansuwanto/go-boilerplate/package/logger"
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

	if err := req.Validate(); err != nil {
		return nil, err
	}

	category := req.RequestToEntity()
	if err := e.categoryRepo.Create(ctx, &category); err != nil {
		logger.WithContext(ctx).WithError(err).Error("failed to create category")
		return nil, err
	}

	return &dto.CategoryCreatorResponse{
		ID:   category.ID,
		Name: category.Name,
	}, nil
}

func (e *CategoryImpl) GetDetail(ctx context.Context, req *dto.CategoryLoaderRequest) (*dto.CategoryLoaderResponse, error) {
	category, err := e.categoryRepo.Load(ctx, db.Filter{Field: "id", Value: req.ID})
	if err != nil {
		logger.WithContext(ctx).WithError(err).Error("failed to load category")
		return nil, err
	}

	return &dto.CategoryLoaderResponse{
		ID:   category.ID,
		Name: category.Name,
	}, nil
}
