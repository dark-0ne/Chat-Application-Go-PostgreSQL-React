package v1

import (
	"log"
	"net/http"

	"github.com/dark-0ne/Chat-Application-Go-PostgreSQL-React/models"
	"github.com/gin-gonic/gin"
)

type NewConversation struct {
	Messages []models.Message `json:"messages"`
}

type ConversationUpdate struct {
	Messages []models.Message `json:"messages"`
}

func GetConversation(c *gin.Context) {

	var conversation models.Conversation

	db, err := models.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Where("id= ?", c.Param("conv_id")).Preload("Messages").First(&conversation).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Conversation not found"})
		return
	}

	c.JSON(http.StatusOK, conversation)

}

func PostConversation(c *gin.Context) {

	var conversation NewConversation

	if err := c.ShouldBindJSON(&conversation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newConversation := models.Conversation{Messages: conversation.Messages}

	db, err := models.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Create(&newConversation).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, newConversation)
}

func UpdateConversation(c *gin.Context) {

	var conversation models.Conversation

	db, err := models.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Where("id = ?", c.Param("conv_id")).First(&conversation).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Conversation not found!"})
		return
	}

	var updateConversation ConversationUpdate

	if err := c.ShouldBindJSON(&updateConversation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Model(&conversation).Updates(models.Conversation{Messages: updateConversation.Messages}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, conversation)

}

func DeleteConversation(c *gin.Context) {

	var conversation models.Conversation

	db, err := models.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Where("id = ?", c.Param("conv_id")).First(&conversation).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Conversation not found!"})
		return
	}

	if err := db.Delete(&conversation).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"conversation": "Conversation deleted"})
}
