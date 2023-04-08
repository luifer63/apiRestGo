package routers

import (
	"encoding/json"
	"net/http"

	"github.com/luifer63/apiRestGo/bd"
	"github.com/luifer63/apiRestGo/models"

)

// ModificarPerfil -> modifica el perfil del usuario
func ModificarPerfil(w http.ResponseWriter, r *http.Request){

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Datos incorrectos " + err.Error(), 400)
		return
	}

	var status bool

	status, err = bd.ModificoRegistro(t, IDUsuario)

	if err != nil {
		http.Error(w, "Ocurrió un error al modificar el registro, intente nuevamente " + err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se ha modificado el regisrtro", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}