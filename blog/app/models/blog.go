package models

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Article struct {
	ID    int64  `json:"ID"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type User struct {
	ID   int64  `json:"ID"`
	Name string `json:"name"`
	VIP  bool   `json:"VIP"`
}

type ArticleRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type OperationMessage struct {
	Message string `json:"message"`
}

type FindByIDRequest struct {
	ID int64 `json:"ID"`
}

type UserIDRequest struct {
	UserID int64 `json:"userID"`
}

type UpdateArticleRequest struct {
	Article
}

func DecodeArticleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request ArticleRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeUpdateArticleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request UpdateArticleRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func DecodeFindArticleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request FindByIDRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		fmt.Println(err.Error() + "xxxxxxxx")
		return nil, err
	}
	return request, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
