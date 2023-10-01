package datamodel

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/think-free/ABCFitness-challenge/internal/errors"
)

type BaseBooking struct {
	ClassID string    `json:"class"`
	UserID  string    `json:"user"`
	Date    time.Time `json:"date"`
}

type CreateBookingRequest struct {
	BaseBooking
}

type Booking struct {
	ID string `json:"id"`
	BaseBooking
}

type BookingFullInfo struct {
	Booking
	Class *Class `json:"class"`
	User  *User  `json:"user"`
}

func NewBooking(ctx context.Context, req *CreateBookingRequest) (*Booking, error) {
	id := uuid.New().String()

	b := &Booking{
		ID:          id,
		BaseBooking: req.BaseBooking,
	}

	if !b.isValid() {
		return nil, errors.ErrorValidationError()
	}

	return b, nil
}

func (b *Booking) isValid() bool {
	// TODO: add true validation
	if b.ClassID == "" || b.UserID == "" {
		return false
	}

	if b.Date.IsZero() {
		return false
	}

	return true
}
