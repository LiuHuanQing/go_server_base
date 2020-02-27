package routers

import "github.com/gin-gonic/gin"

func Setup() {
	r := gin.New()
	r.Run(":5050")
}
