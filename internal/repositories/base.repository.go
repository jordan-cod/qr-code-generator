package repositories

import (
	"context"
	"log"

	"gorm.io/gorm"
)

type repository[T any] struct {
	db *gorm.DB
}

func NewRepository[T any](db *gorm.DB) *repository[T] {
	if db == nil {
		log.Fatal("Banco de dados não inicializado corretamente")
	}
	return &repository[T]{
		db: db,
	}
}

func (r *repository[T]) Create(ctx context.Context, entity *T) error {
	return r.db.WithContext(ctx).Create(entity).Error
}

func (r *repository[T]) Update(ctx context.Context, id string, entity *T) error {
	return r.db.WithContext(ctx).Model(&entity).Where("id = ?", id).Updates(entity).Error
}

func (r *repository[T]) Delete(ctx context.Context, id string, entity *T) error {
	return r.db.WithContext(ctx).Delete(&entity, "id = ?", id).Error
}

func (r *repository[T]) GetAll(ctx context.Context) (*[]T, error) {
	var entities []T
	err := r.db.WithContext(ctx).Find(&entities).Error
	if err != nil {
		return nil, err
	}
	return &entities, nil
}

func (r *repository[T]) GetByID(ctx context.Context, id string) (*T, error) {
	var entity T
	err := r.db.WithContext(ctx).First(&entity, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *repository[T]) GetBy(ctx context.Context, conditions map[string]interface{}) (*T, error) {
	var entity T
	err := r.db.WithContext(ctx).Where(conditions).First(&entity).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}
