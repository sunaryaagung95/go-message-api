package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// func enableCors(w *http.ResponseWriter) {
// 	(*w).Header().Set("Access-Control-Allow-Origin", "*")
// }

// JSON response
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	//enableCors(&w)
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Println(w, "%s", err.Error())
	}
}

// ERROR response
func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		JSON(w, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	JSON(w, http.StatusBadRequest, nil)
}
