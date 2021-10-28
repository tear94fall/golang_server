package server

import (
	"fmt"
	"main/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequestData struct {
	Email  string `id:"email"`
	Passwd string `passwd:"passwd"`
}

func Login(c *gin.Context) {
	data := &LoginRequestData{}

	err := c.ShouldBind(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad reqeust",
		})
		return
	}

	Db, _ := c.MustGet("mysql").(*database.MysqlConn)
	loginQuery := fmt.Sprintf("SELECT email, passwd FROM member WHERE email = \"%s\" and passwd = \"%s\"", data.Email, data.Passwd)
	err = Db.RunQuery(loginQuery)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"email":  data.Email,
		"passwd": data.Passwd,
	})
}
