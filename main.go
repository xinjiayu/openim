package main

import (
	"github.com/gogf/gf/frame/g"
	_ "openim/boot"
	"openim/library/version"
	_ "openim/router"
)

var (
	BuildVersion = "0.0"
	BuildTime    = ""
	CommitID     = ""
)

func main() {
	version.ShowLogo(BuildVersion, BuildTime, CommitID)
	g.Server().Run()
}
