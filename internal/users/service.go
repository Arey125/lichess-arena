package users

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/donseba/go-htmx"
)

type Service struct {
	model UserModel
	htmx  *htmx.HTMX
}

func NewService(model UserModel) Service {
	return Service{model, htmx.New()}
}

func (s Service) Register(mux *http.ServeMux) {
	mux.Handle("/", templ.Handler(home()))
}
