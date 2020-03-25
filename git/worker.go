package git

import (
	"fmt"
	"odo"
)

// GitWorker git work flow
type GitWorker struct {
	context *odo.Context
}

// Work git work
func (g *GitWorker) Work(c *odo.Context) (string, error) {

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
