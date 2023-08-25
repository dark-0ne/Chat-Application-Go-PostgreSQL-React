package routers

import (
	"log"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"

	"github.com/dark-0ne/Chat-Application-Go-PostgreSQL-React/middlewares"
	"github.com/dark-0ne/Chat-Application-Go-PostgreSQL-React/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	authMiddleware := middlewares.GetJWTMiddleware()

	r.POST("/login", authMiddleware.LoginHandler)

	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	apiv1 := r.Group("/api/v1")
	apiv1.GET("/refresh_token", authMiddleware.RefreshHandler)
	apiv1.Use(authMiddleware.MiddlewareFunc())
	{
		apiv1.GET("/users", v1.GetAllUsers)
		apiv1.GET("/user/:uid", v1.GetUser)
		apiv1.POST("/user", v1.PostUser)
		apiv1.PATCH("/user/:uid", v1.UpdateUser)
		apiv1.DELETE("/user/:uid", v1.DeleteUser)

		apiv1.GET("/conversation/:conv_id", v1.GetConversation)
		apiv1.POST("/conversation", v1.PostConversation)
		apiv1.PATCH("/conversation/:conv_id", v1.UpdateConversation)
		apiv1.DELETE("/conversation/:conv_id", v1.DeleteConversation)

		apiv1.GET("/conversation/:conv_id/messages", v1.GetAllMessages)
		apiv1.GET("/conversation/:conv_id/message/:msg_id", v1.GetMessage)
		apiv1.POST("/conversation/:conv_id/message", v1.PostMessage)
		apiv1.PATCH("/conversation/:conv_id/message/:msg_id", v1.UpdateMessage)
		apiv1.DELETE("/conversation/:conv_id/message/:id", v1.DeleteMessage)
	}

	return r
}
