package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/isaaccodekill/my-cal/services"
	"golang.org/x/oauth2"
)

type CalendarController struct {
	Config *oauth2.Config
}

func NewCalendarController(config *oauth2.Config) *CalendarController {
	return &CalendarController{
		Config: config,
	}
}

func (calController *CalendarController) ConnectCalendar(c *gin.Context) {
	url := services.GenerateAuthUrl(calController.Config)
	c.JSON(200, gin.H{
		"message": url,
	})
}

// the redirect url should be the same as the one in the google console
func (calController *CalendarController) RedirectAndSaveCalendar(c *gin.Context) {

}
