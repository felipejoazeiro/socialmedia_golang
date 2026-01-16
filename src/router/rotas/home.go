package rotas

import "webapp/src/controllers"

var rotaPaginaPrincipal = Rota{
	URI:                "/",
	Metodo:             "GET",
	Funcao:             controllers.CarregarPaginaPrincipal,
	RequerAutenticacao: true,
}
