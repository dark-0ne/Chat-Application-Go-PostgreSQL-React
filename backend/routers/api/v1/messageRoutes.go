package v1

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/dark-0ne/Chat-Application-Go-PostgreSQL-React/models"
	"github.com/dark-0ne/Chat-Application-Go-PostgreSQL-React/util"
)

type NewMessage struct {
	Text     string `json:"text" binding:"required"`
	Read     bool   `json:"read" gorm:"default false"`
	SenderID uint   `json:"sender_id" binding:"required"`
}

type MessageUpdate struct {
	Text           string `json:"text"`
	Read           bool   `json:"read"`
	SenderID       uint   `json:"sender_id"`
	ConversationID uint   `json:"conversations_id"`
}

func GetAllMessages(c *gin.Context) {

	var messages []models.Message

	db, err := models.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Where("conversation_id= ?", c.Param("conv_id")).Find(&messages).Error; err != nil || len(messages) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Messages not found"})
		return
	}

	c.JSON(http.StatusOK, messages)

}

func GetMessage(c *gin.Context) {

	var message models.Message

	db, err := models.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Where("id= ?", c.Param("msg_id")).First(&message).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Message not found"})
		return
	}

	c.JSON(http.StatusOK, message)

}

func PostMessage(c *gin.Context) {

	// Check if conversaiton exists
	var conversation models.Conversation

	db, err := models.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Where("id = ?", c.Param("conv_id")).First(&conversation).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Conversation not found!"})
		return
	}

	// Create a new message
	var message NewMessage

	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newMessage := models.Message{Text: message.Text, Read: message.Read, SenderID: message.SenderID, ConversationID: util.Str2Uint(c.Param("conv_id"))}

	if err := db.Create(&newMessage).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Update conversation to include new message

	conversation.Messages = append(conversation.Messages, newMessage)
	fmt.Printf("%+v", conversation)

	if err := db.Model(&conversation).Updates(models.Conversation{Messages: conversation.Messages}).Error; err != nil {
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

	if err := db.Where("id = ?", c.Param("msg_id")).First(&message).Error; err != nil {
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

	if err := db.Where("id = ?", c.Param("msg_id")).First(&message).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Message not found!"})
		return
	}

	if err := db.Delete(&message).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Message deleted"})

}
