package git

import (
	"bytes"
	"odo"
)

// Merge git merge
func Merge(c *odo.Context) (string, error) {
	bu := bytes.NewBuffer(nil)
	var branches = c.AppEnvInfo.DevBranches
	c.Session.Stderr = bu
	c.Session.Stdout = bu
	for _, br := range branches {
		bu.WriteString("Begin merge [" + br + "]\n")
		bs, err := c.Session.Command("git", "merge", br).Output()
		if err != nil {
			return bu.String(), err
		}
		bu.Write(bs)
	}
	bu.WriteString("Merge done.")
	return bu.String(), nil
}
