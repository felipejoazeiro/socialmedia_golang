package modelos

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"webapp/src/config"
	"webapp/src/requisicoes"
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

func BuscarDadosDoUsuario(canal chan<- Usuario, usuarioId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios/%d", config.ApiUrl, usuarioId)
	res, err := requisicoes.FazerReqComAuth(r, http.MethodGet, url, nil)
	if err != nil {
		canal <- Usuario{}
		return
	}
	defer res.Body.Close()
	
	var usuario Usuario
	if err = json.NewDecoder(res.Body).Decode(&usuario); err != nil {
		canal <- Usuario{}
		return
	}

	canal <- usuario

}

func BuscarSeguidores(canal chan<- []Usuario, usuarioId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios/%d/seguidores", config.ApiUrl, usuarioId)
	res, err := requisicoes.FazerReqComAuth(r, http.MethodGet, url, nil)
	if err != nil {
		canal <- []Usuario{}
		return
	}
	defer res.Body.Close()
	
	var seguidores []Usuario
	if err = json.NewDecoder(res.Body).Decode(&seguidores); err != nil {
		canal <- []Usuario{}
		return
	}
	canal <- seguidores
}

func BuscarSeguindo(canal chan<- []Usuario, usuarioId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios/%d/seguindo", config.ApiUrl, usuarioId)
	res, err := requisicoes.FazerReqComAuth(r, http.MethodGet, url, nil)
	if err != nil {
		canal <- []Usuario{}
		return
	}
	defer res.Body.Close()
	
	var seguindo []Usuario
	if err = json.NewDecoder(res.Body).Decode(&seguindo); err != nil {
		canal <- []Usuario{}
		return
	}
	canal <- seguindo
}

func BuscarPublicacoesDoUsuario(canal chan<- []Publicacao, usuarioId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuarios/%d/publicacoes", config.ApiUrl, usuarioId)
	res, err := requisicoes.FazerReqComAuth(r, http.MethodGet, url, nil)
	if err != nil {
		canal <- []Publicacao{}
		return
	}
	defer res.Body.Close()
	var publicacoes []Publicacao
	if err = json.NewDecoder(res.Body).Decode(&publicacoes); err != nil {
		canal <- []Publicacao{}
		return
	}
	canal <- publicacoes
}
