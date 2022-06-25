package models

type User struct {
	ID   int64  `json:"ID"`
	Name string `json:"name"`
	VIP  bool   `json:"VIP"`
}
