package main

import (
	"github.com/CodyGuo/jumpadmin/cmd"

	"github.com/CodyGuo/glog"
)

func main() {
	if err := cmd.Execute(); err != nil {
		glog.Fatal(err)
	}
}
