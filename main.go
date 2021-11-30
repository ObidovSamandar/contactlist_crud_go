package main

import (
	"github.com/gin-gonic/gin"
	"github.com/obidovsamandar/task-list-crud/controllers"
	"github.com/obidovsamandar/task-list-crud/dbconnector"
)

func main() {

	dbconnector.DBClientConnector()

	route := gin.Default()

	route.POST("/createcontact", controllers.CreateContact)
	route.DELETE("/deletecontact/:id", controllers.DeleteContact)
	route.GET("/allcontacts", controllers.GetAllContacts)
	route.GET("/contact/:id", controllers.GetSpecificContact)
	route.PUT("/contact/:id", controllers.UpdateContact)

	route.Run(":8081")
}
