package controller

import (
	"filestoreapp/service"
	"net/http"

	"github.com/gorilla/mux"
)

// All the Handler of file store and heath api
func FileAppRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/store", service.StoreFileHandler).Methods(http.MethodPost)
	r.HandleFunc("/store", service.ListFilesHandler).Methods(http.MethodGet)
	r.HandleFunc("/store", service.UpdateFileHandler).Methods(http.MethodPatch)
	r.HandleFunc("/store", service.DeleteFileHandler).Methods(http.MethodDelete)
	r.HandleFunc("/store/wordcount", service.FindWordCountHandler).Methods(http.MethodGet)
	r.HandleFunc("/readiness", service.CheckReadyHandler).Methods(http.MethodGet)
	r.HandleFunc("/health", service.CheckHealthHandler).Methods(http.MethodGet)

	return r
}
