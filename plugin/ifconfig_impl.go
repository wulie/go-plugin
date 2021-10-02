package main

import (
	"fmt"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"github.com/wulie/go-plugin/common"
	"os"
)

type IfconfigerLocal struct {
	logger hclog.Logger
}

func (i *IfconfigerLocal) Ifconfig() []*common.NetInfo {
	i.logger.Debug("message from IfconfigerLocal.Ifconfig   77777")
	fmt.Println("666666666")

	return make([]*common.NetInfo, 0)
}

var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "Ifconfig_plugin",
	MagicCookieValue: "hello",
}

func main() {
	logger := hclog.New(&hclog.LoggerOptions{
		Level:      hclog.Trace,
		Output:     os.Stderr,
		JSONFormat: true,
	})

	ifconfiger := &IfconfigerLocal{
		logger: logger,
	}
	var pluginMap = map[string]plugin.Plugin{
		"ifconfiger": &common.IfconfigerPlugin{Impl: ifconfiger},
	}
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
	})

}
