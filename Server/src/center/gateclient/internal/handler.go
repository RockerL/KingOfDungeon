package internal

import (
	"center/conf"
	"center/gameclient"
	"center/msg"
	"github.com/name5566/leaf/gate"
)

func init() {
	skeleton.RegisterChanRPC("GameServerRegistered", rpcGameServerRegistered)
}

var (
	clientGate *gate.Gate
	clientGateCloseSig = make(chan bool, 1)
)

//当所有的游戏服务器注册成功则开放客户端加入
func rpcGameServerRegistered(args []interface{}){
	clientGate = &gate.Gate{
		MaxConnNum:      conf.Server.CLMaxConnNum,
		PendingWriteNum: conf.PendingWriteNum,
		MaxMsgLen:       conf.MaxMsgLen,
		WSAddr:          conf.Server.WSAddr,
		HTTPTimeout:     conf.HTTPTimeout,
		CertFile:        conf.Server.CertFile,
		KeyFile:         conf.Server.KeyFile,
		TCPAddr:         conf.Server.CLTCPAddr,
		LenMsgLen:       conf.LenMsgLen,
		LittleEndian:    conf.LittleEndian,
		Processor:       msg.Processor,
		AgentChanRPC:    gameclient.ChanRPC,
	}

	go func() {
		clientGate.Run(clientGateCloseSig)
	}()
}
