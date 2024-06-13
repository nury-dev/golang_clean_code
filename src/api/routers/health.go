package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/nazarow/golang-clean/api/handlers"
)

func Health(r *gin.RouterGroup) {
	handler := handlers.NewHealthHandler()

	r.GET("/", handler.Health)
}
