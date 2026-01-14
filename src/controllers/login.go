package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"webapp/src/respostas"
)

func FazerLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	usuario, err := json.Marshal(map[string]string{
		"email": r.FormValue("email"),
		"senha": r.FormValue("senha"),
	})
	if err != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroApi{Mensagem: "Erro ao converter usuário para JSON"})
		return
	}

	res, err := http.Post("http://api:5000/login", "application/json", bytes.NewBuffer(usuario))
	if err != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroApi{Mensagem: "Erro ao fazer requisição à API"})
		return
	}

	token, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res.StatusCode, string(token))

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		respostas.JSON(w, http.StatusUnauthorized, respostas.ErroApi{Mensagem: "Credenciais inválidas"})
		return
	}
}
