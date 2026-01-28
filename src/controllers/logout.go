package controllers

import (
	"net/http"
	"webapp/src/cookies"
)

// FazerLogout realiza o logout do usuário
func FazerLogout(w http.ResponseWriter, r *http.Request) {
	cookies.Deletar(w)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
