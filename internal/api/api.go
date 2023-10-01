package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/think-free/ABCFitness-challenge/internal/datamodel"
	"github.com/think-free/ABCFitness-challenge/internal/service"
	"github.com/think-free/ABCFitness-challenge/lib/logging"

	ierrors "github.com/think-free/ABCFitness-challenge/internal/errors"
)

type Api struct {
	srv    *service.Service
	router *mux.Router
}

func New(ctx context.Context, srv *service.Service) *Api {
	api := &Api{
		srv:    srv,
		router: mux.NewRouter(),
	}

	api.router.HandleFunc("/users", api.CreateUser).Methods("POST")
	api.router.HandleFunc("/users", api.ListUsers).Methods("GET")
	api.router.HandleFunc("/classes", api.CreateClass).Methods("POST")
	api.router.HandleFunc("/classes", api.ListClasses).Methods("GET")
	api.router.HandleFunc("/bookings", api.CreateBooking).Methods("POST")
	api.router.HandleFunc("/bookings", api.ListBookings).Methods("GET")
	api.router.HandleFunc("/booking", api.GetBooking).Methods("GET")

	return api
}

func (a *Api) Run() {
	logging.Logger(context.Background()).Info("api running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", a.router))
}

// CreateUser accept a CreateUserRequest as json in the body and returns a User as json in the data field
func (a *Api) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := logging.ContextWithLogger(r.Context())
	a.tagRequest(ctx, r)

	var req datamodel.CreateUserRequest
	err := a.decodeRequest(ctx, w, r, &req)
	if err != nil {
		return
	}

	resp, err := a.srv.CreateUser(ctx, &req)
	if err != nil {
		http.Error(w, NewErrorResponse(ctx, err).String(), a.getHttpStatusForError(ctx, err))
		return
	}

	w.WriteHeader(http.StatusCreated)
	a.writeResponse(ctx, w, NewResponse(ctx, resp))
}

// ListUsers returns a list of Users as json in the data field, it accepts offset and count as query params
func (a *Api) ListUsers(w http.ResponseWriter, r *http.Request) {
	ctx := logging.ContextWithLogger(r.Context())
	a.tagRequest(ctx, r)

	req := a.getListRequestParams(ctx, r)

	resp, err := a.srv.ListUsers(ctx, req)
	if err != nil {
		http.Error(w, NewErrorResponse(ctx, err).String(), a.getHttpStatusForError(ctx, err))
		return
	}

	a.writeResponse(ctx, w, NewResponse(ctx, resp))
}

// CreateClass accept a CreateClassRequest as json in the body and returns a Class as json in the data field
func (a *Api) CreateClass(w http.ResponseWriter, r *http.Request) {
	ctx := logging.ContextWithLogger(r.Context())
	a.tagRequest(ctx, r)

	var req datamodel.CreateClassRequest
	err := a.decodeRequest(ctx, w, r, &req)
	if err != nil {
		return
	}

	resp, err := a.srv.CreateClass(ctx, &req)
	if err != nil {
		http.Error(w, NewErrorResponse(ctx, err).String(), a.getHttpStatusForError(ctx, err))
		return
	}

	w.WriteHeader(http.StatusCreated)
	a.writeResponse(ctx, w, NewResponse(ctx, resp))
}

// ListClasses returns a list of Classes as json in the data field, it accepts offset and count as query params
func (a *Api) ListClasses(w http.ResponseWriter, r *http.Request) {
	ctx := logging.ContextWithLogger(r.Context())
	a.tagRequest(ctx, r)

	req := a.getListRequestParams(ctx, r)

	resp, err := a.srv.ListClasses(ctx, req)
	if err != nil {
		http.Error(w, NewErrorResponse(ctx, err).String(), a.getHttpStatusForError(ctx, err))
		return
	}

	a.writeResponse(ctx, w, NewResponse(ctx, resp))
}

// CreateBooking accept a CreateBookingRequest as json in the body and returns a Booking as json in the data field
func (a *Api) CreateBooking(w http.ResponseWriter, r *http.Request) {
	ctx := logging.ContextWithLogger(r.Context())
	a.tagRequest(ctx, r)

	var req datamodel.CreateBookingRequest
	err := a.decodeRequest(ctx, w, r, &req)
	if err != nil {
		return
	}

	resp, err := a.srv.CreateBooking(ctx, &req)
	if err != nil {
		http.Error(w, NewErrorResponse(ctx, err).String(), a.getHttpStatusForError(ctx, err))
		return
	}

	w.WriteHeader(http.StatusCreated)
	a.writeResponse(ctx, w, NewResponse(ctx, resp))
}

// ListBookings returns a list of Bookings as json in the data field, it accepts offset and count as query params
func (a *Api) ListBookings(w http.ResponseWriter, r *http.Request) {
	ctx := logging.ContextWithLogger(r.Context())
	a.tagRequest(ctx, r)

	req := a.getListRequestParams(ctx, r)

	resp, err := a.srv.ListBookings(ctx, req)
	if err != nil {
		http.Error(w, NewErrorResponse(ctx, err).String(), a.getHttpStatusForError(ctx, err))
		return
	}

	a.writeResponse(ctx, w, NewResponse(ctx, resp))
}

// GetBooking returns a Booking as json in the data field, it accepts id as query param
func (a *Api) GetBooking(w http.ResponseWriter, r *http.Request) {
	ctx := logging.ContextWithLogger(r.Context())
	a.tagRequest(ctx, r)

	id := r.URL.Query().Get("id")

	resp, err := a.srv.GetBooking(ctx, id)
	if err != nil {
		http.Error(w, NewErrorResponse(ctx, err).String(), a.getHttpStatusForError(ctx, err))
		return
	}

	a.writeResponse(ctx, w, NewResponse(ctx, resp))
}

// tagRequest adds tags information about the query to the logger
func (a *Api) tagRequest(ctx context.Context, r *http.Request) context.Context {
	ctx = logging.ContextWithLogger(ctx)
	log := logging.Logger(ctx)

	log.SetTag("api.remote", r.Host)
	log.SetTag("api.method", r.Method)
	log.SetTag("api.url", r.URL.String())
	log.SetTag("api.api_uuid", uuid.NewString())

	return ctx
}

// getHttpStatusForError returns the http status code for the given internal error
func (a *Api) getHttpStatusForError(ctx context.Context, err error) int {
	switch {
	case ierrors.IsAlreadyExists(err):
		return http.StatusConflict
	case ierrors.IsValidationError(err):
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError

	}
}

// decodeRequest decodes the json body of the request into the given interface, used for POST requests
func (a *Api) decodeRequest(ctx context.Context, w http.ResponseWriter, r *http.Request, v interface{}) error {
	log := logging.Logger(ctx)

	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		log.Errorf("error decoding request : %v", err)
		http.Error(w, NewErrorResponse(ctx, err).String(), http.StatusBadRequest)
		return err
	}

	return nil
}

// getListRequestParams returns the offset and count query params of the request, used for GET requests
func (a *Api) getListRequestParams(ctx context.Context, r *http.Request) *datamodel.ListRequest {
	log := logging.Logger(ctx)

	offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
	if err != nil {
		log.Debugf("error parsing offset: %v", err)
		offset = -1
	}
	count, err := strconv.Atoi(r.URL.Query().Get("count"))
	if err != nil {
		log.Debugf("error parsing count: %v", err)
		count = -1
	}

	log.SetTag("req.list.offset", offset)
	log.SetTag("req.list.count", count)

	return &datamodel.ListRequest{
		Offset: offset,
		Count:  count,
	}
}

// writeResponse encodes the given interface into the response body, used for all requests, return an error if encoding fails
func (a *Api) writeResponse(ctx context.Context, w http.ResponseWriter, v *Response) {
	log := logging.Logger(ctx)

	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		log.Errorf("error encoding response : %v", err)
		http.Error(w, NewErrorResponse(ctx, err).String(), http.StatusInternalServerError)
	}
}
