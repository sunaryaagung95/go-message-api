package controllers

import (
	"net/http"

	"github.com/sunaryaagung/hello-world/responses"
)

//Home cont
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	responses.JSON(w, http.StatusOK, "This is root")
}
