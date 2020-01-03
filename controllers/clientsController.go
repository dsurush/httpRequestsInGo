package controllers

import (
	"encoding/json"
	"market/db"
	"market/models"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

//
func getID(w http.ResponseWriter, ps httprouter.Params) (int, bool) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.WriteHeader(400)
		return 0, false
	}
	return id, true
}

// DeleteClientHandler "Удаляет клиента по ID"
func DeleteClientHandler(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {
	_, ok := getID(w, pr)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	//json.NewDecoder(r.Body).Decode(&requestBody)
	id := pr.ByName("id")
	err := db.DeleteClient(id)
	if err != nil {
		// w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	return

}

//GetClientsHandler "Вывод всех клиентов на сервере сортированный по  имени"
func GetClientsHandler(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {
	res, err := db.GetClients()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(res)
	return
}

// CreateNewClientHandler "для роута "/products" "
func CreateNewClientHandler(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {
	var requesteBody models.CreateNewClient

	json.NewDecoder(r.Body).Decode(&requesteBody)

	res, err := db.AddClient(requesteBody)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(res)
	return
}
