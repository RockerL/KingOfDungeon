package conf

import (
	"encoding/json"
	"github.com/name5566/leaf/log"
	"io/ioutil"
	"strconv"
)

var Server struct {
	LogLevel              string
	LogPath               string
	WSAddr                string
	CertFile              string
	KeyFile               string
	TCPAddr               string
	MaxConnNum            int
	ConsolePort           int
	ProfilePath           string
	DBAddr                string
	CTAddr                string
	MapLoad               []uint32
	MapRandom             []int64
	DefaultContentValue   uint32
	BlockContentInitValue map[string]int
}

func init() {
	data, err := ioutil.ReadFile("conf/server.json")
	if err != nil {
		log.Fatal("%v", err)
	}
	err = json.Unmarshal(data, &Server)
	if err != nil {
		log.Fatal("%v", err)
	}

	for k, v := range Server.BlockContentInitValue {
		i, _ := strconv.Atoi(k)
		t := uint16(i)
		BlockContentInitValue[t] = uint32(v)
	}

	for i := 0; i < len(Server.MapLoad); i++ {
		MapGenRandom[Server.MapLoad[i]] = Server.MapRandom[i]
	}
}

func GetContentInitValue(t uint16) uint32 {
	v, ok := BlockContentInitValue[t]
	if ok {
		return v
	}
	return Server.DefaultContentValue
}
