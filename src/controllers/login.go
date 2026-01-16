package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/modelos"
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

	url := fmt.Sprintf("%s/login", config.ApiUrl)
	res, err := http.Post(url, "application/json", bytes.NewBuffer(usuario))
	if err != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroApi{Mensagem: "Erro ao fazer requisição à API"})
		return
	}

	token, _ := io.ReadAll(res.Body)

	fmt.Println(res.StatusCode, string(token))

	defer res.Body.Close()

	if res.StatusCode >= 400 {
		respostas.TratarStatusCodeErro(w, res)
		return
	}

	var dadosAuth modelos.DadosAutenticacao

	if err = json.NewDecoder(res.Body).Decode(&dadosAuth); err != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroApi{Mensagem: err.Error()})
		return
	}

	if err = cookies.Salvar(w, dadosAuth.ID, dadosAuth.Token); err != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroApi{Mensagem: "Erro ao salvar cookies"})
		return
	}

	respostas.JSON(w, http.StatusOK, nil)
}
