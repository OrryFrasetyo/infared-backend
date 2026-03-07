package router

import (
	"infared-backend/internal/delivery/http/handler"
	"infared-backend/internal/delivery/http/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter(
	userHandler *handler.UserHandler,
	itemHandler *handler.ItemHandler,
	requestHandler *handler.RequestHandler,
) *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong! 🏓 Server InfaRed API menyala!",
		})
	})

	v1 := r.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/login", userHandler.Login)
		}

		protected := v1.Group("")
		protected.Use(middleware.AuthGuard())
		{
			protected.POST("/auth/register", middleware.RoleGuard("admin"), userHandler.RegisterRelawan)

			protected.POST("/items", middleware.RoleGuard("admin"), itemHandler.CreateItem)
			protected.GET("/items", itemHandler.GetAllItems)

			protected.POST("/requests/chat", requestHandler.ChatToAI)
		}
	}

	return r
}
