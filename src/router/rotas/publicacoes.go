package rotas

import "webapp/src/controllers"

var rotasPublicacoes = []Rota{
	{
		URI:                "/publicacoes",
		Metodo:             "POST",
		Funcao:             controllers.CriarPublicacao,
		RequerAutenticacao: true,
	},
}
