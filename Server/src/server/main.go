package main

import (
	"fmt"
	"github.com/name5566/leaf"
	lconf "github.com/name5566/leaf/conf"
	"server/center"
	"server/client"
	"server/conf"
	"server/gate"
)

func main() {
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.Server.ConsolePort
	lconf.ProfilePath = conf.Server.ProfilePath

	fmt.Println("game server start")

	leaf.Run(
		center.Module,
		client.Module,
		gate.Module,
	)
}
