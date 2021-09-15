package main

import (
	"sing3demons/go-swagger/docs"
	"sing3demons/go-swagger/handler"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func main() {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := gin.Default()
	r.GET("/healthcheck", handler.HealthCheckHandler)
	v1 := r.Group("/api/v1")
	{
		customers := v1.Group("/customers")
		{
			customersHandler := handler.CustomerHandler{}
			customers.GET(":id", customersHandler.GetCustomer)
			customers.GET("", customersHandler.ListCustomers)
			customers.POST("", customersHandler.CreateCustomer)
			customers.DELETE(":id", customersHandler.DeleteCustomer)
			customers.PATCH(":id", customersHandler.UpdateCustomer)
		}
	}

	// use ginSwagger middleware to serve the API docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}
