package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
	if room.AdminID == 0 {
		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("admin_id is required"))
		return
	}

	roomCreated, err := room.CreateRoom(server.DB)
	if err != nil {
		formattedError := formatError(err.Error())
		responses.ERROR(w, http.StatusBadRequest, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, roomCreated.ID))
	responses.JSON(w, http.StatusOK, roomCreated)
}

//GetAllRoom func
func (server *Server) GetAllRoom(w http.ResponseWriter, r *http.Request) {
	room := models.Room{}
	rooms, err := room.GetAllRoom(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, rooms)
}

// GetOneRoom func
func (server *Server) GetOneRoom(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rid, err := strconv.Atoi(vars["id"])
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	room := models.Room{}
	roomGotten, err := room.GetOneRoom(server.DB, rid)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, roomGotten)

}

// DeleteRoom func
func (server *Server) DeleteRoom(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	room := models.Room{}
	rid, err := strconv.Atoi(vars["id"])
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	_, err = room.DeleteRoom(server.DB, rid)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", rid))
	responses.JSON(w, http.StatusNoContent, "")

}
