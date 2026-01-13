package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/respostas"
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

	res, err := http.Post("http://localhost:5000/usuarios", "application/json", bytes.NewBuffer(usuario))
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
