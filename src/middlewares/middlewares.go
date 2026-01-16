package middlewares

import (
	"log"
	"net/http"
	"webapp/src/cookies"
)

// Logger é um middleware que registra as informações da requisição
func Logger(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		nextFunc(w, r)
	}
}

// Autenticar é um middleware que verifica se o usuário está autenticado
func Autenticar(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		valores, err := cookies.Ler(r)
		if err != nil || valores["token"] == "" {
			http.Error(w, "Não autorizado", http.StatusUnauthorized)
			return
		}
		nextFunc(w, r)
	}
}
