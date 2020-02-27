package routers

import "github.com/gin-gonic/gin"

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// apiv1 := r.Group("/api/v1")
	return r

}
