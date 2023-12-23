package controllers

import (
	models "Models/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePost(create *gin.Context) {
	var inputs []models.Create
	if err := create.ShouldBindJSON(&inputs); err != nil {
		create.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, input := range inputs {
		post := models.Post{
			TDN:         input.TDN,
			TrainName:   input.TrainName,
			TrainType:   input.TrainType,
			TrainNumber: input.TrainNumber,
		}
		models.DB.Create(&post)
	}

	create.JSON(http.StatusOK, gin.H{"data": "Posts created successfully"})
}
