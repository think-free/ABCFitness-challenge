package api_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/think-free/ABCFitness-challenge/internal/api"
	"github.com/think-free/ABCFitness-challenge/internal/cliparams"
	"github.com/think-free/ABCFitness-challenge/internal/database"
	"github.com/think-free/ABCFitness-challenge/internal/datamodel"
	"github.com/think-free/ABCFitness-challenge/internal/service"
)

type Message struct {
	Status string          `json:"status"`
	Data   json.RawMessage `json:"data,omitempty"`
}

func DecodeBody(body *bytes.Buffer, v interface{}) error {
	m := &Message{}
	err := json.NewDecoder(body).Decode(m)
	if err != nil {
		return err
	}

	return json.Unmarshal(m.Data, v)
}

func Test(t *testing.T) {
	ctx := context.Background()
	db := database.New(context.Background(), cliparams.New())
	srv := service.New(ctx, db)
	api := api.New(context.Background(), srv)

	// Users management
	user := &datamodel.CreateUserRequest{
		BaseUser: datamodel.BaseUser{
			Name:    "John",
			Surname: "Doe",
			Email:   "john.doe@example.com",
			Phone:   "+34123456789",
		},
	}

	u := createUser(t, api, user, false)
	assert.NotEmpty(t, u.ID)
	assert.Equal(t, user.Name, u.Name)
	assert.Equal(t, user.Surname, u.Surname)
	assert.Equal(t, user.Email, u.Email)
	assert.Equal(t, user.Phone, u.Phone)

	lstU := listUsers(t, api, &datamodel.ListRequest{
		Offset: 0,
		Count:  10,
	})
	assert.Equal(t, 1, len(lstU))
	assert.Equal(t, u.ID, lstU[0].ID)

	// Classes management
	startDate := time.Now().AddDate(0, 0, -20)
	endDate := time.Now().AddDate(0, 0, -10)
	class := &datamodel.CreateClassRequest{
		BaseClass: datamodel.BaseClass{
			Name:          "Yoga",
			Studio:        "Studio 1",
			StartDate:     &startDate,
			EndDate:       &endDate,
			DailyCapacity: 10,
		},
	}

	c := createClass(t, api, class, false)
	assert.NotEmpty(t, c.ID)
	assert.Equal(t, class.Name, c.Name)
	assert.Equal(t, class.Studio, c.Studio)
	assert.Equal(t, class.StartDate, c.StartDate)
	assert.Equal(t, class.EndDate, c.EndDate)
	assert.Equal(t, class.DailyCapacity, c.DailyCapacity)

	lstC := listClasses(t, api, &datamodel.ListRequest{
		Offset: 0,
		Count:  10,
	})
	assert.Equal(t, 1, len(lstC))
	assert.Equal(t, c.ID, lstC[0].ID)

	// Bookings management
	booking := &datamodel.CreateBookingRequest{
		BaseBooking: datamodel.BaseBooking{
			UserID:  u.ID,
			ClassID: c.ID,
			Date:    startDate.AddDate(0, 0, -15),
		},
	}

	b := createBooking(t, api, booking, false)
	assert.NotEmpty(t, b.ID)
	assert.Equal(t, booking.UserID, b.UserID)
	assert.Equal(t, booking.ClassID, b.ClassID)
	assert.Equal(t, booking.Date, b.Date)

	lstB := listBookings(t, api, &datamodel.ListRequest{
		Offset: 0,
		Count:  10,
	})
	assert.Equal(t, 1, len(lstB))
	assert.Equal(t, b.ID, lstB[0].ID)

	fb := getBooking(t, api, b.ID)
	assert.Equal(t, b.ID, fb.ID)

	// Error cases
	u = createUser(t, api, user, true)
	assert.Nil(t, u)

	c = createClass(t, api, class, true)
	assert.Nil(t, c)

	b = createBooking(t, api, booking, true)
	assert.Nil(t, b)
}

func createUser(t *testing.T, api *api.Api, user *datamodel.CreateUserRequest, shouldFail bool) *datamodel.User {
	body, err := json.Marshal(user)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/users", bytes.NewReader(body))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.CreateUser)
	handler.ServeHTTP(rr, req)
	if shouldFail {
		assert.NotEqual(t, http.StatusCreated, rr.Code)
		return nil
	}
	assert.Equal(t, http.StatusCreated, rr.Code)

	var createdUser datamodel.User
	err = DecodeBody(rr.Body, &createdUser)
	assert.NoError(t, err)

	return &createdUser
}

func listUsers(t *testing.T, api *api.Api, listUser *datamodel.ListRequest) []*datamodel.User {
	body, err := json.Marshal(listUser)
	assert.NoError(t, err)

	req, err := http.NewRequest("GET", "/users", bytes.NewReader(body))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.ListUsers)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

	var users []*datamodel.User
	err = DecodeBody(rr.Body, &users)
	assert.NoError(t, err)

	return users
}

func createClass(t *testing.T, api *api.Api, class *datamodel.CreateClassRequest, shouldFail bool) *datamodel.Class {
	body, err := json.Marshal(class)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/classes", bytes.NewReader(body))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.CreateClass)

	handler.ServeHTTP(rr, req)
	if shouldFail {
		assert.NotEqual(t, http.StatusCreated, rr.Code)
		return nil
	}
	assert.Equal(t, http.StatusCreated, rr.Code)

	var createdClass datamodel.Class
	err = DecodeBody(rr.Body, &createdClass)
	assert.NoError(t, err)

	return &createdClass
}

func listClasses(t *testing.T, api *api.Api, listClasses *datamodel.ListRequest) []*datamodel.Class {
	body, err := json.Marshal(listClasses)
	assert.NoError(t, err)

	req, err := http.NewRequest("GET", "/classes", bytes.NewReader(body))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.ListClasses)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var classes []*datamodel.Class
	err = DecodeBody(rr.Body, &classes)
	assert.NoError(t, err)

	return classes
}

func createBooking(t *testing.T, api *api.Api, booking *datamodel.CreateBookingRequest, shouldFail bool) *datamodel.Booking {
	body, err := json.Marshal(booking)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/bookings", bytes.NewReader(body))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.CreateBooking)

	handler.ServeHTTP(rr, req)
	if shouldFail {
		assert.NotEqual(t, http.StatusCreated, rr.Code)
		return nil
	}
	assert.Equal(t, http.StatusCreated, rr.Code)

	var createdBooking datamodel.Booking
	err = DecodeBody(rr.Body, &createdBooking)
	assert.NoError(t, err)

	return &createdBooking
}

func listBookings(t *testing.T, api *api.Api, listBookings *datamodel.ListRequest) []*datamodel.Booking {
	body, err := json.Marshal(listBookings)
	assert.NoError(t, err)

	req, err := http.NewRequest("GET", "/bookings", bytes.NewReader(body))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.ListBookings)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var bookings []*datamodel.Booking
	err = DecodeBody(rr.Body, &bookings)
	assert.NoError(t, err)

	return bookings
}

func getBooking(t *testing.T, api *api.Api, id string) *datamodel.BookingFullInfo {
	url := fmt.Sprintf("/booking?id=%s", id)

	req, err := http.NewRequest("GET", url, nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.GetBooking)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var bookingResp datamodel.BookingFullInfo
	err = DecodeBody(rr.Body, &bookingResp)
	assert.NoError(t, err)

	return &bookingResp
}
