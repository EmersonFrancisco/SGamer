package models

import "time"

// User represtenta um usuario que utilizara a rede social
type User struct {
	ID         uint64    `json:"id,omitempty"` // omitempty não repassa informação para json quando nulo
	Username   string    `json:"username,omitempty"`
	Nick       string    `json:"nick,omitempty"`
	Email      string    `json:"email,omitempty"`
	Pass       string    `json:"pass,omitempty"`
	CreateDate time.Time `json:"CreateDate,omitempty"`
}
