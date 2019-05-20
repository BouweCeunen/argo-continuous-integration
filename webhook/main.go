package main

import (
    "github.com/gin-gonic/gin"
    "fmt"
    "github.com/buger/jsonparser"
    "bytes"
    "io/ioutil"
    "os/exec"
    "strconv"
    "time"
    "strings"
)

func main() {
    gin.DisableConsoleColor()
    gin.SetMode(gin.ReleaseMode)

    r := gin.Default()

    r.POST("/", func(c *gin.Context) {
        c.JSON(202, gin.H{})

        buf := new(bytes.Buffer)
	    buf.ReadFrom(c.Request.Body)
        
        git_repo,_ := jsonparser.GetString(buf.Bytes(), "repository", "full_name")
        git_revision,_ := jsonparser.GetString(buf.Bytes(), "push", "changes", "[0]", "new", "name")
        
        git_repo_name := strings.Split(git_repo, "/")[1]
        full_git_repo := "git@bitbucket.org:" + git_repo + ".git"
        timestamp := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)

        argo_filename := "argo" + timestamp + ".yml"

        input, _ := ioutil.ReadFile("argo.yml")
        temp := bytes.Replace(input, []byte("<git_repo_name>"), []byte(git_repo_name), -1)
        temp2 := bytes.Replace(temp, []byte("<git_repo_full>"), []byte(full_git_repo), -1)
        output := bytes.Replace(temp2, []byte("<git_revision>"), []byte(git_revision), -1)
        ioutil.WriteFile(argo_filename, output, 0666)

        command_output, err := exec.Command("sh", "-c", "./argo submit " + argo_filename).CombinedOutput()
        if err != nil {
            fmt.Printf("Accepted webhook request, did NOT start Argo workflow: git_repo=%q,git_revision=%q, because of: %q\n", git_repo, git_revision, string(err.Error()))
        } else {
            fmt.Printf("Accepted webhook request, started Argo workflow: git_repo=%q,git_revision=%q, with message: %q\n", git_repo, git_revision, string(command_output))
        }
    })
    r.Run(":3000")
}
