package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sunaryaagung95/go-message-api/models"
	"github.com/sunaryaagung95/go-message-api/responses"
)

// CreateRoom func
func (server *Server) CreateRoom(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	room := models.Room{}
	err = json.Unmarshal(body, &room)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	roomCreated, err := room.CreateRoom(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, roomCreated.ID))
	responses.JSON(w, http.StatusOK, roomCreated)
}
