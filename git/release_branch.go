package git

import (
	"odo"
)

// ReleaseBranch release branch
func ReleaseBranch(c *odo.Context) (string, error) {
	out := c.Session.Stdout
	out.Write([]byte("Begin checkout [" + c.AppEnvInfo.Branch + "]\n"))
	_, err := c.Session.Command("git", "checkout", "-b", c.AppEnvInfo.Branch).Output()
	return "git checkout", err
}
