package routers

import(
	"encoding/json"
	"net/http"
	"github.com/luifer63/apiRestGo/bd"
	"github.com/luifer63/apiRestGo/models"
)


// Regisrtro() -> Registro de usuarios

func Registro(w http.ResponseWriter, r *http.Request){

	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil{
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0{
		http.Error(w, "Email es requerido "+err.Error(), 400)
		return
	}

	if len(t.Password) < 6{
		http.Error(w, "Debe especificar una contraseña de al menos 6 carárteres "+err.Error(), 400)
		return
	}

	_,encontrado,_ := bd.ChequeoYaExisteUsuario(t.Email)

	if encontrado {
		http.Error(w, "Ya existe un usuario con ese Email"+err.Error(), 400)
		return
	}

	_,status,err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar registrar el usuario "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado insertar el registro del usuario"+err.Error(), 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}