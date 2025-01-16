package usecaseimpl

import (
	"context"

	"github.com/raihansuwanto/go-boilerplate/app/helper/db"
	"github.com/raihansuwanto/go-boilerplate/app/product/repo"
	"github.com/raihansuwanto/go-boilerplate/app/product/usecase/dto"
	"github.com/sirupsen/logrus"
)

type ProductImpl struct {
	categoryRepo repo.CategoryRepo
	productRepo  repo.ProductRepo
}

func NewProduct(categoryRepo repo.CategoryRepo, productRepo repo.ProductRepo) *ProductImpl {
	return &ProductImpl{
		categoryRepo: categoryRepo,
		productRepo:  productRepo,
	}
}

func (p *ProductImpl) Create(ctx context.Context, cmd *dto.ProductCreatorRequest) (*dto.ProductCreatorResponse, error) {

	if err := cmd.Validate(); err != nil {
		return nil, err
	}

	_, err := p.categoryRepo.Load(ctx, db.Filter{Field: "id", Value: cmd.CategoryID})
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("failed to load category")
		return nil, err
	}

	product := cmd.RequestToEntity()

	if err := p.productRepo.Create(ctx, &product); err != nil {
		logrus.WithContext(ctx).WithError(err).Error("failed to create product")
		return nil, err
	}

	return &dto.ProductCreatorResponse{
		ID:   product.ID,
		Name: product.Name,
	}, nil
}

func (p *ProductImpl) Load(ctx context.Context, cmd *dto.ProductLoaderRequest) (*dto.ProductLoaderResponse, error) {
	product, err := p.productRepo.Load(ctx, db.Filter{Field: "id", Value: cmd.ID})
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("failed to load product")
		return nil, err
	}

	category, err := p.categoryRepo.Load(ctx, db.Filter{Field: "id", Value: product.CategoryID})
	if err != nil {
		logrus.WithContext(ctx).WithError(err).Error("failed to load category")
		return nil, err
	}

	return &dto.ProductLoaderResponse{
		ID:           product.ID,
		Name:         product.Name,
		Description:  product.Description,
		Price:        product.Price,
		CategoryName: category.Name,
		CategoryID:   product.CategoryID,
	}, nil
}
