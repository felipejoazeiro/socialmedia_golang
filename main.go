package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
)

func main() {
	fmt.Println("Rodando webapp")

	r := router.Gerar()
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":3000", nil))
}
