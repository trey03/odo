package git

import (
	"odo"
)

// Merge git merge
func Merge(c *odo.Context) (string, error) {
	out := c.Session.Stdout
	var branches = c.AppEnvInfo.DevBranches

	for _, br := range branches {
		out.Write([]byte("Begin merge [" + br + "]\n"))
		bs, err := c.Session.Command("git", "merge", br).Output()
		if err != nil {
			return "git merge", err
		}
		out.Write([]byte(bs))
	}
	out.Write([]byte("Merge done."))
	return "git merge", nil
}
