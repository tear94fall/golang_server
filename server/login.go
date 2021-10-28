package server

import (
	"fmt"
	"log"
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
	err = LoginQuery(Db, loginQuery)
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

func LoginQuery(conn *database.MysqlConn, query string) error {
	m := &database.Member{}

	rows, err := conn.Conn.Query(query)
	if err != nil {
		log.Printf("run query fail : [%s] error $%v", query, err)
		return fmt.Errorf("run query fail : [%s] error $%v", query, err)
	}
	defer rows.Close()

	count := 0

	for rows.Next() {
		err := rows.Scan(&m.Email, &m.Passwd)
		if err != nil {
			log.Fatal(err)
		}
		m.PrintMember()
		count += 1
	}

	if count == 0 {
		log.Printf("login fail")
		return fmt.Errorf("login fail")
	}

	return nil
}
