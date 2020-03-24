package main

import (
	"fmt"
	"odo"
	"odo/git"
	"odo/sh"
	"os"
)

func main() {
	var appId = "odo123"
	//构建appConfig信息
	//Info("get application config by appId", appId)
	url := "https://github.com/trey03/odo.git"
	appname := "odo"
	workDir := "/Users/chenfeng/work/opensource/flow-wks/"

	os.RemoveAll(workDir + "/" + appname)

	//TODO 根据appId获取配置信息，如git url
	appConfig := &odo.AppConfig{
		AppId:   appId,
		AppName: appname,
		GitUrl:  url,
		WorkDir: workDir,
	}

	session := sh.NewSession()
	session.SetDir(workDir)
	session.PipeFail = true

	appEnvInfo := &odo.AppEnvInfo{
		Env:    "dev",
		Branch: "feature/20200324_dev_01",
	}

	context := &odo.Context{
		AppConfig:  appConfig,
		Session:    session,
		AppEnvInfo: appEnvInfo,
	}

	var gitWorker = new(git.GitWorker)
	_, err := gitWorker.Work(context)

	if err != nil {
		fmt.Println("error:", err)
	}

	//time.Sleep(10 * time.Second)
}
