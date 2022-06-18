package routes

import (
	"go-wildlife/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GetAnimals(c *gin.Context) {
	wildlife, err := models.AnimalsGot()

	checkErr(err)

	if wildlife == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Records Found"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": wildlife})
	}
}

func GetAnimalById(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "GetAnimalById " + id + " Called"})
}

func AddAnimal(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "AddAnimal Called"})
}

func UpdateAnimal(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "UpdateAnimal Called"})
}

func DeleteAnimal(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "DeleteAnimal " + id + " Called"})
}
