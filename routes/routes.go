package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAnimals(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetAnimals Called"})
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
