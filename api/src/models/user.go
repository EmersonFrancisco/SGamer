package models

import "time"

// User represtenta um usuario que utilizara a rede social
type User struct {
	ID         uint      `json:"id,omitempty"` // omitempty não repassa informação para json quando nulo
	Username   string    `json:"nome,omitempty"`
	Nick       string    `json:"nick,omitempty"`
	Email      string    `json:"email,omitempty"`
	Pass       string    `json:"password,omitempty"`
	CreateDate time.Time `json:"CreateDate,omitempty"`
}
