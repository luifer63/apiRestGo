package routers

import(
	"net/http"
	"github.com/luifer63/apiRestGo/bd"
	"github.com/luifer63/apiRestGo/models"
)

// BajaRelacion -> Dar de baja una relacion entre usuarios
func BajaRelacion(w http.ResponseWriter, r *http.Request){
	ID := r.URL.Query().Get("id")
	if len(ID) < 1{
		http.Error(w, "El parámetro es obligatorio", http.StatusBadRequest)
		return
	}

	var t models.Relacion

	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.BorroRelacion(t)

	if err != nil{
		http.Error(w, "Ocurrió un error al Borrar relacion "+err.Error(), http.StatusBadRequest)
		return
	}

	if !status{
		http.Error(w, "No se ha logrado Borrar la relacion ", http.StatusBadRequest)
		return
	} 

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}