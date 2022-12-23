package router

import (
	"go-postgres-crud/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/get-all-warning", controller.GetAll).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/get-warning/{period}/{day_begin}/{colour}", controller.FindWarning).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/insert", controller.InsertWarning).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/update/{period}/{day_begin}/{colour}", controller.UpdateWarning).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/delete/{period}/{day_begin}/{colour}", controller.Delete).Methods("POST", "OPTIONS")

	return router
}
