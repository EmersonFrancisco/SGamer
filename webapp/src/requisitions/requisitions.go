package requisitions

import (
	"io"
	"net/http"
	"webapp/src/cookies"
)

// RequisitionsWhithAuthentication é utilizado para colocar o token na requisição
func RequisitionsWhithAuthentication(r *http.Request, method, url string, data io.Reader) (*http.Response, error) {
	request, erro := http.NewRequest(method, url, data)
	if erro != nil {
		return nil, erro
	}

	cookie, _ := cookies.Read(r)
	request.Header.Add("Authorization", "Bearer "+cookie["token"])
	client := &http.Client{}
	response, erro := client.Do(request)
	if erro != nil {
		return nil, erro
	}

	return response, nil
}
