package routes

import (
	"os"
	"tes-database-pq/docs"
	"tes-database-pq/handlers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "tes-database-pq/docs"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	// Set Swagger host dinamis dari environment variable
	if host := os.Getenv("SWAGGER_HOST"); host != "" {
		docs.SwaggerInfo.Host = host
	}

	// Swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// User routes
	users := r.Group("/users")
	{
		users.GET("", handlers.GetAllUsers)
		users.GET("/:id", handlers.GetUser)
		users.POST("", handlers.CreateUser)
		users.PUT("/:id", handlers.UpdateUser)
		users.DELETE("/:id", handlers.DeleteUser)
	}

	// Root endpoint
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to User CRUD API",
			"docs":    "/swagger/index.html",
		})
	})

	return r
}
