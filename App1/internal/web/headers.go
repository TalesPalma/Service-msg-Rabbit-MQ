package web

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/TalesPalma/GolangRabbitMQ/internal/models"
	"github.com/TalesPalma/GolangRabbitMQ/internal/rabbit"
	"github.com/gin-gonic/gin"
)

func index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}

func loadMessages(ctx *gin.Context) {
	var list []struct {
		Id   uint   `json:"ID"`
		Text string `json:"Text"`
	}

	response, error := http.Get("http://localhost:8081/messages")

	if error != nil {
		ctx.HTML(http.StatusOK, "error.html", gin.H{"Error": error})
		return
	}

	defer response.Body.Close()

	if error := json.NewDecoder(response.Body).Decode(&list); error != nil {
		ctx.HTML(http.StatusOK, "error.html", gin.H{"Error": error})
		return
	}

	for i := range list {
		fmt.Println("LISTA:", list[i], " FIM")
	}

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

	rabbit.NewRabbit().SendMessage(content) // Ja manda para a fila correta
	ctx.JSON(http.StatusOK, message)
}

func deleteMessage(ctx *gin.Context) {
	id := ctx.Param("id")
	log.Println("id:", id)

	url := "http://localhost:8081/messages/" + id

	req, err := http.NewRequest(http.MethodDelete, url, nil)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err})
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": err})
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": resp.Status})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Status": resp.StatusCode})
}
