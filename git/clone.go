package git

import (
	"odo"
)

// Clone clone git from remote
func Clone(c *odo.Context) (string, error) {
	out := c.Session.Stdout
	out.Write([]byte("Begin git clone [" + c.AppConfig.AppName + "]\n"))

	_, err := c.Session.Command("git", "clone", c.AppConfig.GitUrl).Output()
	out.Write([]byte("Clone done.\n"))

	//进入工程目录
	newWorkDir := c.AppConfig.WorkDir + c.AppConfig.AppName
	out.Write([]byte("\nEnter [" + c.AppConfig.AppName + "] directory [" + newWorkDir + "]\n"))
	c.Session.SetDir(newWorkDir)
	return "git", err
}
