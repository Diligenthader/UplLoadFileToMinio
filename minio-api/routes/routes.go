package routes

import "github.com/gin-gonic/gin"

func SetRoute(r *gin.Engine) *gin.Engine {
	Test := r.Group("/Test")
	{
		UpFiletoMinio(Test)
	}
	return r
}
