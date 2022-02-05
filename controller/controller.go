package controller

import (
	"filestoreapp/service"
	"net/http"

	"github.com/gorilla/mux"
)

func FileAppRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/store", service.StoreFile).Methods(http.MethodPost)
	r.HandleFunc("/store", service.ListFiles).Methods(http.MethodGet)
	r.HandleFunc("/store", service.UpdateFile).Methods(http.MethodPatch)
	r.HandleFunc("/store", service.DeleteFile).Methods(http.MethodDelete)

	return r
}
