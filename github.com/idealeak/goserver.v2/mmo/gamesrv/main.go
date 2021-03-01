package main

import (
	_ "github.com/idealeak/goserver.v2/mmo"

	"github.com/idealeak/goserver.v2/core"
	"github.com/idealeak/goserver.v2/core/module"
)

func main() {
	defer core.ClosePackages()
	core.LoadPackages("config.json")

	waiter := module.Start()
	waiter.Wait()
}
