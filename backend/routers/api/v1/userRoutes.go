package v1

import (
	"log"
	"net/http"

	"github.com/dark-0ne/Chat-Application-Go-PostgreSQL-React/models"
	"github.com/gin-gonic/gin"
)

type NewUser struct {
	Username      string                `json:"username" binding:"required"`
	Bio           string                `json:"bio"`
	Password      string                `json:"password" binding:"required"`
	Conversations []models.Conversation `json:"conversations"`
}

type UserUpdate struct {
	Username      string                `json:"username"`
	Bio           string                `json:"bio"`
	Password      string                `json:"password"`
	Conversations []models.Conversation `json:"conversations"`
}

func GetAllUsers(c *gin.Context) {

	var users []models.User

	db, err := models.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Find(&users).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)

}

func GetUser(c *gin.Context) {

	var user models.User

	db, err := models.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Where("id= ?", c.Param("uid")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)

}

func PostUser(c *gin.Context) {

	var user NewUser

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser := models.User{Username: user.Username, Bio: user.Bio, Password: user.Password, Conversations: user.Conversations}

	db, err := models.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, newUser)
}

func UpdateUser(c *gin.Context) {

	var user models.User

	db, err := models.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Where("id = ?", c.Param("uid")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found!"})
		return
	}

	var updateUser UserUpdate

	if err := c.ShouldBindJSON(&updateUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Model(&user).Updates(models.User{Username: updateUser.Username, Bio: updateUser.Bio, Password: updateUser.Password, Conversations: updateUser.Conversations}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)

}

func DeleteUser(c *gin.Context) {

	var user models.User

	db, err := models.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Where("id = ?", c.Param("uid")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found!"})
		return
	}

	if err := db.Unscoped().Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})

}
