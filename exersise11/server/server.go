package server

import (
	"calandary/helpers"
	"calandary/store"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Server struct {
	Addr string
	Store  *store.Store
	router *http.ServeMux
	Handler http.Handler
}

func IntServer(db *store.Store) *Server {
	return &Server{
		Addr:    ":8080",
		Store: db,
		router: http.NewServeMux(),
	}
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

func (s *Server)CreateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseForm()
	if !helpers.CheckError(err) {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}
	user_id := r.PostForm.Get("user_id")
	if user_id == "" {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}
	date, err := time.Parse("2006-01-02",  r.PostForm.Get("date"))
	if !helpers.CheckError(err) {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}
	title :=  r.PostForm.Get("title")
	if title == "" {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}
	resp, err := s.Store.CreateNew(user_id, title, date)
	if !helpers.CheckError(err) {
		http.Error(w, "Error creating event", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func (s *Server)UpdateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseForm()
	if !helpers.CheckError(err) {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}
	id := r.PostForm.Get("id")
	if id == "" {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}
	user_id := r.PostForm.Get("user_id")
	if user_id == "" {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}
	date, err := time.Parse("2006-01-02",  r.PostForm.Get("date"))
	if !helpers.CheckError(err) {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}
	title :=  r.PostForm.Get("title")
	if title == "" {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}
	resp, err := s.Store.UpdateEvent(id, user_id, title, date)
	if !helpers.CheckError(err) {
		http.Error(w, "Error creating event", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func (s *Server)DeleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	err := r.ParseForm()
	if !helpers.CheckError(err) {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}
	id := r.PostForm.Get("id")
	if id == "" {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}
	user_id := r.PostForm.Get("user_id")
	if user_id == "" {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}
	resp, err := s.Store.DeleteEvent(id, user_id)
	if !helpers.CheckError(err) {
		http.Error(w, "Error creating event", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func (s *Server)EventsForDay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	query := r.URL.Query()

	user_id := query.Get("user_id")
	date_str := query.Get("date")
	fmt.Println(user_id, date_str)
	if user_id == "" || date_str == "" {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}
	date, err := time.Parse("2006-01-02",  date_str)
	if !helpers.CheckError(err) {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}
	resp, err := s.Store.GetEventsByDay(user_id, date)	
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func (s *Server)EventsForWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	query := r.URL.Query()

	user_id := query.Get("user_id")
	date_str := query.Get("date")
	fmt.Println(user_id, date_str)
	if user_id == "" || date_str == "" {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}
	date, err := time.Parse("2006-01-02",  date_str)
	if !helpers.CheckError(err) {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}
	resp, err := s.Store.GetEventsByWeek(user_id, date)	
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func (s *Server)EventsForMonth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	query := r.URL.Query()

	user_id := query.Get("user_id")
	date_str := query.Get("date")
	fmt.Println(user_id, date_str)
	if user_id == "" || date_str == "" {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}
	date, err := time.Parse("2006-01-02",  date_str)
	if !helpers.CheckError(err) {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}
	resp, err := s.Store.GetEventsByMonth(user_id, date)	
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}
func (s *Server) Start() {
	s.router.HandleFunc("/create_event", s.CreateEvent)
	s.router.HandleFunc("/update_event", s.UpdateEvent)
	s.router.HandleFunc("/delete_event", s.DeleteEvent)
	s.router.HandleFunc("/events_for_day/", s.EventsForDay)
	s.router.HandleFunc("/events_for_week/", s.EventsForWeek)
	s.router.HandleFunc("/events_for_month/", s.EventsForMonth)

	// Используем LoggerMiddleware для логирования запросов.
	loggedMux := LoggerMiddleware(s.router)

	// server := &http.Server{
	// 	Addr:    ":8080",
	// 	Handler: loggedMux,
	// }
	s.Handler = loggedMux
	fmt.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(s.Addr, s.Handler))
	
}