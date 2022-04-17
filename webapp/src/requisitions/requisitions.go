package requisitions

import (
	"io"
	"net/http"
)

// RequisitionsWhithAuthentication é utilizado para colocar o toke na requisição
func RequisitionsWhithAuthentication(r *http.Request, method, url string, datas io.Reader) (*http.Response, error) {
	request, erro := http.NewRequest(method, url, datas)
	if erro != nil {
		return nil, erro
	}

	return request.Response, nil
}
