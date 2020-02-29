package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "Hello World!")
	})

	r.GET("/fib/:n", func(c *gin.Context) {
		n, _ := strconv.Atoi(c.Param("n"))
		fibn := fib(n)
		c.JSON(200, fmt.Sprintf("fib(%d)=%d", n, fibn))
	})

	r.GET("/sleep/:n", func(c *gin.Context) {
		n, _ := strconv.Atoi(c.Param("n"))
		time.Sleep(time.Duration(n) * time.Millisecond)
		c.JSON(200, fmt.Sprintf("sleep(%d)", n))
	})

	port := os.Getenv("PORT")
	r.Run(fmt.Sprintf(":%s", port))
}
