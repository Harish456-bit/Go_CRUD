package controllers

import (
	models "Models/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReadPost(create *gin.Context) {
	var posts []models.Post
	tdn := create.Param("tdn")
	models.DB.Where("tdn = ?", tdn).First(&posts)
	if err := models.DB.Find(&posts).Error; err != nil {
		create.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	create.JSON(http.StatusOK, gin.H{"data": posts})
}

func FindPost(create *gin.Context) {
	var post models.Post
	tdn := create.Param("tdn")
	if err := models.DB.Where("tdn = ?", tdn).First(&post).Error; err != nil {
		create.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}
	create.JSON(http.StatusOK, gin.H{"Data": post})
}

func UpdatePost(create *gin.Context) {
	var post models.Post
	if err := models.DB.Where("tdn = ?", create.Param("tdn")).First(&post).Error; err != nil {
		create.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}
	var input models.Update
	if err := create.ShouldBindJSON(&input); err != nil {
		create.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Model(&post).Updates(models.Post{
		TrainName:   input.TrainName,
		TrainType:   input.TrainType,
		TrainNumber: input.TrainNumber,
	})

	create.JSON(http.StatusOK, gin.H{"data": post})
}

func DeletePost(create *gin.Context) {
	var post models.Post
	if err := models.DB.Where("tdn = ?", create.Param("tdn")).First(&post).Error; err != nil {
		create.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "record not found"})
		return
	}

	models.DB.Delete(&post)
	create.JSON(http.StatusOK, gin.H{"data": post})
}
