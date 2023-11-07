package main

import (
	"github.com/gin-gonic/gin"
	"github.com/isaaccodekill/my-cal/config"
	"github.com/isaaccodekill/my-cal/controllers"
)

var appConfig *config.Config

func main() {

	// Load .env file
	appConfig = config.LoadConfig()

	// create oauth2 config
	calConfig := config.CreateCalConfigContext(appConfig)

	router := gin.Default()

	// create CalendarController to inject CalendarConfig
	calController := controllers.NewCalendarController(calConfig)

	calRouter := router.Group("/cal")
	calRouter.GET("/connect", calController.ConnectCalendar)

	router.Run(":8080")
}
