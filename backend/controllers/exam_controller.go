package controllers

import (
	"context"
	"ezam/config"
	"ezam/middleware"
	"ezam/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ExamRoute(r *gin.RouterGroup) {
	r.POST("/create", middleware.AuthMiddleware, create_exam)
}

func create_exam(c *gin.Context) {
	var exam models.ExamCreateDTO
	if err := c.ShouldBindJSON(&exam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Doesn't match the schema json"})
		return
	}

	result, err := config.ExamColl.InsertOne(context.TODO(), exam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error inserting exam"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"status": "Exam Insertion Successfull", "data": result})
}
