package main

import (
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"github.com/wulie/go-plugin/common"
	"os"
)

type IfconfigerLocal struct {
	logger hclog.Logger
}

func (i *IfconfigerLocal) Ifconfig() string {
	i.logger.Debug("message from IfconfigerLocal.Ifconfig")
	return "hello ifconfig "
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
