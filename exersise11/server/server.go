package server

import (
	"log"
	"net/http"
	"time"
	"fmt"
)

type Server struct {
	// config Config
	// store  Store
	router *http.ServeMux
}



func (s *Server)helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		end := time.Now()
		latency := end.Sub(start)
		clientIP := r.RemoteAddr
		method := r.Method
		uri := r.RequestURI
		log.Printf(
			"%s - %s %s %s - %v",
			clientIP, method, uri, r.Proto, latency,
		)
	})
}

func (s *Server) Start() error{
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.helloHandler)

	// Используем LoggerMiddleware для логирования запросов.
	loggedMux := LoggerMiddleware(mux)

	server := &http.Server{
		Addr:    ":8080",
		Handler: loggedMux,
	}

	fmt.Println("Server started on port 8080")
	log.Fatal(server.ListenAndServe())
	return nil
}