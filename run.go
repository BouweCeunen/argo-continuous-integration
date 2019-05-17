package main

import "k8s.io/client-go/rest"
import "github.com/gin-gonic/gin"
import "fmt"

var webhook = "/webhook"

func main() {
    rest.InClusterConfig()

    gin.DisableConsoleColor()
    r := gin.Default()

    r.POST(webhook, func(c *gin.Context) {
        c.JSON(202, gin.H{})

        fmt.Printf("Started Argo workflow %q in namespace %q.\n", workflow_name)
    })
    r.Run()
}
