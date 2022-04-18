package models

import "time"

// Post representa os dados de uma publicação
type Post struct {
	Id         uint64    `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	AuthorId   uint64    `json:"authorid,omitempty"`
	AuthorNick string    `json:"authornick,omitempty"`
	Likes      uint64    `json:"likes"`
	DatePost   time.Time `json:"datepost,omitempty"`
}
