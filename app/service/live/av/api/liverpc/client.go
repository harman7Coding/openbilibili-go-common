// Code generated by liverpcgen, DO NOT EDIT.
// source: *.proto files under this directory
// If you want to change this file, Please see README in go-common/app/tool/liverpc/protoc-gen-liverpc/
package liverpc

import (
	"go-common/app/service/live/av/api/liverpc/v0"
	"go-common/app/service/live/av/api/liverpc/v1"
	"go-common/library/net/rpc/liverpc"
)

// Client that represents a liverpc av service api
type Client struct {
	cli *liverpc.Client
	// V0PayGoods presents the controller in liverpc
	V0PayGoods v0.PayGoods
	// V0PayLive presents the controller in liverpc
	V0PayLive v0.PayLive
	// V1PayLive presents the controller in liverpc
	V1PayLive v1.PayLive
	// V1Pk presents the controller in liverpc
	V1Pk v1.Pk
}

// DiscoveryAppId the discovery id is not the tree name
var DiscoveryAppId = "live.av"

// New a Client that represents a liverpc live.av service api
// conf can be empty, and it will use discovery to find service by default
// conf.AppID will be overwrite by a fixed value DiscoveryAppId
// therefore is no need to set
func New(conf *liverpc.ClientConfig) *Client {
	if conf == nil {
		conf = &liverpc.ClientConfig{}
	}
	conf.AppID = DiscoveryAppId
	var realCli = liverpc.NewClient(conf)
	cli := &Client{cli: realCli}
	cli.clientInit(realCli)
	return cli
}

func (cli *Client) GetRawCli() *liverpc.Client {
	return cli.cli
}

func (cli *Client) clientInit(realCli *liverpc.Client) {
	cli.V0PayGoods = v0.NewPayGoodsRpcClient(realCli)
	cli.V0PayLive = v0.NewPayLiveRpcClient(realCli)
	cli.V1PayLive = v1.NewPayLiveRpcClient(realCli)
	cli.V1Pk = v1.NewPkRpcClient(realCli)
}
