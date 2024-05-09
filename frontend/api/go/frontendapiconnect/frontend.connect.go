// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: frontendapi/frontend.proto

package frontendapiconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	api "github.com/curioswitch/tasuke/frontend/api"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// FrontendServiceName is the fully-qualified name of the FrontendService service.
	FrontendServiceName = "frontendapi.FrontendService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// FrontendServiceSaveUserProcedure is the fully-qualified name of the FrontendService's SaveUser
	// RPC.
	FrontendServiceSaveUserProcedure = "/frontendapi.FrontendService/SaveUser"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	frontendServiceServiceDescriptor        = api.File_frontendapi_frontend_proto.Services().ByName("FrontendService")
	frontendServiceSaveUserMethodDescriptor = frontendServiceServiceDescriptor.Methods().ByName("SaveUser")
)

// FrontendServiceClient is a client for the frontendapi.FrontendService service.
type FrontendServiceClient interface {
	// Saves information for a user. This method works both for a new or existing user.
	// The user is identified by the firebase ID token included in the authorization header.
	SaveUser(context.Context, *connect.Request[api.SaveUserRequest]) (*connect.Response[api.SaveUserResponse], error)
}

// NewFrontendServiceClient constructs a client for the frontendapi.FrontendService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewFrontendServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) FrontendServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &frontendServiceClient{
		saveUser: connect.NewClient[api.SaveUserRequest, api.SaveUserResponse](
			httpClient,
			baseURL+FrontendServiceSaveUserProcedure,
			connect.WithSchema(frontendServiceSaveUserMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// frontendServiceClient implements FrontendServiceClient.
type frontendServiceClient struct {
	saveUser *connect.Client[api.SaveUserRequest, api.SaveUserResponse]
}

// SaveUser calls frontendapi.FrontendService.SaveUser.
func (c *frontendServiceClient) SaveUser(ctx context.Context, req *connect.Request[api.SaveUserRequest]) (*connect.Response[api.SaveUserResponse], error) {
	return c.saveUser.CallUnary(ctx, req)
}

// FrontendServiceHandler is an implementation of the frontendapi.FrontendService service.
type FrontendServiceHandler interface {
	// Saves information for a user. This method works both for a new or existing user.
	// The user is identified by the firebase ID token included in the authorization header.
	SaveUser(context.Context, *connect.Request[api.SaveUserRequest]) (*connect.Response[api.SaveUserResponse], error)
}

// NewFrontendServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewFrontendServiceHandler(svc FrontendServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	frontendServiceSaveUserHandler := connect.NewUnaryHandler(
		FrontendServiceSaveUserProcedure,
		svc.SaveUser,
		connect.WithSchema(frontendServiceSaveUserMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/frontendapi.FrontendService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case FrontendServiceSaveUserProcedure:
			frontendServiceSaveUserHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedFrontendServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedFrontendServiceHandler struct{}

func (UnimplementedFrontendServiceHandler) SaveUser(context.Context, *connect.Request[api.SaveUserRequest]) (*connect.Response[api.SaveUserResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("frontendapi.FrontendService.SaveUser is not implemented"))
}