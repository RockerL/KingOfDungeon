package internal

import (
	"center/conf"
	"center/gameserver"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"shared"
)

type Module struct {
	*gate.Gate //组合了Leaf的Gate类型，一起实现了Module
}

func (m *Module) OnInit() {
	m.Gate = &gate.Gate{
		MaxConnNum:      conf.Server.GSMaxConnNum,
		PendingWriteNum: conf.PendingWriteNum,
		MaxMsgLen:       conf.MaxMsgLen,
		WSAddr:          conf.Server.WSAddr,
		HTTPTimeout:     conf.HTTPTimeout,
		CertFile:        conf.Server.CertFile,
		KeyFile:         conf.Server.KeyFile,
		TCPAddr:         conf.Server.GSTCPAddr,
		LenMsgLen:       conf.LenMsgLen,
		LittleEndian:    conf.LittleEndian,
		Processor:       shared.GSCTProcessor,
		AgentChanRPC:    gameserver.ChanRPC,
	}

	log.Debug("gate server module init")
}
