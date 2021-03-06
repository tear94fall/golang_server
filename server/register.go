package server

import (
	"database/sql"
	"fmt"
	"main/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterRequestData struct {
	Email  string `id:"email"`
	Passwd string `passwd:"passwd"`
	Name   string `name:"name"`
	Age    int    `age:"age"`
	Tel    string `tel:"tel"`
}

func Register(c *gin.Context) {
	data := &RegisterRequestData{}

	err := c.ShouldBind(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad reqeust",
		})
	}

	Db, _ := c.MustGet("mysql").(*database.MysqlConn)

	err = RegisterQuery(Db.Conn, data)
	if err != nil {
		fmt.Println(err)
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

func RegisterQuery(conn *sql.DB, data *RegisterRequestData) error {
	registerQeury := "INSERT INTO member (email, passwd, name, age, tel) VALUES (?, ?, ?, ?, ?)"
	stmt, err := conn.Prepare(registerQeury)

	if err != nil {
		return fmt.Errorf("prepare query fail : [%s] error $%v", registerQeury, err)
	}

	_, err = stmt.Exec(data.Email, data.Passwd, data.Name, data.Age, data.Tel)
	defer stmt.Close()

	if err != nil {
		return fmt.Errorf("execute query fail : [%s] error $%v", registerQeury, err)
	}

	return nil
}
