package modelos

import (
	"net/http"
	"time"
)

type Usuario struct {
	ID          uint64       `json:"id"`
	Nome        string       `json:"nome"`
	Email       string       `json:"email"`
	Nick        string       `json:"nick"`
	CriadoEm    time.Time    `json:"criadoEm"`
	Seguidores  []Usuario    `json:"seguidores,omitempty"`
	Seguindo    []Usuario    `json:"seguindo,omitempty"`
	Publicacoes []Publicacao `json:"publicacoes,omitempty"`
}

func BuscarUsuarioCompleto(usuarioID uint64, r *http.Request) (Usuario, error) {
	canalUsuario := make(chan Usuario)
	canalSeguidores := make(chan []Usuario)
	canalSeguindo := make(chan []Usuario)
	canalPublicacoes := make(chan []Publicacao)

	go BuscarDadosDoUsuario(canalUsuario, usuarioID, r)
	go BuscarSeguidores(canalSeguidores, usuarioID, r)
	go BuscarSeguindo(canalSeguindo, usuarioID, r)
	go BuscarPublicacoesDoUsuario(canalPublicacoes, usuarioID, r)

	usuario := <-canalUsuario
	usuario.Seguidores = <-canalSeguidores
	usuario.Seguindo = <-canalSeguindo
	usuario.Publicacoes = <-canalPublicacoes
	return usuario, nil
}


func BuscarDadosDoUsuario(canal <-chan Usuario, usuarioId uint64, r *http.Request) {

}

func BuscarSeguidores(canal <-chan []Usuario, usuarioId uint64, r *http.Request) {
}

func BuscarSeguindo(canal <-chan []Usuario, usuarioId uint64, r *http.Request) {
}

func BuscarPublicacoesDoUsuario(canal <-chan []Publicacao, usuarioId uint64, r *http.Request) {
}
