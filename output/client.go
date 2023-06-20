// SPDX-FileCopyrightText: 2023 Iván Szkiba
//
// SPDX-License-Identifier: MIT

package output

import "context"

type client struct {
	impl OutputClient
}

func newClient(impl OutputClient) *client {
	return &client{impl: impl}
}

func (c *client) Init(ctx context.Context, params *Params) (*Info, error) {
	resp, err := c.impl.Init(ctx, &InitRequest{Params: params}) // nolint:exhaustruct
	if err != nil {
		return nil, err
	}

	return resp.Info, nil
}

func (c *client) Start(ctx context.Context) error {
	_, err := c.impl.Start(ctx, new(Empty))

	return err
}

func (c *client) Stop(ctx context.Context) error {
	_, err := c.impl.Stop(ctx, new(Empty))

	return err
}

func (c *client) AddMetrics(ctx context.Context, metrics []*Metric) error {
	_, err := c.impl.AddMetrics(ctx, &AddMetricsRequest{Metrics: metrics}) // nolint:exhaustruct

	return err
}

func (c *client) AddSamples(ctx context.Context, samples []*Sample) error {
	_, err := c.impl.AddSamples(ctx, &AddSamplesRequest{Samples: samples}) // nolint:exhaustruct

	return err
}
