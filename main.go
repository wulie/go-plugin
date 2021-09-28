package main

import (
	"fmt"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"github.com/wulie/go-plugin/common"
	"os"
	"os/exec"
)

var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "Ifconfig_plugin",
	MagicCookieValue: "hello",
}

var pluginMap = map[string]plugin.Plugin{
	"ifconfiger": &common.IfconfigerPlugin{},
}

func main() {
	logger := hclog.New(&hclog.LoggerOptions{
		Name:   "plugin",
		Output: os.Stdout,
		Level:  hclog.Debug,
	})

	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
		Cmd:             exec.Command("./plugin/plugin"),
		Logger:          logger,
	})
	defer client.Kill()
	rpc, err := client.Client()
	if err != nil {
		panic(err)
	}
	raw, err := rpc.Dispense("ifconfiger")
	if err != nil {
		panic(err)
	}

	ifconfig := raw.(common.Ifconfiger)
	fmt.Println(ifconfig.Ifconfig())

}
