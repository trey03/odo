package git

type Worker interface{
	work(appId string) 
}

type GitWorker struct{
	appId string,
}

func (g *GitWorker) work(appId string) {
	//根据appId获取配置信息，如git url
	url := "  "
}