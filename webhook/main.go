package main

import (
    "github.com/gin-gonic/gin"
    "fmt"
    "github.com/buger/jsonparser"
    "bytes"
    // "io/ioutil"
    // "os/exec"
)

var webhook = "/webhook"

func main() {
    gin.DisableConsoleColor()
    gin.SetMode(gin.ReleaseMode)
    
    r := gin.Default()

    r.POST(webhook, func(c *gin.Context) {
        c.JSON(202, gin.H{})

        buf := new(bytes.Buffer)
	    buf.ReadFrom(c.Request.Body)

        git_repo,_ := jsonparser.GetString(buf.Bytes(), "person")
        git_revision,_ := jsonparser.GetString(buf.Bytes(), "person")

        // input, _ := ioutil.ReadFile("argo.yml")
        // temp := bytes.Replace(input, []byte("<git_repo>"), []byte(git_repo), -1)
        // output := bytes.Replace(temp, []byte("<git_revision>"), []byte(git_revision), -1)
        // ioutil.WriteFile("argo.yml", output, 0666)

        // command_output, _ := exec.Command("kubectl apply -f argo.yml").CombinedOutput()
        // fmt.Println(string(command_output))

        fmt.Printf("Accepted webhook request, started Argo workflow: git_repo=%q,git_revision=%q\n", git_repo, git_revision)
    })
    r.Run(":3000")
}
