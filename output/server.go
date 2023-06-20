// SPDX-FileCopyrightText: 2023 Iván Szkiba
//
// SPDX-License-Identifier: MIT

package output

import (
	"context"
	"log"
	"sync"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

type server struct {
	UnimplementedOutputServer
	impl Output
}

func newServer(impl Output) *server {
	return &server{impl: impl} // nolint:exhaustruct
}

func (s *server) Init(ctx context.Context, req *InitRequest) (*InitResponse, error) {
	config, err := s.impl.Init(ctx, req.Params)

	return &InitResponse{Info: config}, err // nolint:exhaustruct
}

func (s *server) Start(ctx context.Context, _ *Empty) (*Empty, error) {
	return new(Empty), s.impl.Start(ctx)
}

func (s *server) Stop(ctx context.Context, _ *Empty) (*Empty, error) {
	return new(Empty), s.impl.Stop(ctx)
}

func (s *server) AddMetrics(ctx context.Context, req *AddMetricsRequest) (*Empty, error) {
	return new(Empty), s.impl.AddMetrics(ctx, req.Metrics)
}

func (s *server) AddSamples(ctx context.Context, req *AddSamplesRequest) (*Empty, error) {
	return new(Empty), s.impl.AddSamples(ctx, req.Samples)
}

func Serve(impl Output) {
	plugin.Serve(&plugin.ServeConfig{ // nolint:exhaustruct
		HandshakeConfig: Handshake,
		Plugins: plugin.PluginSet{
			"output": NewPlugin(impl),
		},
		GRPCServer: plugin.DefaultGRPCServer,
	})
}

var loggerOnce sync.Once

func initServerLogger() {
	loggerOnce.Do(func() {
		logger := hclog.New(&hclog.LoggerOptions{JSONFormat: true}) // nolint:exhaustruct

		hclog.SetDefault(logger)
		log.SetFlags(0)
		log.SetOutput(logger.StandardWriter(&hclog.StandardLoggerOptions{InferLevels: true})) // nolint:exhaustruct
	})
}
