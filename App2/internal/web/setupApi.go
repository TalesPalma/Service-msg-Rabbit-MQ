package web

import "github.com/gin-gonic/gin"

func InitWebServer() {
	r := gin.Default()
	routers(r)

	if err := r.Run(":8081"); err != nil {
		panic(err)
	}

}

func routers(r *gin.Engine) {
	r.GET("/", index)
	r.GET("/messages", getMessages)
}
