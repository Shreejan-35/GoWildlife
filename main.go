package main

import (
	"go-wildlife/models"
	"go-wildlife/routes"

	"log"

	"github.com/gin-gonic/gin"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	r := gin.Default() //router

	models.ConnectDb()
	app := r.Group("/api")

	{
		app.GET("animals", routes.GetAnimals)
		app.GET("animal/:id", routes.GetAnimalById)
		app.POST("person", routes.AddAnimal)
		app.PUT("person/:id", routes.UpdateAnimal)
		app.DELETE("person/:id", routes.DeleteAnimal)
	}

	r.Run()
}
