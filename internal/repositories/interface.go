package repositories

import "github.com/silvajuniorow/microAPI/internal/entities"

type ICategoryRepository interface {
	Save(category *entities.Category) error
	List() ([]*entities.Category, error)
}
