package respostas

import (
	"encoding/json"
	"net/http"
)

type ErroApi struct {
	Mensagem string `json:"mensagem"`
}

// Retorna uma resposta em JSON para a requisição
func JSON(w http.ResponseWriter, status int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if erro := json.NewEncoder(w).Encode(dados); erro != nil {
		http.Error(w, "Erro ao converter os dados para JSON", http.StatusInternalServerError)
		return
	}
}

func TratarStatusCodeErro(w http.ResponseWriter, r *http.Response ) {
	var erroApi ErroApi
	if r.StatusCode >= 400 {
		if err := json.NewDecoder(r.Body).Decode(&erroApi); err != nil {
			http.Error(w, "Erro ao ler a resposta da API", http.StatusInternalServerError)
			return
		}
		JSON(w, r.StatusCode, erroApi)
	}
	
}