package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/requisitions"
	"webapp/src/response"

	"github.com/gorilla/mux"
)

// NewPost chama API para criação de uma publicalção no BD
func NewPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	post, erro := json.Marshal(map[string]string{
		"title":   r.FormValue("title"),
		"content": r.FormValue("content"),
	})
	if erro != nil {
		response.JSON(w, http.StatusBadRequest, response.ErroAPI{Erro: erro.Error()})
		return
	}
	url := fmt.Sprintf("%s/post", config.ApiUrl)
	resp, erro := requisitions.RequisitionsWhithAuthentication(r, http.MethodPost, url, bytes.NewBuffer(post))
	if erro != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErroAPI{Erro: erro.Error()})
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		response.TratamentStatusCode(w, resp)
		return
	}

	response.JSON(w, resp.StatusCode, post)

}

// LikePost chama API para curtir uma publicalção no BD
func LikePost(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	postId, erro := strconv.ParseUint(parameters["postId"], 10, 64)
	if erro != nil {
		response.JSON(w, http.StatusBadRequest, response.ErroAPI{Erro: erro.Error()})
		return
	}
	url := fmt.Sprintf("%s/post/%d/like", config.ApiUrl, postId)
	resp, erro := requisitions.RequisitionsWhithAuthentication(r, http.MethodPost, url, nil)
	if erro != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErroAPI{Erro: erro.Error()})
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		response.TratamentStatusCode(w, resp)
		return
	}

	response.JSON(w, resp.StatusCode, nil)
}

// UnlikePost chama API para descurtir uma publicalção no BD
func UnlikePost(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	postId, erro := strconv.ParseUint(parameters["postId"], 10, 64)
	if erro != nil {
		response.JSON(w, http.StatusBadRequest, response.ErroAPI{Erro: erro.Error()})
		return
	}
	url := fmt.Sprintf("%s/post/%d/unlike", config.ApiUrl, postId)
	resp, erro := requisitions.RequisitionsWhithAuthentication(r, http.MethodPost, url, nil)
	if erro != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErroAPI{Erro: erro.Error()})
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		response.TratamentStatusCode(w, resp)
		return
	}
	response.JSON(w, resp.StatusCode, nil)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	postId, erro := strconv.ParseUint(parameters["postId"], 10, 64)
	if erro != nil {
		response.JSON(w, http.StatusBadRequest, response.ErroAPI{Erro: erro.Error()})
		return
	}
	r.ParseForm()
	post, erro := json.Marshal(map[string]string{
		"title":   r.FormValue("title"),
		"content": r.FormValue("content"),
	})
	if erro != nil {
		response.JSON(w, http.StatusBadRequest, response.ErroAPI{Erro: erro.Error()})
		return
	}
	url := fmt.Sprintf("%s/post/%d", config.ApiUrl, postId)
	resp, erro := requisitions.RequisitionsWhithAuthentication(r, http.MethodPut, url, bytes.NewBuffer(post))
	if erro != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErroAPI{Erro: erro.Error()})
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		response.TratamentStatusCode(w, resp)
		return
	}
	response.JSON(w, resp.StatusCode, nil)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	postId, erro := strconv.ParseUint(parameters["postId"], 10, 64)
	if erro != nil {
		response.JSON(w, http.StatusBadRequest, response.ErroAPI{Erro: erro.Error()})
		return
	}
	url := fmt.Sprintf("%s/post/%d", config.ApiUrl, postId)
	resp, erro := requisitions.RequisitionsWhithAuthentication(r, http.MethodDelete, url, nil)
	if erro != nil {
		response.JSON(w, http.StatusInternalServerError, response.ErroAPI{Erro: erro.Error()})
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		response.TratamentStatusCode(w, resp)
		return
	}
	response.JSON(w, resp.StatusCode, nil)
}
