package web

import (
	"log"
	"net/http"

	"github.com/TalesPalma/GolangRabbitMQ/internal/db"
	"github.com/TalesPalma/GolangRabbitMQ/internal/models"
	"github.com/TalesPalma/GolangRabbitMQ/internal/rabbit"
	"github.com/gin-gonic/gin"
)

func index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}

func loadMessages(ctx *gin.Context) {
	var list []models.Message
	db.Db.Find(&list)
	ctx.HTML(http.StatusOK, "messages.html", gin.H{"Items": list})
}

func postMessage(ctx *gin.Context) {
	channel := ctx.PostForm("channel")
	content := ctx.PostForm("content")
	message := models.NewMessage(content, channel)

	if message.NotValid() {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid message"})
		return
	}

	if response := db.Db.Create(&message); response.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": response.Error})
		return
	}

	rabbit.NewRabbit().SendMessage(content)
	ctx.JSON(http.StatusOK, message)
}

func deleteMessage(ctx *gin.Context) {
	id := ctx.Query("id")
	log.Println("id:", id)
	if response := db.Db.Where("id = ?", id).Delete(&models.Message{}); response.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": response.Error})
		return
	}
	ctx.JSON(http.StatusOK, nil)
}
