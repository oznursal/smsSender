package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"smsSender/controller"
	_ "smsSender/docs"
	"smsSender/worker"
)

//	@title			Swagger Sms Sender API
//	@version		2.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		localhost:8080
// @BasePath	/api/v1
func main() {
	r := gin.Default()
	c := controller.NewController()
	worker.Start()

	v1 := r.Group("/api/v1")
	{
		messages := v1.Group("/messages")
		{
			messages.GET("", c.ListMessages)
		}

		scheduler := v1.Group("/scheduler")
		{
			scheduler.PUT("start", c.Start)
			scheduler.PUT("stop", c.Stop)
		}

	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
