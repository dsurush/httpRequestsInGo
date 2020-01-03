package controllers

import (
	"encoding/json"
	"market/db"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GetProductsHandler "для роута "/products" "
func GetProductsHandler(w http.ResponseWriter, r *http.Request, pr httprouter.Params) {
	res, err := db.GetProducts()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(res)
	return
}
