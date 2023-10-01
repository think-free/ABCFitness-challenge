package datamodel

type ListRequest struct {
	Offset int `json:"offset"`
	Count  int `json:"count"`
}
