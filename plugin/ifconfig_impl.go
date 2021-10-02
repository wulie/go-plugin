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
	nets := make([]*common.NetInfo, 0)

	i.logger.Debug("message from IfconfigerLocal.Ifconfig   77777")
	fmt.Println("666666666")
	nets = append(nets, &common.NetInfo{
		Name: "eth0",
		Ip:   "192.168.100.169",
		Mac:  "23:67:df:sd:67:23",
	})
	nets = append(nets, &common.NetInfo{
		Name: "eth1",
		Ip:   "192.168.0.169",
		Mac:  "23:67:df:sd:67:67",
	})
	nets = append(nets, &common.NetInfo{
		Name: "eth2",
		Ip:   "192.168.1.169",
		Mac:  "23:67:df:sd:67:67",
	})
	return nets
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
