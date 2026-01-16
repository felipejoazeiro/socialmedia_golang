package requisicoes

import (
	"io"
	"net/http"
	"webapp/src/cookies"
)

func FazerReqComAuth(r *http.Request, metodo, url string, dados io.Reader) (*http.Response, error) {
	req, erro := http.NewRequest(metodo, url, dados)
	if erro != nil {
		return nil, erro
	}
	cookie, _ := cookies.Ler(r)
	req.Header.Add("Authorization", "Bearer "+cookie["token"])
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}	
	return res, nil
}
