package server

import (
	"database/sql"
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
	err = LoginQuery(Db.Conn, data)
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

func LoginQuery(conn *sql.DB, data *LoginRequestData) error {
	m := &database.Member{}

	LoginQuery := "SELECT email, passwd FROM member WHERE email = ? and passwd = ?"
	stmt, err := conn.Prepare(LoginQuery)

	if err != nil {
		return fmt.Errorf("prepare query fail : [%s] error $%v", LoginQuery, err)
	}

	err = stmt.QueryRow(data.Email, data.Passwd).Scan(&m.Email, &m.Passwd)
	defer stmt.Close()

	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("no row in databases. query : %s", LoginQuery)
		}
	}

	if m.Email == "" || m.Passwd == "" {
		return fmt.Errorf("Login Fail. email : %s, passwd : %s", m.Email, m.Passwd)
	}

	return nil
}
