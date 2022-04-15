package models

import (
	"errors"
	"strings"
	"time"
)

// Post representa a estrutura de uma publicação do usuário
type Post struct {
	Id         uint64    `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	AuthorId   uint64    `json:"authorid,omitempty"`
	AuthorNick string    `json:"authornick,omitempty"`
	Likes      uint64    `json:"likes"`
	DatePost   time.Time `json:"datepost,omitempty"`
}

// Prepare verificar se a publicação está de acordo com as regras
func (post *Post) Prepare() error {

	if erro := post.validated(); erro != nil {
		return erro
	}
	post.format()
	return nil
}

func (post *Post) validated() error {
	switch {
	case post.Title == "":
		return errors.New("O titulo é obrigatorio, e não pode estar em branco")
	case post.Content == "":
		return errors.New("O Conteudo é obrigatorio e não pode estar em branco")
	}
	return nil

}

func (post *Post) format() {
	post.Title = strings.TrimSpace(post.Title)
	post.Content = strings.TrimSpace(post.Content)
}
