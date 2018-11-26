package internal

import (
	"center/conf"
	"center/login"
	"center/msg"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

type Module struct {
	*gate.Gate				//组合了Leaf的Gate类型，一起实现了Module
}

func (m *Module) OnInit() {
	m.Gate = &gate.Gate{
		MaxConnNum:      conf.Server.MaxConnNum,
		PendingWriteNum: conf.PendingWriteNum,
		MaxMsgLen:       conf.MaxMsgLen,
		WSAddr:          conf.Server.WSAddr,
		HTTPTimeout:     conf.HTTPTimeout,
		CertFile:        conf.Server.CertFile,
		KeyFile:         conf.Server.KeyFile,
		TCPAddr:         conf.Server.TCPAddr,
		LenMsgLen:       conf.LenMsgLen,
		LittleEndian:    conf.LittleEndian,
		Processor:       msg.Processor,
		AgentChanRPC:    login.ChanRPC,
	}

	log.Debug("gate module init")
}
