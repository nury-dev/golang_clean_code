package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/nazarow/golang-clean/api/handlers"
	"github.com/nazarow/golang-clean/api/middlewares"
	"github.com/nazarow/golang-clean/config"
)

func User(router *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewUsersHandler(cfg)

	router.POST("/send-otp", middlewares.OtpLimiter(cfg), h.SendOtp)
	router.POST("/login-by-username", h.LoginByUsername)
	router.POST("/register-by-username", h.RegisterByUsername)
	router.POST("/login-by-mobile", h.RegisterLoginByMobileNumber)
}
