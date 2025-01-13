package types

import (
	context "context"

	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
)

// Query defines the gRPC querier service.
type QueryServer interface {
	// PendingBundles queries all pending MEV bundles
	PendingBundles(context.Context, *QueryPendingBundlesRequest) (*QueryPendingBundlesResponse, error)
}

// QueryPendingBundlesRequest is the request type for querying pending bundles
type QueryPendingBundlesRequest struct {
}

func (m *QueryPendingBundlesRequest) Reset()         { *m = QueryPendingBundlesRequest{} }
func (m *QueryPendingBundlesRequest) String() string { return proto.CompactTextString(m) }
func (*QueryPendingBundlesRequest) ProtoMessage()    {}

// QueryPendingBundlesResponse is the response type for querying pending bundles
type QueryPendingBundlesResponse struct {
	Bundles []Bundle `protobuf:"bytes,1,rep,name=bundles,proto3" json:"bundles"`
}

func (m *QueryPendingBundlesResponse) Reset()         { *m = QueryPendingBundlesResponse{} }
func (m *QueryPendingBundlesResponse) String() string { return proto.CompactTextString(m) }
func (*QueryPendingBundlesResponse) ProtoMessage()    {}

// QueryClient is the client API for Query service
type QueryClient interface {
	// PendingBundles queries all pending MEV bundles
	PendingBundles(ctx context.Context, in *QueryPendingBundlesRequest, opts ...grpc.CallOption) (*QueryPendingBundlesResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

// NewQueryClient creates a new QueryClient instance
func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) PendingBundles(ctx context.Context, in *QueryPendingBundlesRequest, opts ...grpc.CallOption) (*QueryPendingBundlesResponse, error) {
	out := new(QueryPendingBundlesResponse)
	err := c.cc.Invoke(ctx, "/sei.mev.v1beta1.Query/PendingBundles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}
