package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// User
	r.POST("/login", LoginUser)
	r.GET("/user", TokenAuthMiddleware(), getUser)
	r.GET("/user/:id", TokenAuthMiddleware(), getSingleUser)
	r.POST("/user", TokenAuthMiddleware(), createUser)
	r.PUT("/user/:id", TokenAuthMiddleware(), editUser)
	r.DELETE("/user/:id", TokenAuthMiddleware(), deleteUser)

	// Data
	r.GET("/data", TokenAuthMiddleware(), getData)
	r.POST("/data", TokenAuthMiddleware(), createData)
	r.DELETE("/data/:id", TokenAuthMiddleware(), deleteData)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
