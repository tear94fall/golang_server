package server

import (
	"fmt"
	"log"
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
	data := RegisterRequestData{}

	err := c.ShouldBind(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad reqeust",
		})
	}

	Db, _ := c.MustGet("mysql").(*database.MysqlConn)
	registerQuery := fmt.Sprintf("INSERT INTO "+
		" member (email, passwd, name, age, tel) "+
		" VALUES (\"%s\", \"%s\", \"%s\", %d, \"%s\") ",
		data.Email,
		data.Passwd,
		data.Name,
		data.Age,
		data.Tel)

	err = RegisterQuery(Db, registerQuery)
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

func RegisterQuery(conn *database.MysqlConn, query string) error {
	result, err := conn.Conn.Exec(query)
	if err != nil {
		log.Printf("run query fail : [%s] error $%v", query, err)
		return fmt.Errorf("run query fail : [%s] error $%v", query, err)
	}

	if err != nil {
		log.Fatal(err)
	}

	n, _ := result.RowsAffected()
	if n == 1 {
		fmt.Println("1 row inserted.")
	}

	return nil
}
