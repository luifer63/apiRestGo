package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/luifer63/apiRestGo/middlew"

	"github.com/luifer63/apiRestGo/routers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Manejadores() -> configuro handler y pongo a escuchar el servidor
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
