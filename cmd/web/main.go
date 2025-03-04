package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"

	database "lichess-arena/internal/db"
	"lichess-arena/internal/users"

	"github.com/gorilla/sessions"
)

func main() {
	dsn := os.Getenv("DB")
	if dsn == "" {
		panic("DB environment variable is not set")
	}

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		panic(err)
	}
	if port == 0 {
		panic("PORT environment variable is not set")
	}

    storageSecret := os.Getenv("STORAGE_SECRET")
    if storageSecret == "" {
        panic("STORAGE_SECRET environment variable is not set")
    }

	db := database.Connect(dsn)
    sessionStore := sessions.NewCookieStore([]byte(storageSecret))

	mux := http.NewServeMux()
	userModel := users.NewModel(db)

	staticFileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("GET /static/", http.StripPrefix("/static", staticFileServer))

	userService := users.NewService(userModel, sessionStore)
	userService.Register(mux)

	server := http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	fmt.Printf("Listening on http://127.0.0.1:%d\n", port)
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
