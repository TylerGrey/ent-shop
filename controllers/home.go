package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HomeController ...
type HomeController struct{}

// Home ...
func (hc HomeController) Home(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{
		"Title": "HOME",
	})
}
