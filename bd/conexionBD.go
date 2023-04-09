package bd

import(
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoCN -> es el objeto de conexion a BD
var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://xxxxxx:xxxxxxxxxxxx@cluster0.igqhpqx.mongodb.net/test")

// ConectarBD() -> funcion para conectar a base de datos
func ConectarBD() *mongo.Client{
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil{
		log.Fatal(err.Error())
		return client
	} 

	log.Println("Conexión exitosa de la BD")
	return client
}

// ChequeoConexion() es el ping a la BD
func ChequeoConexion() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil{
		return 0
	}
	return 1
}
