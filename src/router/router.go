package router

import (
	"github.com/gorilla/mux"
)

// Retorna as portas configuradas
func Gerar() *mux.Router {
	return mux.NewRouter()
}
