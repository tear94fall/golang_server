package main

import (
	"fmt"
	"main/database"
	"main/server"

	"github.com/gin-gonic/gin"
)

func main() {
	printStartMsg()

	r := SetupRouter()
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

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(MiddleMysql())

	r.POST("/api/v1/login", server.Login)
	r.POST("/api/v1/register", server.Register)

	return r
}

// Middleware
func MiddleMysql() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.InitMysqlConn()

		if db == nil {
			panic("init mysql database conn fail")
		}

		c.Set("mysql", db)
		c.Next()
	}
}

// Logger
func MiddleLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}
