package controllers

import "github.com/sunaryaagung95/go-message-api/middlewares"

// RunRouters func
func (s *Server) RunRouters() {
	// Home Router
	s.Router.HandleFunc("/", middlewares.SetJSON(s.Home)).Methods("GET")

	// Login Router
	s.Router.HandleFunc("/api/login", middlewares.SetJSON(s.Login)).Methods("POST")

	// User Router
	s.Router.HandleFunc("/api/users", middlewares.SetJSON(s.GetAllUsers)).Methods("GET")
	s.Router.HandleFunc("/api/users/{id}", middlewares.SetJSON(s.GetOneUser)).Methods("GET")
	s.Router.HandleFunc("/api/users", middlewares.SetJSON(s.AddUser)).Methods("POST")
	s.Router.HandleFunc("/api/users/{id}", middlewares.SetJSON(s.UpdateUser)).Methods("PUT")
	s.Router.HandleFunc("/api/users/{id}", middlewares.SetJSON(s.DeleteUser)).Methods("DELETE")
}
