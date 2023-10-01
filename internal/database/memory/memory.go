package memory

import (
	"context"

	"github.com/think-free/ABCFitness-challenge/internal/datamodel"
	"github.com/think-free/ABCFitness-challenge/internal/errors"
)

// Memory implements the Database interface with a memory data collection that is not persistent
type Memory struct {
	users    []*datamodel.User
	classes  []*datamodel.Class
	bookings []*datamodel.Booking
}

func New(ctx context.Context) *Memory {
	return &Memory{}
}

func (m *Memory) SaveUser(ctx context.Context, u *datamodel.User) error {
	for _, user := range m.users {
		if user.Name == u.Name && user.Surname == u.Surname && user.Email == u.Email && user.Phone == u.Phone {
			return errors.ErrorAlreadyExists()
		}
	}

	m.users = append(m.users, u)
	return nil
}

func (m *Memory) GetUserByID(ctx context.Context, id string) (*datamodel.User, error) {
	for _, user := range m.users {
		if user.ID == id {
			return user, nil
		}
	}

	return nil, errors.ErrorNotFound()
}

func (m *Memory) GetUserID(ctx context.Context, u *datamodel.User) (string, error) {
	for _, user := range m.users {
		if user.Name == u.Name && user.Surname == u.Surname && user.Email == u.Email && user.Phone == u.Phone {
			return user.ID, nil
		}
	}

	return "", errors.ErrorNotFound()
}

func (m *Memory) ListUsers(ctx context.Context, offset, count int) ([]*datamodel.User, error) {
	return m.users, nil
}

func (m *Memory) SaveClass(ctx context.Context, cl *datamodel.Class) error {
	for _, class := range m.classes {
		if class.Studio == cl.Studio && class.Name == cl.Name && class.StartDate.Unix() == cl.StartDate.Unix() {
			return errors.ErrorAlreadyExists()
		}
	}

	m.classes = append(m.classes, cl)
	return nil
}

func (m *Memory) GetClassByID(ctx context.Context, id string) (*datamodel.Class, error) {
	for _, class := range m.classes {
		if class.ID == id {
			return class, nil
		}
	}

	return nil, errors.ErrorNotFound()
}

func (m *Memory) GetClassID(ctx context.Context, cl *datamodel.Class) (string, error) {
	for _, class := range m.classes {
		if class.Studio == cl.Studio && class.Name == cl.Name && class.StartDate.Unix() == cl.StartDate.Unix() {
			return class.ID, nil
		}
	}

	return "", errors.ErrorNotFound()
}

func (m *Memory) ListClasses(ctx context.Context, offset, count int) ([]*datamodel.Class, error) {
	return m.classes, nil
}

func (m *Memory) SaveBooking(ctx context.Context, b *datamodel.Booking) error {
	for _, booking := range m.bookings {
		if booking.UserID == b.UserID && booking.ClassID == b.ClassID && booking.Date == b.Date {
			return errors.ErrorAlreadyExists()
		}
	}

	// TODO : due to the lack of time, the booking is not checked if it is in the class date range nor if the class is full

	m.bookings = append(m.bookings, b)
	return nil
}

func (m *Memory) GetBookingID(ctx context.Context, b *datamodel.Booking) (string, error) {
	for _, booking := range m.bookings {
		if booking.UserID == b.UserID && booking.ClassID == b.ClassID && booking.Date == b.Date {
			return booking.ID, nil
		}
	}

	return "", errors.ErrorNotFound()
}

func (m *Memory) GetBookingByID(ctx context.Context, id string) (*datamodel.Booking, error) {
	for _, booking := range m.bookings {
		if booking.ID == id {
			return booking, nil
		}
	}

	return nil, errors.ErrorNotFound()
}

func (m *Memory) ListBookings(ctx context.Context, offset, count int) ([]*datamodel.Booking, error) {
	return m.bookings, nil
}
