package controller

import (
	jwtDomain "DemoSite/domain/jwt"
	"DemoSite/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthControllerInit(server *gin.Engine) {
	server.POST("/auth/login", login)
}

// @Summary User LogIn
// @Tags Auth
// @version 1.0
// @produce application/json
// @param params body models.User true "params"
// @Success 200 {string} json "{"token":""}"
// @Failure 400 {string} json "{"msg":"fail"}"
// @Router /auth/login [post]
func login(c *gin.Context) {
	var user models.User
	err := c.Bind(&user)

	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"msg": err.Error(),
		})
		return
	}

	userInfo := jwtDomain.UserInfo{user.ID, user.Name}
	token, _ := jwtDomain.GenToken(userInfo)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
