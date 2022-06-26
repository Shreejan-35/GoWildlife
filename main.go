package main

import (
	"go-wildlife/models"
	"go-wildlife/routes"

	"log"

	"github.com/gin-gonic/contrib/static"

	"github.com/gin-gonic/gin"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	r := gin.Default() //router

	r.Use(static.Serve("/", static.LocalFile("./templates", true)))

	models.ConnectDb()
	app := r.Group("/api")

	{
		app.GET("animals", routes.GetAnimals)
		app.GET("animal/:id", routes.GetAnimalById)
		app.POST("addanimal", routes.AddAnimal)
		app.PUT("upanimal/:id", routes.UpdateAnimal)
		app.DELETE("delanimal/:id", routes.DeleteAnimal)
	}

	r.Run()
}
