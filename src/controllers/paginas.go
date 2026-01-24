package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/modelos"
	"webapp/src/requisicoes"
	"webapp/src/respostas"
	"webapp/src/utils"
)

// Renderiza a tela de login
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
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
