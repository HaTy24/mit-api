package baseService

import (
	"gorm.io/gorm"
)

// Định nghĩa BaseService chung với generics
type BaseService[T any] struct {
	DB *gorm.DB
}

// Create entity
func (c *BaseService[T]) Create(entity *T) error {
	return c.DB.Create(entity).Error
}

// Find entity by ID
func (c *BaseService[T]) FindById(id uint, entity *T) error {
	return c.DB.First(entity, id).Error
}

// Update entity
func (c *BaseService[T]) Update(entity *T) error {
	return c.DB.Save(entity).Error
}

// Delete entity
func (c *BaseService[T]) Delete(id uint, entity *T) error {
	return c.DB.Delete(entity, id).Error
}
