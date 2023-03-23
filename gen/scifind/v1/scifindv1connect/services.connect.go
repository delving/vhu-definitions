// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: scifind/v1/services.proto

package scifindv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/delving/scifind/gen/scifind/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// ObjectServiceName is the fully-qualified name of the ObjectService service.
	ObjectServiceName = "scifind.v1.ObjectService"
	// TimelineServiceName is the fully-qualified name of the TimelineService service.
	TimelineServiceName = "scifind.v1.TimelineService"
)

// ObjectServiceClient is a client for the scifind.v1.ObjectService service.
type ObjectServiceClient interface {
	GetObject(context.Context, *connect_go.Request[v1.GetObjectRequest]) (*connect_go.Response[v1.GetObjectResponse], error)
}

// NewObjectServiceClient constructs a client for the scifind.v1.ObjectService service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewObjectServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) ObjectServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &objectServiceClient{
		getObject: connect_go.NewClient[v1.GetObjectRequest, v1.GetObjectResponse](
			httpClient,
			baseURL+"/scifind.v1.ObjectService/GetObject",
			opts...,
		),
	}
}

// objectServiceClient implements ObjectServiceClient.
type objectServiceClient struct {
	getObject *connect_go.Client[v1.GetObjectRequest, v1.GetObjectResponse]
}

// GetObject calls scifind.v1.ObjectService.GetObject.
func (c *objectServiceClient) GetObject(ctx context.Context, req *connect_go.Request[v1.GetObjectRequest]) (*connect_go.Response[v1.GetObjectResponse], error) {
	return c.getObject.CallUnary(ctx, req)
}

// ObjectServiceHandler is an implementation of the scifind.v1.ObjectService service.
type ObjectServiceHandler interface {
	GetObject(context.Context, *connect_go.Request[v1.GetObjectRequest]) (*connect_go.Response[v1.GetObjectResponse], error)
}

// NewObjectServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewObjectServiceHandler(svc ObjectServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/scifind.v1.ObjectService/GetObject", connect_go.NewUnaryHandler(
		"/scifind.v1.ObjectService/GetObject",
		svc.GetObject,
		opts...,
	))
	return "/scifind.v1.ObjectService/", mux
}

// UnimplementedObjectServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedObjectServiceHandler struct{}

func (UnimplementedObjectServiceHandler) GetObject(context.Context, *connect_go.Request[v1.GetObjectRequest]) (*connect_go.Response[v1.GetObjectResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("scifind.v1.ObjectService.GetObject is not implemented"))
}

// TimelineServiceClient is a client for the scifind.v1.TimelineService service.
type TimelineServiceClient interface {
	GetTimeline(context.Context, *connect_go.Request[v1.GetTimelineRequest]) (*connect_go.Response[v1.GetTimelineResponse], error)
}

// NewTimelineServiceClient constructs a client for the scifind.v1.TimelineService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewTimelineServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) TimelineServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &timelineServiceClient{
		getTimeline: connect_go.NewClient[v1.GetTimelineRequest, v1.GetTimelineResponse](
			httpClient,
			baseURL+"/scifind.v1.TimelineService/GetTimeline",
			opts...,
		),
	}
}

// timelineServiceClient implements TimelineServiceClient.
type timelineServiceClient struct {
	getTimeline *connect_go.Client[v1.GetTimelineRequest, v1.GetTimelineResponse]
}

// GetTimeline calls scifind.v1.TimelineService.GetTimeline.
func (c *timelineServiceClient) GetTimeline(ctx context.Context, req *connect_go.Request[v1.GetTimelineRequest]) (*connect_go.Response[v1.GetTimelineResponse], error) {
	return c.getTimeline.CallUnary(ctx, req)
}

// TimelineServiceHandler is an implementation of the scifind.v1.TimelineService service.
type TimelineServiceHandler interface {
	GetTimeline(context.Context, *connect_go.Request[v1.GetTimelineRequest]) (*connect_go.Response[v1.GetTimelineResponse], error)
}

// NewTimelineServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewTimelineServiceHandler(svc TimelineServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/scifind.v1.TimelineService/GetTimeline", connect_go.NewUnaryHandler(
		"/scifind.v1.TimelineService/GetTimeline",
		svc.GetTimeline,
		opts...,
	))
	return "/scifind.v1.TimelineService/", mux
}

// UnimplementedTimelineServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedTimelineServiceHandler struct{}

func (UnimplementedTimelineServiceHandler) GetTimeline(context.Context, *connect_go.Request[v1.GetTimelineRequest]) (*connect_go.Response[v1.GetTimelineResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("scifind.v1.TimelineService.GetTimeline is not implemented"))
}