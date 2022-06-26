package models

import (
	"context"
	"encoding/json"
	"net/http"
)

type Article struct {
	ID    int64  `json:"ID"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type ArticleRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type OperationMessage struct {
	Message string `json:"message"`
}

type FindArticleResponse struct {
	ID int64 `json:"ID"`
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
	var request FindArticleResponse
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
