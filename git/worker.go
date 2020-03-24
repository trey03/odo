package git

import "odo"

type GitWorker struct {
	context *odo.Context
}

// Work git work
func (g *GitWorker) Work(c *odo.Context) (string, error) {

	//拉取代码库
	msg, err := Clone(c)

	if err != nil {
		return msg, err
	}

	//进入工程目录

	//创建一个release分支
	// ... checking out to commit

	//合并多个开发分支
	//var branchs = []string{"release/test_01","release/test_02","release/test_03"}
	//for _, br := range branchs {
	//TODO 太吐血了，go-git还没有merge或re
	//}
	return msg, nil
}
