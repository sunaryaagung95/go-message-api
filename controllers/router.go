package controllers

import "github.com/sunaryaagung/hello-world/middlewares"

// RunRouters func
func (s *Server) RunRouters() {
	//Home Route
	s.Router.HandleFunc("/", middlewares.SetJSON(s.Home)).Methods("GET")

	//User Route
	s.Router.HandleFunc("/api/users", middlewares.SetJSON(s.GetAllUsers)).Methods("GET")
	s.Router.HandleFunc("/api/users/{id}", middlewares.SetJSON(s.GetOneUser)).Methods("GET")
	s.Router.HandleFunc("/api/users", middlewares.SetJSON(s.AddUser)).Methods("POST")
	s.Router.HandleFunc("/api/users/{id}", middlewares.SetJSON(s.UpdateUser)).Methods("PUT")
	s.Router.HandleFunc("/api/users/{id}", middlewares.SetJSON(s.DeleteUser)).Methods("DELETE")
}
