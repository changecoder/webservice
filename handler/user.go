package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserLogin User Login Handler
func UserLogin(c *gin.Context) {

}

// UserStatus User Status Handler
func UserStatus(c *gin.Context) {
	cookie, err := c.Cookie("cc_u")

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"authenticated": false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"authenticated": true,
		"cookie":        cookie,
	})
}

// UserInfo User Info Handler
func UserInfo(c *gin.Context) {

}
