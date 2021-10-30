package server

import (
	"database/sql"
	"fmt"
	"main/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeleteRequestData struct {
	Email  string `id:"email"`
	Passwd string `passwd:"passwd"`
}

func Delete(c *gin.Context) {
	data := &UpdateRequestData{}

	err := c.ShouldBind(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad reqeust",
		})
	}

	Db, _ := c.MustGet("mysql").(*database.MysqlConn)

	err = DeleteQuery(Db.Conn, data)
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

func DeleteQuery(conn *sql.DB, data *UpdateRequestData) error {
	DeleteQuery := "DELETE FROM member WHERE email = ? and passwd = ?"
	stmt, err := conn.Prepare(DeleteQuery)

	if err != nil {
		return fmt.Errorf("prepare query fail : [%s] error $%v", DeleteQuery, err)
	}

	_, err = stmt.Exec(data.Email, data.Passwd)
	defer stmt.Close()

	if err != nil {
		return fmt.Errorf("execute query fail : [%s] error $%v", DeleteQuery, err)
	}

	return nil
}
