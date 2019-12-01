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
	s.Router.HandleFunc("/api/users/{id}", middlewares.SetAuth(middlewares.SetJSON(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/api/users/{id}", middlewares.SetAuth(s.DeleteUser)).Methods("DELETE")

	// Room router
	s.Router.HandleFunc("/api/rooms", middlewares.SetAuth(middlewares.SetJSON(s.GetAllRoom))).Methods("GET")
	s.Router.HandleFunc("/api/rooms/{id}", middlewares.SetAuth(middlewares.SetJSON(s.GetOneRoom))).Methods("GET")
	s.Router.HandleFunc("/api/rooms", middlewares.SetAuth(middlewares.SetJSON(s.CreateRoom))).Methods("POST")
	s.Router.HandleFunc("/api/rooms/{id}", middlewares.SetAuth(s.DeleteRoom)).Methods("DELETE")

	// Member Router
	s.Router.HandleFunc("/api/members/{id}", middlewares.SetAuth(middlewares.SetJSON(s.GetMember))).Methods("GET")
	s.Router.HandleFunc("/api/members", middlewares.SetAuth(middlewares.SetJSON(s.AddMember))).Methods("POST")

	// Message Router
	s.Router.HandleFunc("/api/messages/{id}", middlewares.SetAuth(middlewares.SetJSON(s.GetMessage))).Methods("GET")
	s.Router.HandleFunc("/api/messages", middlewares.SetAuth(middlewares.SetJSON(s.CreateMessage))).Methods("POST")
	s.Router.HandleFunc("/api/messages/{id}", middlewares.SetAuth(s.DeleteMessage)).Methods("DELETE")

}
