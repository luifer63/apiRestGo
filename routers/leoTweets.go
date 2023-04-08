package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/luifer63/apiRestGo/bd"

)
// LeoTweets -> leo tweets
func LeoTweets(w http.ResponseWriter, r *http.Request){

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe enviar el ID", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviar la página", http.StatusBadRequest)
		return
	}

	pagina , err := strconv.Atoi(r.URL.Query().Get("pagina"))

	if err != nil{
		http.Error(w, "Debe enviar una pagina con un valor mayor a cero", http.StatusBadRequest)
		return
	}

	pag := int64(pagina)

	respuesta, correcto := bd.LeoTweets(ID, pag)

	if !correcto {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}