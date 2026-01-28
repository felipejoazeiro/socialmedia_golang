package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/modelos"
	"webapp/src/requisicoes"
	"webapp/src/respostas"
	"webapp/src/utils"

	"github.com/gorilla/mux"
)

// Renderiza a tela de login
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Ler(r)

	if cookie["token"] != "" {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}
	utils.ExecutarTemplate(w, "login.html", nil)
}

// Renderiza a tela de cadastro de usuário
func CarregarPaginaDeCadastroDeUsuario(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "cadastro.html", nil)
}

// Renderiza a página principal após o login
func CarregarPaginaPrincipal(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/publicações", config.ApiUrl)

	res, err := requisicoes.FazerReqComAuth(r, http.MethodGet, url, nil)

	if res.StatusCode >= 400 {
		respostas.TratarStatusCodeErro(w, res)
		return
	}

	var publi []modelos.Publicacao
	if err = json.NewDecoder(res.Body).Decode(&publi); err != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroApi{Mensagem: err.Error()})
		return
	}

	fmt.Println(res.StatusCode, err)

	cookie, _ := cookies.Ler(r)
	usuarioId, _ := strconv.ParseUint(cookie["id"], 10, 64)

	utils.ExecutarTemplate(w, "home.html", struct {
		Publicacoes []modelos.Publicacao
		UsuarioID   uint64
	}{
		Publicacoes: publi,
		UsuarioID:   usuarioId,
	})
}

// Renderiza a página de edição de publicação
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

func CarregarPaginaDeUsuarios(w http.ResponseWriter, r *http.Request) {
	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))
	url := fmt.Sprintf("%s/usuarios?usuario=%s", config.ApiUrl, nomeOuNick)

	
}
