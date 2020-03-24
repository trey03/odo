package main

import (
	"bytes"
	"fmt"
	"odo/sh"
)

func main() {
	session := sh.NewSession()
	bu := bytes.NewBuffer(nil)
	session.Stderr = bu
	session.PipeFail = true
	session.SetDir("/Users/chenfeng/work/opensource/flow-wks")
	//err := session.Command("git", "pull").WriteStdout("/tmp/listing.txt")
	session.Command("git", "pull")
	err := session.WriteStdout("/tmp/listing.txt")
	fmt.Println("--------1")
	//err := session.WriteStdout("/tmp/listing.txt")

	fmt.Println("-------2")

	if err != nil {
		fmt.Println(err)
		fmt.Println(bu)
	}

}
