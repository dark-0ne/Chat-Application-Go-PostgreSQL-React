package routers

import (
	"log"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"

	"github.com/dark-0ne/Chat-Application-Go-PostgreSQL-React/middlewares"
	"github.com/dark-0ne/Chat-Application-Go-PostgreSQL-React/services"
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

		apiv1.GET("/users", services.GetUsers)
		apiv1.GET("/user/:id", services.GetUser)
		apiv1.POST("/user", services.PostUser)
		apiv1.PATCH("/user/:id", services.UpdateUser)
		apiv1.DELETE("/user/:id", services.DeleteUser)
	}

	return r
}
