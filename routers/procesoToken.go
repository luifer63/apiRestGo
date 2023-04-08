package routers

import(
	"errors"
	"strings"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/luifer63/apiRestGo/bd"
	"github.com/luifer63/apiRestGo/models"
)

// Emaul -> valor usado en todos los endpoints
var Email string
// IDUsuario -> valor usado en todos los endpoints
var IDUsuario string

///ProcesoToken -> proceso token para extraer sus valores
func ProcesoToken(tk string)(*models.Claim, bool, string, error){

	miClave := []byte("MastersDelDesarrollo_GrupoDeFacebook")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")		
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if err == nil{
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}
	if !tkn.Valid{
		return claims, false, string(""), errors.New("token invalido")
	}
	return claims, false, string(""), err
}