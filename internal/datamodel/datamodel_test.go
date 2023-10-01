package datamodel_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/think-free/ABCFitness-challenge/internal/datamodel"
)

func TestUser(t *testing.T) {
	ctx := context.Background()

	name := "Elon"
	surname := "Musk"
	email := "elon.musk@example.com"
	phone := "+341234567890"

	user, err := datamodel.NewUser(ctx, &datamodel.CreateUserRequest{
		BaseUser: datamodel.BaseUser{
			Name:    name,
			Surname: surname,
			Email:   email,
			Phone:   phone,
		},
	})
	require.NotNil(t, user)
	require.NoError(t, err)
	require.NotEmpty(t, user.ID)
	require.Equal(t, name, user.Name)
	require.Equal(t, surname, user.Surname)
	require.Equal(t, email, user.Email)
	require.Equal(t, phone, user.Phone)
}

func TestClass(t *testing.T) {
	ctx := context.Background()

	name := "Yoga"
	studio := "Yoga Studio"
	startDate := time.Now().AddDate(0, 0, -20)
	endDate := time.Now().AddDate(0, 0, -10)
	dailyCapacity := 10

	class, err := datamodel.NewClass(ctx, &datamodel.CreateClassRequest{
		BaseClass: datamodel.BaseClass{
			Name:          name,
			Studio:        studio,
			StartDate:     &startDate,
			EndDate:       &endDate,
			DailyCapacity: dailyCapacity,
		},
	})

	require.NotNil(t, class)
	require.NoError(t, err)
	require.NotEmpty(t, class.ID)
	require.Equal(t, name, class.Name)
	require.Equal(t, studio, class.Studio)
	require.Equal(t, startDate.Unix(), class.StartDate.Unix())
	require.Equal(t, endDate.Unix(), class.EndDate.Unix())
	require.Equal(t, dailyCapacity, class.DailyCapacity)
}

func TestBooking(t *testing.T) {
	ctx := context.Background()

	cID := "class-id"
	uID := "user-id"
	date := time.Now().AddDate(0, 0, -15)

	booking, err := datamodel.NewBooking(ctx, &datamodel.CreateBookingRequest{
		BaseBooking: datamodel.BaseBooking{
			ClassID: cID,
			UserID:  uID,
			Date:    date,
		},
	})

	require.NotNil(t, booking)
	require.NoError(t, err)
	require.NotEmpty(t, booking.ID)
	require.Equal(t, cID, booking.ClassID)
	require.Equal(t, uID, booking.UserID)
	require.Equal(t, date.Unix(), booking.Date.Unix())
}
