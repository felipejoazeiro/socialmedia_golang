package rotas

import "webapp/src/controllers"

var rotasPublicacoes = []Rota{
	{
		URI:                "/publicacoes",
		Metodo:             "POST",
		Funcao:             controllers.CriarPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publicacoes/{publicacaoId}/curtir",
		Metodo:             "POST",
		Funcao:             controllers.CurtirPublicacao,
		RequerAutenticacao: true,
	},
	{
		URI:                "/publicacoes/{publicacaoId}/descurtir",
		Metodo:             "POST",
		Funcao:             controllers.DescurtirPublicacao,
		RequerAutenticacao: true,
	},
}
