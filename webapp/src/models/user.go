package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"
	"webapp/src/config"
	"webapp/src/requisitions"
)

// User representa uma pessoa utilizando a rede social
type User struct {
	Id         uint64    `json:"id"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Nick       string    `json:"nick"`
	CreateDate time.Time `json:"createdate"`
	Followers  []User    `json:"followers"`
	Following  []User    `json:"following"`
	Posts      []Post    `json:"posts"`
}

// SearchCompleteUser faz todas as requisições para API para completar cadastro
func SearchCompleteUser(userId uint64, r *http.Request) (User, error) {
	chUser := make(chan User)
	chFollowers := make(chan []User)
	chFollowing := make(chan []User)
	chPosts := make(chan []Post)
	var wg sync.WaitGroup
	wg.Add(4)
	go func() {
		defer wg.Done()
		SearchUser(&chUser, userId, r)
	}()
	go func() {
		defer wg.Done()
		SearchFollowers(&chFollowers, userId, r)
	}()
	go func() {
		defer wg.Done()
		SearchFollowing(&chFollowing, userId, r)
	}()
	go func() {
		defer wg.Done()
		SearchPosts(&chPosts, userId, r)
	}()

	var user User
	user = <-*&chUser
	user.Followers = <-*&chFollowers
	user.Following = <-*&chFollowing
	user.Posts = <-*&chPosts
	wg.Wait()
	if user.Id == 0 {
		return user, errors.New("Erro ao buscar o Usuario")
	}
	if user.Followers == nil {
		return user, errors.New("Erro ao buscar o Seguidores")
	}
	if user.Following == nil {
		return user, errors.New("Erro ao buscar os que segue")
	}
	if user.Posts == nil {
		return user, errors.New("Erro ao buscar as publicações")
	}
	return user, nil
}

// SearchUser busca os dados do usuário na API
func SearchUser(chUsuario *chan User, userId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/user/%d", config.ApiUrl, userId)
	resp, erro := requisitions.RequisitionsWhithAuthentication(r, http.MethodGet, url, nil)
	if erro != nil {
		*chUsuario <- User{}
		return
	}
	defer resp.Body.Close()

	var user User
	if erro = json.NewDecoder(resp.Body).Decode(&user); erro != nil {
		*chUsuario <- User{}
		return
	}

	*chUsuario <- user
}

// SearchFollowers busca os seguidores do usuário na API
func SearchFollowers(chFollowers *chan []User, userId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/user/%d/followers", config.ApiUrl, userId)
	resp, erro := requisitions.RequisitionsWhithAuthentication(r, http.MethodGet, url, nil)
	if erro != nil {
		*chFollowers <- nil
		return
	}
	defer resp.Body.Close()

	var users []User
	if erro = json.NewDecoder(resp.Body).Decode(&users); erro != nil {
		*chFollowers <- nil
		return
	}
	if users == nil {
		*chFollowers <- make([]User, 0)
		return
	}
	*chFollowers <- users
}

// SearchFollowing busca os usuários seguidores do usuário na API
func SearchFollowing(chFollowing *chan []User, userId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/user/%d/following", config.ApiUrl, userId)

	resp, erro := requisitions.RequisitionsWhithAuthentication(r, http.MethodGet, url, nil)
	if erro != nil {
		*chFollowing <- nil
		return
	}
	defer resp.Body.Close()

	var users []User
	if erro = json.NewDecoder(resp.Body).Decode(&users); erro != nil {
		*chFollowing <- nil
		return
	}

	if users == nil {
		*chFollowing <- make([]User, 0)
		return
	}
	*chFollowing <- users
}

// SearchPosts busca os as publicações do usuário na API
func SearchPosts(chPosts *chan []Post, userId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/user/%d/post", config.ApiUrl, userId)
	resp, erro := requisitions.RequisitionsWhithAuthentication(r, http.MethodGet, url, nil)
	if erro != nil {
		*chPosts <- nil
		return
	}
	defer resp.Body.Close()

	var posts []Post
	if erro = json.NewDecoder(resp.Body).Decode(&posts); erro != nil {
		*chPosts <- nil
		return
	}

	if posts == nil {
		*chPosts <- make([]Post, 0)
		return
	}

	*chPosts <- posts
}
