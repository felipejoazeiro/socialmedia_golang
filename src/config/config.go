package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	ApiUrl   = "http://api:5000"
	Porta    = ":3000"
	HashKey  []byte //Autenticar o Cooki
	BlockKey []byte //Criptografar o Cookie
)

// Carregar inicializa as variáveis de ambiente
func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	ApiUrl = os.Getenv("API_URL")
}
