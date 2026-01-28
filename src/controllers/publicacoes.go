package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/modelos"
	"webapp/src/requisicoes"
	"webapp/src/respostas"
	"webapp/src/utils"

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

func CarregarPaginaDeEdicaoDePublicacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroApi{Mensagem: "ID da publicação inválido"})
		return
	}

	url := fmt.Sprintf("%s/publicações/%d", config.ApiUrl, publicacaoID)

	res, err := requisicoes.FazerReqComAuth(r, http.MethodGet, url, nil)
	if err != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroApi{Mensagem: err.Error()})
		return
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		respostas.TratarStatusCodeErro(w, res)
		return
	}
	var publicacao modelos.Publicacao
	if err := json.NewDecoder(res.Body).Decode(&publicacao); err != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroApi{Mensagem: "Erro ao ler a publicação"})
		return
	}
	utils.ExecutarTemplate(w, "editar-publicacao.html", publicacao)
}

func EditarPublicacao(w http.ResponseWriter, r *http.Request) {
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
	url := fmt.Sprintf("%s/publicações/%d/editar", config.ApiUrl, publicacaoId)

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