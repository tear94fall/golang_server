package main

import (
	"fmt"
	"main/database"
	"main/server"

	"github.com/gin-gonic/gin"
)

func main() {
	printStartMsg()

	mysql := database.InitMysqlConn()
	if mysql == nil {
		panic(fmt.Errorf("cannot read database conf file"))
	}

	defer mysql.Conn.Close()

	r := SetupRouter(mysql)
	r.Run()
}

func printStartMsg() {
	colorReset := "\033[0m"
	colorCyan := "\033[36m"

	fmt.Print(string(colorCyan))
	fmt.Println("     ______         _____                                 ")
	fmt.Println("    / ____/___     / ___/___  ______   _____  _____       ")
	fmt.Println("   / / __/ __ \\    \\__ \\/ _ \\/ ___/ | / / _ \\/ ___/  ")
	fmt.Println("  / /_/ / /_/ /   ___/ /  __/ /   | |/ /  __/ /           ")
	fmt.Println("  \\____/\\____/   /____/\\___/_/    |___/\\___/_/        ")
	fmt.Println("                                                          ")
	fmt.Print(string(colorReset))
}

func SetupRouter(mysql *database.MysqlConn) *gin.Engine {
	r := gin.Default()

	r.Use(MiddleMysql(mysql))

	r.POST("/api/v1/login", server.Login)
	r.POST("/api/v1/register", server.Register)
	r.POST("/api/v1/update", server.Update)
	r.POST("/api/v1/delete", server.Delete)

	return r
}

// Mysql Middleware
func MiddleMysql(mysql *database.MysqlConn) gin.HandlerFunc {
	return func(c *gin.Context) {
		if mysql == nil {
			panic("init mysql database conn fail")
		}

		c.Set("mysql", mysql)
		c.Next()
	}
}

// Logger
func MiddleLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}
