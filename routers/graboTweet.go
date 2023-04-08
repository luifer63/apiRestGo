package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/luifer63/apiRestGo/bd"
	"github.com/luifer63/apiRestGo/models"
)

// GraboTweet -> permite grabar el mensaje en bd
func GraboTweet(w http.ResponseWriter, r *http.Request){

	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GraboTweet{
		UserId: IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha: time.Now(),
	}
	var status bool
	_, status, err = bd.InsertoTweet(registro)

	if err != nil{
		http.Error(w, "Error al ingresar el tweet, reintente nuevamente" + err.Error(), 400)
		return
	}

	if !status{
		http.Error(w, "No se ha logrado insertar el Tweet" , 400)
		return
	} 

	w.WriteHeader(http.StatusCreated)
}