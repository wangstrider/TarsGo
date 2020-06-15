package main

import (
	"github.com/wangstrider/TarsGo/tars"
)

func main() { //Init servant
	cfg := tars.GetServerConfig()                                                //Get Config File Object
	proto := new(CustomProtocolImp)                                              //New Proto
	tars.AddServantWithProtocol(proto, cfg.App+"."+cfg.Server+".CustomProtoObj") //Register Servant
	tars.Run()
}
