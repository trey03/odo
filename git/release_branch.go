package git

import (
	"bytes"
	"odo"
)

// ReleaseBranch release branch
func ReleaseBranch(c *odo.Context) (string,error){

	bu := bytes.NewBuffer(nil)
	bu.WriteString("Begin checkout [" + c.AppEnvInfo.Branch + "]\n")
	c.Session.Stderr = bu
	c.Session.Stdout = bu
	_,err := c.Session.Command("git","checkout","-b",c.AppEnvInfo.Branch).Output()

	return bu.String(), err
}