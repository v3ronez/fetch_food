package pkg

import (
	"errors"
	"io"
	"net/http"
)

type HttpService interface {
	Get(url string) ([]byte, error)
}

type HttpClientService struct {
	client *http.Client
}

type HttpStatus struct {
	code int
	msg  string
}

func NewHttpClient() *HttpClientService {
	return &HttpClientService{
		client: &http.Client{},
	}
}

func (h *HttpClientService) Get(url string) ([]byte, error) {
	r, err := h.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	if isValid := h.CheckStatus(r); isValid.code != 200 {
		return nil, errors.New(isValid.msg)
	}
	b, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	return b, err
}

func (h *HttpClientService) CheckStatus(response *http.Response) HttpStatus {
	return HttpStatus{
		code: response.StatusCode,
		msg:  response.Status,
	}
}
