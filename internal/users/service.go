package users

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/donseba/go-htmx"
	"github.com/gorilla/sessions"
)

type Service struct {
	model UserModel
	htmx  *htmx.HTMX
    sessionStore *sessions.CookieStore
}

func NewService(model UserModel, sessionStore *sessions.CookieStore) Service {
	return Service{model, htmx.New(), sessionStore}
}

func (s Service) Register(mux *http.ServeMux) {
	mux.HandleFunc("/", s.homePage)
	mux.Handle("GET /login", templ.Handler(login()))
	mux.Handle("GET /sign-up", templ.Handler(signUp()))
	mux.HandleFunc("POST /sign-up", s.signUpPost)
}
