package routes

import (
	"github.com/gin-gonic/gin"
	"v1/controller"
)

func UpFiletoMinio(r *gin.RouterGroup) {
	user := r.Group("file")
	{
		user.POST("/upload", controller.UploadFile)
	}
}
