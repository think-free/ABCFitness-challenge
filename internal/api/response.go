package api

import (
	"context"
	"encoding/json"
	"time"

	"github.com/think-free/ABCFitness-challenge/lib/logging"
)

const (
	StatusOK    = "ok"
	StatusError = "error"
)

// Response object of the api, used to return data to the client in the data field
// There is two types of responses, one for errors and one for correct data
type Response struct {
	Status   string          `json:"status"`
	Data     json.RawMessage `json:"data,omitempty"`
	Metadata Metadata        `json:"metadata"`
}

type Metadata struct {
	CreatedAt string `json:"createdAt,omitempty"`
}

func NewResponse(ctx context.Context, data interface{}) *Response {
	dt, err := json.Marshal(data)
	if err != nil {
		return NewErrorResponse(ctx, err)
	}
	r := &Response{
		Status: StatusOK,
		Data:   dt,
		Metadata: Metadata{
			CreatedAt: time.Now().Format(time.RFC3339),
		},
	}
	return r
}

func NewErrorResponse(ctx context.Context, err error) *Response {
	dt, err := json.Marshal(err.Error())
	if err != nil {
		logging.Logger(ctx).Errorf("error marshaling error while build error response: %v", err)
		return NewErrorResponse(ctx, err)
	}
	r := &Response{
		Status: StatusError,
		Data:   dt,
		Metadata: Metadata{
			CreatedAt: time.Now().Format(time.RFC3339),
		},
	}
	return r
}

func (r *Response) String() string {
	js, _ := json.Marshal(r)
	return string(js)
}
