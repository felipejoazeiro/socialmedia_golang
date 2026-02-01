package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/requisicoes"
	"webapp/src/respostas"

	"github.com/gorilla/mux"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	usuario, err := json.Marshal(map[string]string{
		"nome":  r.FormValue("nome"),
		"email": r.FormValue("email"),
		"senha": r.FormValue("senha"),
	})

	if err != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroApi{Mensagem: err.Error()})
	}

	url := fmt.Sprintf("%s/usuarios", config.ApiUrl)
	res, err := http.Post(url, "application/json", bytes.NewBuffer(usuario))
	if err != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroApi{Mensagem: err.Error()})
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		respostas.TratarStatusCodeErro(w, res)
		return
	}

	respostas.JSON(w, res.StatusCode, nil)

	fmt.Println("Usuário criado com sucesso!")
}

func PararDeSeguirUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioId, err := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if err != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroApi{Mensagem: "ID do usuário inválido"})
		return
	}
	url := fmt.Sprintf("%s/usuarios/%d/parar-de-seguir", config.ApiUrl, usuarioId)
	res, err := requisicoes.FazerReqComAuth(r, http.MethodPost, url, nil)
	if err != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroApi{Mensagem: err.Error()})
		return
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		respostas.TratarStatusCodeErro(w, res)
		return
	}

	respostas.JSON(w, res.StatusCode, nil)

}

func SeguirUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioId, err := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if err != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroApi{Mensagem: "ID do usuário inválido"})
		return
	}
	url := fmt.Sprintf("%s/usuarios/%d/seguir", config.ApiUrl, usuarioId)
	res, err := requisicoes.FazerReqComAuth(r, http.MethodPost, url, nil)
	if err != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroApi{Mensagem: err.Error()})
		return
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		respostas.TratarStatusCodeErro(w, res)
		return
	}
	respostas.JSON(w, res.StatusCode, nil)
}

func EditarUsuario(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	usuario, err := json.Marshal(map[string]string{
		"nome":  r.FormValue("nome"),
		"email": r.FormValue("email"),
		"senha": r.FormValue("senha"),
	})

	if err != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroApi{Mensagem: err.Error()})
		return
	}

	cookie, _ := cookies.Ler(r)
	usuarioId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	url := fmt.Sprintf("%s/usuarios/%d", config.ApiUrl, usuarioId)
	res, err := requisicoes.FazerReqComAuth(r, http.MethodPut, url, bytes.NewBuffer(usuario))

	if err != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroApi{Mensagem: err.Error()})
		return
	}

	defer res.Body.Close()

	if res.StatusCode >= 400 {
		respostas.TratarStatusCodeErro(w, res)
		return
	}

	respostas.JSON(w, res.StatusCode, nil)
}
