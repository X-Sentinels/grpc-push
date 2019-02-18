package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/X-Sentinels/grpc-push/server/controller"
	"github.com/X-Sentinels/grpc-push/server/g"
	gs "github.com/X-Sentinels/grpc-push/server/grpc-server"
)

func main() {
	cfg := flag.String("c", "cfg.json", "configuration file")
	version := flag.Bool("v", false, "show version")
	flag.Parse()

	if *version {
		fmt.Println(g.VERSION)
		os.Exit(0)
	}

	g.ParseConfig(*cfg)

	g.NotifMessage = make(chan g.Message, g.Config().ChannelCache)

	go gs.Start()

	go controller.StartGin(g.Config().Http.Listen)

	select {}
}
