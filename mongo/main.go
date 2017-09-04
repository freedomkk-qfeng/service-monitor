package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/freedomkk-qfeng/service-monitor/mongo/cron"
	"github.com/freedomkk-qfeng/service-monitor/mongo/funcs"
	"github.com/freedomkk-qfeng/service-monitor/mongo/g"
)

func main() {

	cfg := flag.String("c", "cfg.json", "configuration file")
	version := flag.Bool("v", false, "show version")
	check := flag.Bool("check", false, "check collector")

	flag.Parse()

	if *version {
		fmt.Println(g.VERSION)
		os.Exit(0)
	}

	g.ParseConfig(*cfg)

	//g.InitRootDir()
	//g.InitRpcClients()

	if *check {
		funcs.CheckCollector()
		os.Exit(0)
	}

	funcs.BuildMappers()

	cron.Collect()

	select {}

}
