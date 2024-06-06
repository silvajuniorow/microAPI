package entities

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
}

func NewCategory(name string) (*Category, error) {
	category := &Category{
		ID:        uuid.New(),
		Name:      name,
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
	}

	// Business rules
	err := category.IsValid()
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (c *Category) IsValid() error {
	if len(c.Name) < 5 {
		return fmt.Errorf("name most be greater thab 5, got %d", len(c.Name))
	}
	return nil
}
