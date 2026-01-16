package controllers

import (
	"fmt"
	"net/http"
	"webapp/src/config"
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

	res, err := http.Get(url)
	if err != nil || res.StatusCode >= 400 {
		http.Error(w, "Não foi possível carregar a página principal", http.StatusInternalServerError)
		return
	}
	utils.ExecutarTemplate(w, "home.html", nil)
}
