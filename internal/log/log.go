package log

import (
	"fmt"
	"net/http"
)

func LogRequests(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Printf("%s - %s %s\n", r.RemoteAddr, r.Method, r.URL.RequestURI())
        next.ServeHTTP(w,r)
    })
}
