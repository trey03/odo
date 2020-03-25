package git

import (
	"bytes"
	"odo"
)

// Clone clone git from remote
func Clone(c *odo.Context) (string, error) {
	bu := bytes.NewBuffer(nil)
	bu.WriteString("Begin git clone [" + c.AppConfig.AppName + "]\n")
	c.Session.Stderr = bu
	c.Session.Stdout = bu
	_, err := c.Session.Command("git", "clone", c.AppConfig.GitUrl).Output()
	bu.WriteString("Clone done.\n")

	//进入工程目录
	newWorkDir := c.AppConfig.WorkDir+c.AppConfig.AppName
	bu.WriteString("\nEnter ["+c.AppConfig.AppName+"] directory ["+newWorkDir+"]\n")
	c.Session.SetDir(newWorkDir)
	return bu.String(), err
}
