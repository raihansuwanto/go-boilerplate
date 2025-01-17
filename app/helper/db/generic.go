package db

import (
	"context"

	"github.com/go-pg/pg/v10"
	"github.com/raihansuwanto/go-boilerplate/package/errors"
	"github.com/raihansuwanto/go-boilerplate/package/logger"
)

type GenericRepository[T any] interface {
	LoadMany(ctx context.Context, filters ...Filter) ([]*T, error)
	Load(ctx context.Context, filters ...Filter) (*T, error)
	Create(ctx context.Context, entity *T) error
	InsertMany(ctx context.Context, entities []*T) ([]*T, error)
	Delete(ctx context.Context, id interface{}) error
	Update(ctx context.Context, entity *T, filters ...Filter) (*T, error)
	Replace(ctx context.Context, entity *T, filters ...Filter) error
}

// only support for equal operator
// example: SELECT * FROM table WHERE field = value
type Filter struct {
	Field string
	Value interface{}
}

type genericRepository[T any] struct {
	db *pg.DB
}

func NewGenericRepository[T any](db *pg.DB) GenericRepository[T] {
	return &genericRepository[T]{db}
}

func (r *genericRepository[T]) LoadMany(ctx context.Context, filters ...Filter) ([]*T, error) {
	var entities []*T
	query := r.db.WithContext(ctx).Model(&entities)

	for _, filter := range filters {
		query = query.Where(filter.Field+" = ?", filter.Value)
	}

	err := query.Select()
	if err != nil {
		logger.WithContext(ctx).WithError(err).Error("failed to load entities")
		return nil, errors.NewInternalSystemError()
	}

	return entities, nil
}

func (r *genericRepository[T]) Load(ctx context.Context, filters ...Filter) (*T, error) {
	var entity T
	query := r.db.WithContext(ctx).Model(&entity)

	for _, filter := range filters {
		query = query.Where(filter.Field+" = ?", filter.Value)
	}

	err := query.Select()
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, errors.NewDataNotFoundError()
		}
		logger.WithContext(ctx).WithError(err).Error("failed to load entity")
		return nil, errors.NewInternalSystemError().WithMessage(err.Error())
	}

	return &entity, nil
}

func (r *genericRepository[T]) Create(ctx context.Context, entity *T) error {
	_, err := r.db.WithContext(ctx).Model(entity).Insert()
	if err != nil {
		logger.WithContext(ctx).WithError(err).Error("failed to create entity")
		return errors.NewInternalSystemError().WithMessage(err.Error())
	}

	return nil
}

func (r *genericRepository[T]) InsertMany(ctx context.Context, entities []*T) ([]*T, error) {
	_, err := r.db.WithContext(ctx).Model(&entities).Insert()
	if err != nil {
		logger.WithContext(ctx).WithError(err).Error("failed to insert entities")
		return nil, errors.NewInternalSystemError().WithMessage(err.Error())
	}

	return entities, nil
}

func (r *genericRepository[T]) Delete(ctx context.Context, id interface{}) error {
	var entity T
	_, err := r.db.WithContext(ctx).Model(&entity).Where("id = ?", id).Delete()
	if err != nil {
		logger.WithContext(ctx).WithError(err).Error("failed to delete entity")
		return errors.NewInternalSystemError().WithMessage(err.Error())
	}
	return nil
}

func (r *genericRepository[T]) Update(ctx context.Context, entity *T, filters ...Filter) (*T, error) {
	query := r.db.WithContext(ctx).Model(entity)

	for _, filter := range filters {
		query = query.Where(filter.Field+" = ?", filter.Value)
	}

	_, err := query.Update()
	if err != nil {
		logger.WithContext(ctx).WithError(err).Error("failed to update entity")
		return nil, errors.NewInternalSystemError().WithMessage(err.Error())
	}

	return entity, nil
}

func (r *genericRepository[T]) Replace(ctx context.Context, entity *T, filters ...Filter) error {
	query := r.db.WithContext(ctx).Model(entity)

	for _, filter := range filters {
		query = query.Where(filter.Field+" = ?", filter.Value)
	}

	_, err := query.Insert()
	if err != nil {
		logger.WithContext(ctx).WithError(err).Error("failed to replace entity")
		return errors.NewInternalSystemError().WithMessage(err.Error())
	}
	return nil
}
