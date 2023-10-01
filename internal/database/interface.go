package database

import (
	"context"

	"github.com/think-free/ABCFitness-challenge/internal/cliparams"
	"github.com/think-free/ABCFitness-challenge/internal/database/memory"
	"github.com/think-free/ABCFitness-challenge/internal/datamodel"
	"github.com/think-free/ABCFitness-challenge/lib/logging"
)

const (
	DatabaseMemory = "memory"
)

type Database interface {
	SaveUser(ctx context.Context, u *datamodel.User) error
	GetUserByID(ctx context.Context, id string) (*datamodel.User, error)
	GetUserID(ctx context.Context, u *datamodel.User) (string, error)
	ListUsers(ctx context.Context, offset, count int) ([]*datamodel.User, error)

	SaveClass(ctx context.Context, cl *datamodel.Class) error
	GetClassByID(ctx context.Context, id string) (*datamodel.Class, error)
	GetClassID(ctx context.Context, cl *datamodel.Class) (string, error)
	ListClasses(ctx context.Context, offset, count int) ([]*datamodel.Class, error)

	SaveBooking(ctx context.Context, b *datamodel.Booking) error
	GetBookingID(ctx context.Context, b *datamodel.Booking) (string, error)
	GetBookingByID(ctx context.Context, id string) (*datamodel.Booking, error)
	ListBookings(ctx context.Context, offset, count int) ([]*datamodel.Booking, error)
}

func New(ctx context.Context, cp *cliparams.ClientParameters) Database {
	log := logging.Logger(ctx)

	log.Infof("initializing database '%s'", cp.DatabaseType)

	switch cp.DatabaseType {
	case DatabaseMemory:
		return memory.New(ctx)
	default:
		return nil
	}
}
