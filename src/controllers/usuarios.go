package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	usuario, err := json.Marshal(map[string]string{
		"nome":  r.FormValue("nome"),
		"email": r.FormValue("email"),
		"senha": r.FormValue("senha"),
	})

	if err != nil {
		http.Error(w, "Erro ao converter dados do usuário", http.StatusBadRequest)
		return
	}

	res, err := http.Post("http://localhost:5000/usuarios", "application/json", bytes.NewBuffer(usuario))
	if err != nil {
		http.Error(w, "Erro ao enviar requisição para a API", http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	fmt.Println(res.Body)
}
