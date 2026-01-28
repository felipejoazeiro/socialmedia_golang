package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/requisicoes"
	"webapp/src/respostas"

	"github.com/gorilla/mux"
)

// CriarPublicacao cria uma nova publicação
func CriarPublicacao(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	publicacao, err := json.Marshal(map[string]string{
		"titulo":   r.FormValue("titulo"),
		"conteudo": r.FormValue("conteudo"),
	})

	if err != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroApi{Mensagem: err.Error()})
		return
	}

	url := fmt.Sprintf("%s/publicações", config.ApiUrl)

	res, err := requisicoes.FazerReqComAuth(r, http.MethodPost, url, bytes.NewBuffer(publicacao))
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

func CurtirPublicacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	publicacaoId, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroApi{Mensagem: "ID da publicação inválido"})
		return
	}
	url := fmt.Sprintf("%s/publicações/%d/curtir", config.ApiUrl, publicacaoId)

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

func DescurtirPublicacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	publicacaoId, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroApi{Mensagem: "ID da publicação inválido"})
		return
	}
	url := fmt.Sprintf("%s/publicações/%d/descurtir", config.ApiUrl, publicacaoId)

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

func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	publicacaoId, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroApi{Mensagem: "ID da publicação inválido"})
		return
	}
	r.ParseForm()
	publicacao, err := json.Marshal(map[string]string{
		"titulo":   r.FormValue("titulo"),
		"conteudo": r.FormValue("conteudo"),
	})
	if err != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroApi{Mensagem: err.Error()})
		return
	}
	url := fmt.Sprintf("%s/publicações/%d/atualizar", config.ApiUrl, publicacaoId)

	res, err := requisicoes.FazerReqComAuth(r, http.MethodPut, url, bytes.NewBuffer(publicacao))

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

func DeletarPublicacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	publicacaoId, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroApi{Mensagem: "ID da publicação inválido"})
		return
	}
	url := fmt.Sprintf("%s/publicacoes/%d", config.ApiUrl, publicacaoId)

	res, err := requisicoes.FazerReqComAuth(r, http.MethodDelete, url, nil)
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