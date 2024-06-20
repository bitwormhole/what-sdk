package main

import (
	"flag"
	"os"

	"github.com/bitwormhole/what-sdk/modules/whatsdk"
	"github.com/starter-go/starter"
)

func main() {

	var configFilePath = ""
	flag.StringVar(&configFilePath, "p", "", "本地配置文件的路径")
	flag.Parse()

	if configFilePath == "" {
		panic("缺少命令行参数：-p")
	}

	addProps := map[string]string{
		"application.properties.file":    configFilePath,
		"application.properties.enabled": "1",
	}

	// run
	m := whatsdk.ModuleBenchmark()
	args := os.Args

	i := starter.Init(args)
	i.GetProperties().Import(addProps)
	i.MainModule(m)
	i.WithPanic(true)
	i.Run()
}
