package usecase

import (
	"context"

	"github.com/raihansuwanto/go-boilerplate/app/product/usecase/dto"
)

type Category interface {
	Create(ctx context.Context, req *dto.CategoryCreatorRequest) (*dto.CategoryCreatorResponse, error)
	GetDetail(ctx context.Context, req *dto.CategoryLoaderRequest) (*dto.CategoryLoaderResponse, error)
}

type Product interface {
	Create(ctx context.Context, cmd *dto.ProductCreatorRequest) (*dto.ProductCreatorResponse, error)
	GetDetail(ctx context.Context, cmd *dto.ProductLoaderRequest) (*dto.ProductLoaderResponse, error)
}
