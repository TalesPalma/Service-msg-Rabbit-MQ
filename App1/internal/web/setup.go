package web

import (
	"text/template"

	"github.com/gin-gonic/gin"
)

func InitServer() {
	r := gin.Default()
	routers(r)

	templates, err := LoadTemplates()
	if err != nil {
		panic(err)
	}
	r.SetFuncMap(template.FuncMap{})
	r.SetHTMLTemplate(templates)
	r.Static("/css", "./templates/css")
	err = r.Run(":8080")
	if err != nil {
		panic(err)
	}
}

func routers(r *gin.Engine) {
	r.GET("/", index)
	r.GET("/message", loadMessages)
	r.POST("/message", postMessage)
	r.DELETE("/message/:id", deleteMessage)
}
