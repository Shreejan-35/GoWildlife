package main

import (
	"go-wildlife/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() //router

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
