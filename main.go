package main

import (
	_ "github.com/ItsFimall/ForceHiddifyConnection/hiddify_extension"

	"github.com/hiddify/hiddify-core/extension/server"
)

func main() {
	server.StartTestExtensionServer()
}
