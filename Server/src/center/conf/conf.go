package conf

import (
	"log"
	"time"
)

var (
	// log conf
	LogFlag = log.LstdFlags

	// gate conf
	PendingWriteNum        = 2000
	MaxMsgLen       uint32 = 102400
	HTTPTimeout            = 10 * time.Second
	LenMsgLen              = 4
	LittleEndian           = true

	// skeleton conf
	GoLen              = 10000
	TimerDispatcherLen = 10000
	AsynCallLen        = 10000
	ChanRPCLen         = 10000
)
