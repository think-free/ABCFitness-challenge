package datamodel

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/think-free/ABCFitness-challenge/internal/errors"
)

type BaseClass struct {
	Studio        string     `json:"studio"`
	Name          string     `json:"class_name"`
	StartDate     *time.Time `json:"start_date"`
	EndDate       *time.Time `json:"end_date"`
	DailyCapacity int        `json:"capacity"`
}

type Class struct {
	ID string `json:"id"`
	BaseClass
}

type CreateClassRequest struct {
	BaseClass
}

func NewClass(ctx context.Context, req *CreateClassRequest) (*Class, error) {
	id := uuid.New().String()

	c := &Class{
		ID:        id,
		BaseClass: req.BaseClass,
	}

	if !c.isValid() {
		return nil, errors.ErrorValidationError()
	}

	return c, nil
}

func (c *Class) isValid() bool {
	// TODO: add true validation
	if c.Studio == "" || c.Name == "" {
		return false
	}

	if c.StartDate.After(*c.EndDate) {
		return false
	}

	if c.DailyCapacity <= 0 {
		return false
	}

	return true
}
