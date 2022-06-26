package models

import (
	"context"
	"encoding/json"
	"net/http"
)

type User struct {
	ID   int64  `json:"ID"`
	Name string `json:"name"`
	VIP  bool   `json:"VIP"`
}

type UserRequest struct {
	Name string `json:"name"`
	VIP  bool   `json:"vip"`
}

type OperationMessage struct {
	Message string `json:"message"`
}

type UserResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"`
}

type FindUserRequest struct {
	ID int64 `json:"ID"`
}

type UpdateUserRequest struct {
	User
}

func DecodeUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request UserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeUpdateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeFindUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request FindUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
