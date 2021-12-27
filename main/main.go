package main

import (
	"fmt"
	"main/database"
	"main/server"
	"main/utility"

	"github.com/gin-gonic/gin"
)

func main() {
	printStartMsg()

	mysql, err := database.InitMysqlConn()
	if mysql == nil {
		utility.CmdClear()
		printFailMsg()
		fmt.Println(err)
		return
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

func printFailMsg() {
	colorReset := "\033[0m"
	colorRed := "\033[31m"

	fmt.Print(string(colorRed))
	fmt.Println("     _____                              ____                 ______      _ __     ")
	fmt.Println("    / ___/___  ______   _____  _____   / __ \\__  ______     / ____/___ _(_) /    ")
	fmt.Println("    \\__ \\/ _ \\/ ___/ | / / _ \\/ ___/  / /_/ / / / / __ \\   / /_  / __ `/ / / ")
	fmt.Println("   ___/ /  __/ /   | |/ /  __/ /     / _, _/ /_/ / / / /  / __/ / /_/ / / /      ")
	fmt.Println("  /____/\\___/_/    |___/\\___/_/     /_/ |_|\\__,_/_/ /_/  /_/    \\__,_/_/_/   ")
	fmt.Println("                                                                                 ")
	fmt.Print(string(colorReset))
}

func SetupRouter(mysql *database.MysqlConn) *gin.Engine {
	r := gin.Default()

	r.Use(MiddleMysql(mysql))

	r.POST("/api/v1/login", server.Login)
	r.POST("/api/v1/register", server.Register)
	r.POST("/api/v1/update", server.Update)

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
