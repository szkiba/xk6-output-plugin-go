// SPDX-FileCopyrightText: 2023 Iván Szkiba
//
// SPDX-License-Identifier: MIT

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.3
// source: proto/output.proto

package output

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Output_Init_FullMethodName       = "/xk6_output_plugin.output.Output/Init"
	Output_Start_FullMethodName      = "/xk6_output_plugin.output.Output/Start"
	Output_Stop_FullMethodName       = "/xk6_output_plugin.output.Output/Stop"
	Output_AddMetrics_FullMethodName = "/xk6_output_plugin.output.Output/AddMetrics"
	Output_AddSamples_FullMethodName = "/xk6_output_plugin.output.Output/AddSamples"
)

// OutputClient is the client API for Output service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OutputClient interface {
	// Init is called before registering the output plugin.
	//
	// Init receives the environment variables of the k6 process as parameters.
	// In addition, standard command line arguments can be used to configure the plugin.
	// A description of the plugin and various configuration parameters for the xk6-output-plugin can be returned.
	Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*InitResponse, error)
	// Start is called before the k6 Engine tries to use the output and should be
	// used for any long initialization tasks.
	Start(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	// Flush all remaining metrics and finalize the test run.
	Stop(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	// AddMetrics is called on all metrics, the plugin can use it to save metric parameters.
	//
	// The call is made before AddSample is called with the given metric's sample.
	AddMetrics(ctx context.Context, in *AddMetricsRequest, opts ...grpc.CallOption) (*Empty, error)
	// AddSamples receives samples of the metrics periodically.
	AddSamples(ctx context.Context, in *AddSamplesRequest, opts ...grpc.CallOption) (*Empty, error)
}

type outputClient struct {
	cc grpc.ClientConnInterface
}

func NewOutputClient(cc grpc.ClientConnInterface) OutputClient {
	return &outputClient{cc}
}

func (c *outputClient) Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*InitResponse, error) {
	out := new(InitResponse)
	err := c.cc.Invoke(ctx, Output_Init_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *outputClient) Start(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Output_Start_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *outputClient) Stop(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Output_Stop_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *outputClient) AddMetrics(ctx context.Context, in *AddMetricsRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Output_AddMetrics_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *outputClient) AddSamples(ctx context.Context, in *AddSamplesRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Output_AddSamples_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OutputServer is the server API for Output service.
// All implementations must embed UnimplementedOutputServer
// for forward compatibility
type OutputServer interface {
	// Init is called before registering the output plugin.
	//
	// Init receives the environment variables of the k6 process as parameters.
	// In addition, standard command line arguments can be used to configure the plugin.
	// A description of the plugin and various configuration parameters for the xk6-output-plugin can be returned.
	Init(context.Context, *InitRequest) (*InitResponse, error)
	// Start is called before the k6 Engine tries to use the output and should be
	// used for any long initialization tasks.
	Start(context.Context, *Empty) (*Empty, error)
	// Flush all remaining metrics and finalize the test run.
	Stop(context.Context, *Empty) (*Empty, error)
	// AddMetrics is called on all metrics, the plugin can use it to save metric parameters.
	//
	// The call is made before AddSample is called with the given metric's sample.
	AddMetrics(context.Context, *AddMetricsRequest) (*Empty, error)
	// AddSamples receives samples of the metrics periodically.
	AddSamples(context.Context, *AddSamplesRequest) (*Empty, error)
	mustEmbedUnimplementedOutputServer()
}

// UnimplementedOutputServer must be embedded to have forward compatible implementations.
type UnimplementedOutputServer struct {
}

func (UnimplementedOutputServer) Init(context.Context, *InitRequest) (*InitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Init not implemented")
}
func (UnimplementedOutputServer) Start(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedOutputServer) Stop(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedOutputServer) AddMetrics(context.Context, *AddMetricsRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMetrics not implemented")
}
func (UnimplementedOutputServer) AddSamples(context.Context, *AddSamplesRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSamples not implemented")
}
func (UnimplementedOutputServer) mustEmbedUnimplementedOutputServer() {}

// UnsafeOutputServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OutputServer will
// result in compilation errors.
type UnsafeOutputServer interface {
	mustEmbedUnimplementedOutputServer()
}

func RegisterOutputServer(s grpc.ServiceRegistrar, srv OutputServer) {
	s.RegisterService(&Output_ServiceDesc, srv)
}

func _Output_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OutputServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Output_Init_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OutputServer).Init(ctx, req.(*InitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Output_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OutputServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Output_Start_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OutputServer).Start(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Output_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OutputServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Output_Stop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OutputServer).Stop(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Output_AddMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OutputServer).AddMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Output_AddMetrics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OutputServer).AddMetrics(ctx, req.(*AddMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Output_AddSamples_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSamplesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OutputServer).AddSamples(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Output_AddSamples_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OutputServer).AddSamples(ctx, req.(*AddSamplesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Output_ServiceDesc is the grpc.ServiceDesc for Output service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Output_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "xk6_output_plugin.output.Output",
	HandlerType: (*OutputServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Init",
			Handler:    _Output_Init_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _Output_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _Output_Stop_Handler,
		},
		{
			MethodName: "AddMetrics",
			Handler:    _Output_AddMetrics_Handler,
		},
		{
			MethodName: "AddSamples",
			Handler:    _Output_AddSamples_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/output.proto",
}
