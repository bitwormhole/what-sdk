package main

import (
	"os"

	"github.com/bitwormhole/what-sdk/modules/whatsdk"
	"github.com/starter-go/starter"
)

func main() {

	m := whatsdk.ModuleBenchmark()
	args := os.Args

	i := starter.Init(args)
	i.MainModule(m)
	i.WithPanic(true)
	i.Run()
}
