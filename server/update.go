package server

import (
	"database/sql"
	"fmt"
	"main/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateRequestData struct {
	Email  string `id:"email"`
	Passwd string `passwd:"passwd"`
	Name   string `name:"name"`
	Age    int    `age:"age"`
	Tel    string `tel:"tel"`
}

func Update(c *gin.Context) {
	data := &UpdateRequestData{}

	err := c.ShouldBind(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad reqeust",
		})
	}

	Db, _ := c.MustGet("mysql").(*database.MysqlConn)

	err = UpdateQuery(Db.Conn, data)
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

func UpdateQuery(conn *sql.DB, data *UpdateRequestData) error {
	UpdateQuery := "UPDATE member SET name=?, age=?, tel=? WHERE email=?"
	stmt, err := conn.Prepare(UpdateQuery)

	if err != nil {
		return fmt.Errorf("prepare query fail : [%s] error $%v", UpdateQuery, err)
	}

	_, err = stmt.Exec(data.Name, data.Age, data.Tel, data.Email)
	defer stmt.Close()

	if err != nil {
		return fmt.Errorf("execute query fail : [%s] error $%v", UpdateQuery, err)
	}

	return nil
}
