package routers

import(
	"encoding/json"
	"net/http"
	"github.com/luifer63/apiRestGo/bd"
	"github.com/luifer63/apiRestGo/models"
)

// ConsultaRelacion -> Consiltar la relacion entre usuarios
func ConsultaRelacion(w http.ResponseWriter, r *http.Request){

	ID := r.URL.Query().Get("id")
	if len(ID) < 1{
		http.Error(w, "El parámetro es obligatorio", http.StatusBadRequest)
		return
	}

	var t models.Relacion

	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	var resp models.RespuestaConsultaRelacion

	status, err := bd.ConsultoRelacion(t)

	if err != nil || !status{
		resp.Status = false
	}else{
		resp.Status = true
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}