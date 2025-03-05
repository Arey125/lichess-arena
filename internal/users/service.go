package users

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/donseba/go-htmx"
	"github.com/alexedwards/scs/v2"
)

type Service struct {
	model UserModel
	htmx  *htmx.HTMX
    session *scs.SessionManager
}

func NewService(model UserModel, sessionManager *scs.SessionManager) Service {
    return Service{model: model, htmx: htmx.New(), session: sessionManager}
}

func (s Service) Register(mux *http.ServeMux) {
	mux.HandleFunc("/", s.homePage)
	mux.Handle("GET /login", templ.Handler(login()))
	mux.Handle("GET /sign-up", templ.Handler(signUp()))
	mux.HandleFunc("POST /sign-up", s.signUpPost)
}
