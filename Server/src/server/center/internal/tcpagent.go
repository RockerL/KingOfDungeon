package internal

import (
	"github.com/name5566/leaf/chanrpc"
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/network"
	"net"
	"reflect"
	"shared"
)

type TCPAgent struct {
	conn      *network.TCPConn
	msgParser *network.MsgParser
	processor network.Processor
	chanRPC   *chanrpc.Server
	userData  interface{}
}

func NewTCPAgent(c net.Conn, msgParser *network.MsgParser, chanRPC *chanrpc.Server) *TCPAgent {
	return &TCPAgent{
		conn:      network.NewTCPConn(c, 10, msgParser),
		msgParser: msgParser,
		processor: shared.GSCTProcessor,
		chanRPC:   chanRPC,
	}
}

func (a *TCPAgent) Run() {
	for {
		data, err := a.msgParser.Read(a.conn)
		if err != nil {
			log.Debug("read message: %v", err)
			break
		}

		if a.processor != nil {
			msg, err := a.processor.Unmarshal(data)
			if err != nil {
				log.Debug("unmarshal message error: %v", err)
				break
			}
			err = a.processor.Route(msg, a)
			if err != nil {
				log.Debug("route message error: %v", err)
				break
			}
		}
	}
}

func (a *TCPAgent) OnClose() {
	if a.chanRPC != nil {
		err := a.chanRPC.Call0("CloseAgent", a)
		if err != nil {
			log.Error("chanrpc error: %v", err)
		}
	}
}

func (a *TCPAgent) WriteMsg(msg interface{}) {
	if a.processor != nil {
		data, err := a.processor.Marshal(msg)
		if err != nil {
			log.Error("marshal message %v error: %v", reflect.TypeOf(msg), err)
			return
		}
		err = a.conn.WriteMsg(data...)
		if err != nil {
			log.Error("write message %v error: %v", reflect.TypeOf(msg), err)
		}
	}
}

func (a *TCPAgent) LocalAddr() net.Addr {
	return a.conn.LocalAddr()
}

func (a *TCPAgent) RemoteAddr() net.Addr {
	return a.conn.RemoteAddr()
}

func (a *TCPAgent) Close() {
	a.conn.Close()
}

func (a *TCPAgent) Destroy() {
	a.conn.Destroy()
}

func (a *TCPAgent) UserData() interface{} {
	return a.userData
}

func (a *TCPAgent) SetUserData(data interface{}) {
	a.userData = data
}
