package controllers

import (
	"net/http"
	"user/config"
	"user/models"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
  var user models.User

  if err := c.ShouldBindJSON(&user); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()});
    return
  }

  if err := config.DB.Create(&user).Error; err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }

  c.JSON(http.StatusCreated, gin.H{"message": "User created", "user": user})
}

func CreateUsers(c *gin.Context) {
	var users []models.User

	if err := c.ShouldBindJSON(&users); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx := config.DB.Begin()
	for _, user := range users {
		if err := tx.Create(&user).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	tx.Commit()

	c.JSON(http.StatusCreated, gin.H{"message": "Users created"})
}

func GetUser(c *gin.Context) {
  id := c.Param("id")
  var user models.User

  if err := config.DB.First(&user, id).Error; err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }

  c.JSON(http.StatusOK, gin.H{"user": user})
}

func GetUsers(c *gin.Context) {
  var users []models.User

  if err := config.DB.Find(&users).Error; err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }

  c.JSON(http.StatusOK, gin.H{"users": users})
}

func UpdateUser(c *gin.Context) {
  id := c.Param("id")
  var user models.User

  if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

  if err := c.ShouldBindJSON(&user); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  if err := config.DB.Save(&user).Error; err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return 
  }

  c.JSON(http.StatusAccepted, gin.H{"message": "User Updated", "user": user})

}

func DeleteUser(c *gin.Context) {
  id := c.Param("id")
  var user models.User

  if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

  if err := config.DB.Delete(&user).Error; err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
  }

  c.JSON(http.StatusAccepted, gin.H{"message":"user deleted", "user": user})
}
