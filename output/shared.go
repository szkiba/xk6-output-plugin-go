// SPDX-FileCopyrightText: 2023 Iván Szkiba
//
// SPDX-License-Identifier: MIT

package output

import (
	"context"

	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

type Output interface {
	Init(context.Context, *Params) (*Info, error)
	Start(context.Context) error
	Stop(context.Context) error
	AddMetrics(context.Context, []*Metric) error
	AddSamples(context.Context, []*Sample) error
}

var Handshake = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "XK6_OUTPUT_PLUGIN",
	MagicCookieValue: "1e427e4a-0ac3-11ee-9271-3fe99b1af532",
}

var PluginMap = map[string]plugin.Plugin{"output": new(Plugin)}

type Plugin struct {
	plugin.Plugin
	impl Output
}

func NewPlugin(impl Output) *Plugin {
	return &Plugin{impl: impl} // nolint:exhaustruct
}

func (p *Plugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	initServerLogger()

	RegisterOutputServer(s, newServer(p.impl))

	return nil
}

func (p *Plugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return newClient(NewOutputClient(c)), nil
}
