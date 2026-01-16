package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/router"
	"webapp/src/utils"
)

/* 
Gerando chave estrangeira, rodar apenas uma vez e copiar o valor gerado para o .env
func init() {
	hashKey := hex.EncodeToString(securecookie.GenerateRandomKey(16))
	fmt.Println(hashKey)
} */

func main() {
	config.Carregar()
	cookies.Configurar()
	utils.CarregarTemplates()
	r := router.Gerar()

	fmt.Println("Servidor rodando na porta", config.Porta)
	log.Fatal(http.ListenAndServe(config.Porta, r))
}
