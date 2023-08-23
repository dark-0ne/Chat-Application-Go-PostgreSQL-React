package v1

import (
	"log"
	"net/http"

	"github.com/dark-0ne/Chat-Application-Go-PostgreSQL-React/models"
	"github.com/gin-gonic/gin"
)

type NewMessage struct {
	Text           string `json:"username" binding:"required"`
	Read           bool   `json:"read" gorm:"default false"`
	SenderID       uint   `json:"sernder_id" binding:"required"`
	ConversationID uint   `json:"conversations_id" binding:"required"`
}

type MessageUpdate struct {
	Text           string `json:"username"`
	Read           bool   `json:"read"`
	SenderID       uint   `json:"sernder_id"`
	ConversationID uint   `json:"conversations_id"`
}

func GetMessage(c *gin.Context) {

	var message models.Message

	db, err := models.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Where("id= ?", c.Param("id")).First(&message).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Message not found"})
		return
	}

	c.JSON(http.StatusOK, message)

}

func PostMessage(c *gin.Context) {

	var message NewMessage

	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newMessage := models.Message{Text: message.Text, Read: message.Read, SenderID: message.SenderID, ConversationID: message.ConversationID}

	db, err := models.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Create(&newMessage).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, newMessage)
}

func UpdateMessage(c *gin.Context) {

	var message models.Message

	db, err := models.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Where("id = ?", c.Param("id")).First(&message).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Message not found!"})
		return
	}

	var updateMessage MessageUpdate

	if err := c.ShouldBindJSON(&updateMessage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Model(&message).Updates(models.Message{Text: updateMessage.Text, Read: updateMessage.Read, SenderID: updateMessage.SenderID, ConversationID: updateMessage.ConversationID}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, message)

}

func DeleteMessage(c *gin.Context) {

	var message models.Message

	db, err := models.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Where("id = ?", c.Param("id")).First(&message).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Message not found!"})
		return
	}

	if err := db.Delete(&message).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Message deleted"})

}
