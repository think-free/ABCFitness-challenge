package service

import (
	"context"

	"github.com/think-free/ABCFitness-challenge/internal/database"
	"github.com/think-free/ABCFitness-challenge/internal/datamodel"
	"github.com/think-free/ABCFitness-challenge/lib/logging"
)

type Service struct {
	db database.Database
}

func New(ctx context.Context, db database.Database) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) CreateUser(ctx context.Context, r *datamodel.CreateUserRequest) (*datamodel.User, error) {
	log := logging.Logger(ctx)

	log.SetTag("req.user.name", r.Name)
	log.SetTag("req.user.surname", r.Surname)
	log.SetTag("req.user.email", r.Email)
	log.SetTag("req.user.phone", r.Phone)

	user, err := datamodel.NewUser(ctx, r)
	if err != nil {
		log.Errorf("error creating user : %v", err)
		return nil, err
	}

	log.SetTag("user.id", user.ID)
	log.SetTag("user.name", user.Name)
	log.SetTag("user.surname", user.Surname)
	log.SetTag("user.email", user.Email)
	log.SetTag("user.phone", user.Phone)

	err = s.db.SaveUser(ctx, user)
	if err != nil {
		uid, errID := s.db.GetUserID(ctx, user)
		if errID == nil {
			log.Errorf("user already exists with id '%s'", uid)
			return nil, err
		}
		log.Errorf("error saving user : %v", err)
		return nil, err
	}

	log.Debugf("user '%s' created", user.ID)

	return user, nil
}

func (s *Service) ListUsers(ctx context.Context, r *datamodel.ListRequest) ([]*datamodel.User, error) {
	log := logging.Logger(ctx)

	log.SetTag("req.list.user.offset", r.Offset)
	log.SetTag("req.list.user.count", r.Count)

	users, err := s.db.ListUsers(ctx, r.Offset, r.Count)
	if err != nil {
		log.Errorf("error listing users : %v", err)
		return nil, err
	}

	if len(users) == 0 {
		log.Warnf("no users found")
		return nil, nil
	}

	log.Debugf("found %d users", len(users))

	return users, nil
}

func (s *Service) CreateClass(ctx context.Context, cl *datamodel.CreateClassRequest) (*datamodel.Class, error) {
	log := logging.Logger(ctx)

	log.SetTag("req.class.studio", cl.Studio)
	log.SetTag("req.class.name", cl.Name)
	log.SetTag("req.class.date.start", cl.StartDate)
	log.SetTag("req.class.date.end", cl.EndDate)
	log.SetTag("req.class.capacity", cl.DailyCapacity)

	class, err := datamodel.NewClass(ctx, cl)
	if err != nil {
		log.Errorf("error creating class : %v", err)
		return nil, err
	}

	log.SetTag("class.id", class.ID)
	log.SetTag("class.name", class.Name)
	log.SetTag("class.studio", class.Studio)
	log.SetTag("class.date.start", class.StartDate)
	log.SetTag("class.date.end", class.EndDate)
	log.SetTag("class.capacity", class.DailyCapacity)

	err = s.db.SaveClass(ctx, class)
	if err != nil {
		cid, errID := s.db.GetClassID(ctx, class)
		if errID == nil {
			log.Errorf("class already exists with id '%s'", cid)
			return nil, err
		}
		log.Errorf("error saving class : %v", err)
		return nil, err
	}

	log.Debugf("class '%s' created", class.ID)

	return class, nil
}

func (s *Service) ListClasses(ctx context.Context, r *datamodel.ListRequest) ([]*datamodel.Class, error) {
	log := logging.Logger(ctx)

	log.SetTag("req.list.classes.offset", r.Offset)
	log.SetTag("req.list.classes.count", r.Count)

	classes, err := s.db.ListClasses(ctx, r.Offset, r.Count)
	if err != nil {
		log.Errorf("error listing classes : %v", err)
		return nil, err
	}

	if len(classes) == 0 {
		log.Warnf("no classes found")
		return nil, nil
	}

	log.Debugf("found %d classes", len(classes))

	return classes, nil
}

func (s *Service) CreateBooking(ctx context.Context, req *datamodel.CreateBookingRequest) (*datamodel.Booking, error) {
	log := logging.Logger(ctx)

	log.SetTag("req.booking.user_id", req.UserID)
	log.SetTag("req.booking.class_id", req.ClassID)
	log.SetTag("req.booking.date", req.Date)

	booking, err := datamodel.NewBooking(ctx, req)
	if err != nil {
		log.Errorf("error creating booking : %v", err)
		return nil, err
	}

	log.SetTag("booking.id", booking.ID)
	log.SetTag("booking.user_id", booking.UserID)
	log.SetTag("booking.class_id", booking.ClassID)
	log.SetTag("booking.date", booking.Date)

	err = s.db.SaveBooking(ctx, booking)
	if err != nil {
		bid, errID := s.db.GetBookingID(ctx, booking)
		if errID == nil {
			log.Errorf("booking already exists with id '%s'", bid)
			return nil, err
		}
		log.Errorf("error saving booking", booking.ID)
		return nil, err
	}

	log.Debugf("booking '%s' created", booking.ID)
	return booking, nil
}

func (s *Service) GetBooking(ctx context.Context, id string) (*datamodel.BookingFullInfo, error) {
	log := logging.Logger(ctx)

	log.SetTag("req.booking.id", id)

	booking, err := s.db.GetBookingByID(ctx, id)
	if err != nil {
		log.Errorf("error getting booking : %v", err)
		return nil, err
	}

	log.SetTag("booking.id", booking.ID)
	log.SetTag("booking.user_id", booking.UserID)
	log.SetTag("booking.class_id", booking.ClassID)
	log.SetTag("booking.date", booking.Date)

	class, err := s.db.GetClassByID(ctx, booking.ClassID)
	if err != nil {
		log.Errorf("error getting class : %v", err)
		return nil, err
	}

	log.SetTag("booking.class.id", class.ID)
	log.SetTag("booking.class.name", class.Name)
	log.SetTag("booking.class.studio", class.Studio)
	log.SetTag("booking.class.date.start", class.StartDate)
	log.SetTag("booking.class.date.end", class.EndDate)
	log.SetTag("booking.class.capacity", class.DailyCapacity)

	user, err := s.db.GetUserByID(ctx, booking.UserID)
	if err != nil {
		log.Errorf("error getting user : %v", err)
		return nil, err
	}

	log.SetTag("booking.user.id", user.ID)
	log.SetTag("booking.user.name", user.Name)
	log.SetTag("booking.user.surname", user.Surname)
	log.SetTag("booking.user.email", user.Email)
	log.SetTag("booking.user.phone", user.Phone)

	bookingFullInfo := &datamodel.BookingFullInfo{
		Booking: *booking,
		Class:   class,
		User:    user,
	}

	log.Debugf("booking '%s' found", booking.ID)

	return bookingFullInfo, nil
}

func (s *Service) ListBookings(ctx context.Context, r *datamodel.ListRequest) ([]*datamodel.Booking, error) {
	log := logging.Logger(ctx)

	log.SetTag("req.list.booking.offset", r.Offset)
	log.SetTag("req.list.booking.count", r.Count)

	bookings, err := s.db.ListBookings(ctx, r.Offset, r.Count)
	if err != nil {
		log.Errorf("error listing bookings : %v", err)
		return nil, err
	}

	if len(bookings) == 0 {
		log.Warnf("no bookings found")
		return nil, nil
	}

	log.Debugf("found %d bookings", len(bookings))

	return bookings, nil
}
