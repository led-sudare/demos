package main

import (
	"flag"
	"math/rand"
	"time"

	"sudare_contents/lib/content"
	"sudare_contents/lib/util"
)

type Configs struct {
	ZmqTarget string `json:"zmqTarget"`
}

func NewConfigs() Configs {
	return Configs{
		ZmqTarget: "0.0.0.0:5510",
	}
}

func main() {
	util.InitColorUtil()
	rand.Seed(time.Now().UnixNano())

	configs := NewConfigs()
	util.ReadConfig(&configs)

	var (
		optInputPort = flag.String("r", configs.ZmqTarget, "Specify IP and port of server zeromq SUB running.")
	)

	flag.Parse()

	endpoint := "tcp://" + *optInputPort

	sender := content.NewContentSender(endpoint)
	contents := []content.CylinderContent{
		content.NewContentSinWideLine(),
		content.NewContentSinLine(),
		content.NewContentCirWave(),
	}
	
	sender.SetContentToPlay(contents, 10*time.Second)

}
