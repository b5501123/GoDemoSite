package controller

import (
	"DemoSite/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UserControllerInit(server *gin.Engine) {
	server.GET("/user/:id", get)
	server.PUT("/user/:id", update)
	server.POST("/user", create)
	server.DELETE("/user/:id", delete)
}

// @Summary User Update
// @Tags User
// @version 1.0
// @produce text/plain
// @param id path int true "id"
// @Success 200 {string} string
// @Router /user/{id} [get]
func get(c *gin.Context) {
	var user models.User
	id, _ := strconv.Atoi(c.Param("id"))
	db := models.DB
	db.Table("User").First(&user, id)
	c.JSON(http.StatusOK, user)
}

// @Summary User Update
// @Tags User
// @version 1.0
// @produce application/json
// @param id path int true "id"
// @param params body models.User true "params"
// @Success 200 {object} models.User "{"msg":"OK"}"
// @Failure 400 {string} json "{"msg":"fail"}"
// @Router /user/{id} [put]
func update(c *gin.Context) {
	var user models.User
	id, _ := strconv.Atoi(c.Param("id"))
	err := c.Bind(&user)

	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"msg": err.Error(),
		})
		return
	}
	db := models.DB
	db.Table("User").Where("ID = ?", id).Update("Name", user.Name)
	user.ID = id
	c.JSON(http.StatusOK, user)
}

// @Summary User Create
// @Tags User
// @version 1.0
// @produce application/json
// @param params body models.User true "params"
// @Success 200 {object} models.User "{"msg":"OK"}"
// @Failure 400 {string} json "{"msg":"fail"}"
// @Router /user [post]
func create(c *gin.Context) {
	var user models.User
	err := c.Bind(&user)

	if err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"msg": err.Error(),
		})
		return
	}
	db := models.DB
	db.Table("User").Create(user)
	c.JSON(http.StatusOK, user)
}

// @Summary User Delete
// @Tags User
// @version 1.0
// @produce text/plain
// @param id path int true "id"
// @Success 200 {string} string
// @Router /user/{id} [delete]
func delete(c *gin.Context) {
	var user models.User
	id, _ := strconv.Atoi(c.Param("id"))
	db := models.DB
	db.Table("User").Where("id = ?", id).Delete(user)
	c.JSON(http.StatusOK, gin.H{
		"msg": "OK",
	})
}
