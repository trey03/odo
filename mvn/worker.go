package mvn

import (
	"io"
	"odo"
	"os"
)

// MvnWorker mvn worker
type MvnWorker struct {
}

// Work mvn clean
func (m *MvnWorker) Work(c *odo.Context) (string, error) {

	f := c.AppConfig.WorkDir+"/mvn.log"
	out, err := os.Create(f)
	if err != nil {
		return "", err
	}
	defer out.Close()
	c.Session.Stderr = out
	c.Session.Stdout = out

	err = build(c)
	if err != nil {
		return "error", err
	}
	return "", err
}

func build(c *odo.Context) error {
	out := c.Session.Stdout

	write("maven building ["+c.AppConfig.AppName+"]\n", out)

	err := c.Session.Command("mvn", "clean", "package").Run()
	if err != nil {
		write("maven build fail.", out)
	} else {
		write("maven build done.", out)
	}
	return err
}

func write(m string, out io.Writer) {
	out.Write([]byte(m))
}
