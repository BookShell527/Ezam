package controllers

import (
	"context"
	"ezam/config"
	"ezam/middleware"
	"ezam/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"golang.org/x/crypto/bcrypt"
)

func StudentRoute(r *gin.RouterGroup) {
	r.POST("/register", register)
	r.POST("/login", login)
}

func register(c *gin.Context) {
	var register_student models.StudentRegisterDTO

	if err := c.ShouldBindJSON(&register_student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte(register_student.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	register_student.Password = string(password)
	result, err := config.StudentColl.InsertOne(context.TODO(), register_student)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("Student inserted: %s", result.InsertedID)
	tokenString, err := middleware.GenerateJwt(register_student.Name, "student")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Can't generate jwt"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"data":  result.InsertedID,
		"token": tokenString,
	})
}

func login(c *gin.Context) {
	var login_student models.StudentLoginDTO
	if err := c.ShouldBindJSON(&login_student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong JSON template"})
		return
	}

	var student_fromDb models.Student
	filter := bson.D{
		{"$or", bson.A{
			bson.D{{"nrp", login_student.Identifier}},
			bson.D{{"name", login_student.Identifier}},
			bson.D{{"email", login_student.Identifier}},
		}},
	}

	if err := config.StudentColl.FindOne(context.TODO(), filter).Decode(&student_fromDb); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Student not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(student_fromDb.Password), []byte(login_student.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Wrong Password"})
		return
	}

	tokenString, err := middleware.GenerateJwt(student_fromDb.Name, "student")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Can't generate jwt token"})
	}

	c.JSON(http.StatusAccepted, gin.H{
		"student": student_fromDb.Id,
		"token":   tokenString,
	})
}
