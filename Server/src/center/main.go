package main

import (
	"center/conf"
	"center/game"
	"center/gate"
	"center/login"

	"github.com/name5566/leaf"
	lconf "github.com/name5566/leaf/conf"
	"github.com/name5566/leaf/log"
)

func main() {
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.Server.ConsolePort
	lconf.ProfilePath = conf.Server.ProfilePath

	log.Release("game center start...")

	//启动每个模块的协程，等待退出信号
	leaf.Run(
		login.Module,		//login模块负责处理玩家的登录，角色管理，分配玩家进入GameServer
		game.Module,		//game模块负责处理游戏相关的内容，主要跟GameServer进程进行交互
		gate.Module,		//最后运行gate模块，允许玩家开始登录
	)
}
