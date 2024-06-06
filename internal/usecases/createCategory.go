package usecases

import (
	"github.com/silvajuniorow/microAPI/internal/entities"
	"github.com/silvajuniorow/microAPI/internal/repositories"
)

type createCategoryUseCase struct {
	repository repositories.ICategoryRepository
}

func NewCreateCategoryUseCase(repository repositories.ICategoryRepository) *createCategoryUseCase {
	return &createCategoryUseCase{repository}
}

func (u *createCategoryUseCase) Execute(name string) error {
	category, err := entities.NewCategory(name)

	if err != nil {
		return err
	}

	err = u.repository.Save(category)

	if err != nil {
		return err
	}

	return nil
}
