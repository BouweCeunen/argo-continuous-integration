package main

import (
    "bytes"
    "fmt"
    "github.com/buger/jsonparser"
    "github.com/gin-gonic/gin"
    "io/ioutil"
    "os/exec"
    "strconv"
    "strings"
    "time"
)

func main() {
    gin.DisableConsoleColor()
    gin.SetMode(gin.ReleaseMode)

    r := gin.Default()

    r.POST("/", func(c *gin.Context) {
        c.JSON(202, gin.H{})

        buf := new(bytes.Buffer)
        _, _ = buf.ReadFrom(c.Request.Body)
        
        gitRepo,_ := jsonparser.GetString(buf.Bytes(), "repository", "full_name")
        gitRevision,_ := jsonparser.GetString(buf.Bytes(), "push", "changes", "[0]", "new", "name")
        
        gitRepoName := strings.Split(gitRepo, "/")[1]
        fullGitRepo := "git@bitbucket.org:" + gitRepo + ".git"
        timestamp := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)

        argoFilename := "argo" + timestamp + ".yml"

        input, _ := ioutil.ReadFile("argo.yml")
        temp := bytes.Replace(input, []byte("<git_repo_name>"), []byte(gitRepoName), -1)
        temp2 := bytes.Replace(temp, []byte("<git_repo_full>"), []byte(fullGitRepo), -1)
        output := bytes.Replace(temp2, []byte("<git_revision>"), []byte(gitRevision), -1)
        _ = ioutil.WriteFile(argoFilename, output, 0666)

        commandOutput, err := exec.Command("sh", "-c", "./argo submit " + argoFilename).CombinedOutput()
        if err != nil {
            fmt.Printf("Accepted webhook request, did NOT start Argo workflow: git_repo=%q,git_revision=%q, because of: %q\n", gitRepo, gitRevision, string(err.Error()))
        } else {
            fmt.Printf("Accepted webhook request, started Argo workflow: git_repo=%q,git_revision=%q, with message: %q\n", gitRepo, gitRevision, string(commandOutput))
        }
    })
    _ = r.Run(":3000")
}
