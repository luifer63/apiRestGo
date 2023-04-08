package routers

import (
	"net/http"
	"github.com/luifer63/apiRestGo/bd"
)

//EliminarTweet -> Elimina tweet
func EliminarTweet(w http.ResponseWriter, r *http.Request){

	ID := r.URL.Query().Get("id")
	if len(ID) < 1{
		http.Error(w, "Debe enviar el ID en la URL", http.StatusBadRequest)
		return
	}

	err := bd.BorroTweet(ID, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar borrar el tweet "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)	
}