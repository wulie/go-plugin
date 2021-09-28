package common

import (
	"github.com/hashicorp/go-plugin"
	"net/rpc"
	"os/exec"
)

type Ifconfiger interface {
	Ifconfig() string
}

type IfconfigerRPC struct {
	client *rpc.Client
}

func (i *IfconfigerRPC) Ifconfig() string {
	var resp string
	err := i.client.Call("Plugin.Ifconfig", new(interface{}), &resp)
	if err != nil {
		panic(err)
	}
	cmd := exec.Command("ifconfig")
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	resp += string(output)
	return resp
}

type IfconfigerRCPServer struct {
	Impl Ifconfiger
}

func (i *IfconfigerRCPServer) Ifconfig(args interface{}, resp *string) error {
	*resp = i.Impl.Ifconfig()
	return nil
}

type IfconfigerPlugin struct {
	Impl Ifconfiger
}

func (i *IfconfigerPlugin) Server(broker *plugin.MuxBroker) (interface{}, error) {
	return &IfconfigerRCPServer{Impl: i.Impl}, nil
}

func (i *IfconfigerPlugin) Client(broker *plugin.MuxBroker, client *rpc.Client) (interface{}, error) {
	return &IfconfigerRPC{client}, nil
}
