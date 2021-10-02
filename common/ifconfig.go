package common

import (
	"fmt"
	"github.com/hashicorp/go-plugin"
	"net/rpc"
)

type NetInfo struct {
	Name string
	Ip   string
	Mac  string
}

func (i *NetInfo) String() string {
	return fmt.Sprintf("name:%s ip:%s mac:%s", i.Name, i.Ip, i.Mac)
}

type Ifconfiger interface {
	Ifconfig() []*NetInfo
}

type IfconfigerRPC struct {
	client *rpc.Client
}

func (i *IfconfigerRPC) Ifconfig() []*NetInfo {
	//var resp string
	nets := make([]*NetInfo, 0)
	err := i.client.Call("Plugin.Ifconfig", new(interface{}), &nets)
	if err != nil {
		panic(err)
	}
	//cmd := exec.Command("ifconfig")
	//output, err := cmd.Output()
	//if err != nil {
	//	panic(err)
	//}
	//nets := make([]*NetInfo,0)
	nets = append(nets, &NetInfo{
		Name: "eth0",
		Ip:   "192.168.100.169",
		Mac:  "23:67:df:sd:67:23",
	})
	nets = append(nets, &NetInfo{
		Name: "eth1",
		Ip:   "192.168.0.169",
		Mac:  "23:67:df:sd:67:67",
	})
	//resp = string(output)
	//fmt.Println(nets)
	return nets
}

type IfconfigerRCPServer struct {
	Impl Ifconfiger
}

func (i *IfconfigerRCPServer) Ifconfig(args interface{}, resp *[]*NetInfo) error {
	fmt.Println("666666666666666666")
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
