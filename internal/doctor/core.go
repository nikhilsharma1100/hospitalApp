package doctor

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllEntities(context *gin.Context) {
	doctors := GetAll()

	context.JSON(http.StatusOK, gin.H{"data": doctors})
}

func GetEntityById(context *gin.Context, id uint) {
	doctor, err := FindUserById(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusOK, gin.H{"data": doctor})
}

func CreateEntity(context *gin.Context) {
	// Read request input here
	var entity Doctor
	Create(entity)

	context.JSON(http.StatusCreated, gin.H{"data": "created"})
}

func UpdateEntity(context *gin.Context, entity Doctor) {
	Update(entity)

	context.JSON(http.StatusOK, gin.H{"data": "updated"})
}
