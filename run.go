package main

import "github.com/gin-gonic/gin"

func main() {
    gin.DisableConsoleColor()
    r := gin.Default()
    r.POST("/webhook", func(c *gin.Context) {
        c.JSON(202, gin.H{})
    })
    r.Run()
}

