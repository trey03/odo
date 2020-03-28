package git

import (
	"fmt"
	"odo"
	"os"
)

// GitWorker git work flow
type GitWorker struct {
	context *odo.Context
}

// Work git work
func (g *GitWorker) Work(c *odo.Context) (string, error) {
	logFile := c.AppConfig.WorkDir + "/git.log"
	out,err := os.Create(logFile)
	if err != nil {
		return "", err
	}
	defer out.Close()
	c.Session.Stderr = out
	c.Session.Stdout = out
	
	//拉取代码库
	msg, err := Clone(c)
	fmt.Println(msg)
	if err != nil {
		return msg, err
	}

	//创建一个release分支
	// ... checking out to commit
	msg, err = ReleaseBranch(c)
	fmt.Println(msg)
	if err != nil {
		return msg, err
	}
	
	//合并多个开发分支
	msg, err = Merge(c)
	fmt.Println(msg)
	if err != nil {
		return msg, err
	}

	return msg, nil
}
