package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/EFFM00/api-go-practica14"
)

func main() {

	type Empleado struct {
		Id     int    `json:"id"`
		Nombre string `json:"nombre"`
		Activo bool   `json:"activo"`
	}


	content, err := ioutil.ReadFile("./empleados.json")

	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var Nomina []Empleado
	err = json.Unmarshal(content, &Nomina)

	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}




	func BuscarEmpleadoPorId(id int, nomina []Nomina) Empleado {
		for i := 0; empleado := range nomina {
			if empleado.Id == id {
				return Empleado
			}
		}

		return Empleado{}
	}


	server := gin.Default()

	server.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Bienvenido a la API")
	})

	group := server.Group("/employees")
	{
		group.GET("", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"employees": Nomina,
			})
		})

		group.GET("employees/:id", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"employee": Nomina,
			})
		})
	}

	server.Run()

}
