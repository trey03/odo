package git

import (
	"bytes"
	"odo"
)

// Clone clone git from remote
func Clone(c *odo.Context) (string, error) {
	bu := bytes.NewBuffer(nil)
	bu.WriteString("begin git clone [" + c.AppConfig.AppName + "]\n")
	c.Session.Stderr = bu
	c.Session.Stdout = bu
	_, err := c.Session.Command("git", "clone", c.AppConfig.GitUrl).Output()
	bu.WriteString("clone done.")
	return bu.String(), err
}
