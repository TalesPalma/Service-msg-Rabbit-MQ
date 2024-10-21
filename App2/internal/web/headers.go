package web

import (
	"fmt"
	"net/http"

	"github.com/TalesPalma/App2/internal/db"
	"github.com/TalesPalma/App2/internal/models"
	"github.com/gin-gonic/gin"
)

func index(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Hello, World!")
}

func getMessages(ctx *gin.Context) {

	var listMessages []models.Message

	if response := db.Db.Find(&listMessages); response.Error != nil {
		errorApi(ctx, response.Error)
		return
	}

	ctx.JSON(http.StatusOK, listMessages)

}

func errorApi(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}

func deleteMessage(ctx *gin.Context) {
	id := ctx.Param("id")
	fmt.Println("id:", id)

	if resp := db.Db.Delete(&models.Message{}, id); resp.Error != nil {
		errorApi(ctx, resp.Error)
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
