package odo

import "odo/sh"

// AppConfig app config
type AppConfig struct {
	AppId   string
	AppName string
	GitUrl  string
	WorkDir string
}
type AppEnvInfo struct {
	Env    string
	Branch string
}

type Worker interface {
	Work(c *AppConfig) (string, error)
}

type Context struct {
	AppConfig  *AppConfig
	Session    *sh.Session
	AppEnvInfo *AppEnvInfo
}
