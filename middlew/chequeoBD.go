package middlew

import (
	"net/http"
	"github.com/luifer63/apiRestGo/bd"
)


func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){

		if bd.ChequeoConexion() == 0{
			http.Error(w, "Conexión Perdida con BD", 500)
			return
		}
		next.ServeHTTP(w, r)
	}	
}