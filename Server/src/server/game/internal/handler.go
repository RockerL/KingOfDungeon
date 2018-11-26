package internal

import (
	"reflect"
)

func init(){
	//handler(&msg.ReqLogin{}, handleMap)
}

func handler(m interface{}, h interface{}){
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

//处理地图内消息
func handleMap(args []interface{}){

}

