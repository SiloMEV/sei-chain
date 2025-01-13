package types

import (
	context "context"

	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// MsgSubmitBundle represents a message to submit a MEV bundle
type MsgSubmitBundle struct {
	Sender    string   `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Txs       []string `protobuf:"bytes,2,rep,name=txs,proto3" json:"txs,omitempty"`
	BlockNum  uint64   `protobuf:"varint,3,opt,name=block_num,json=blockNum,proto3" json:"block_num,omitempty"`
	Timestamp int64    `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (m *MsgSubmitBundle) Reset()         { *m = MsgSubmitBundle{} }
func (m *MsgSubmitBundle) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitBundle) ProtoMessage()    {}

// MsgSubmitBundleResponse defines the response structure
type MsgSubmitBundleResponse struct {
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (m *MsgSubmitBundleResponse) Reset()         { *m = MsgSubmitBundleResponse{} }
func (m *MsgSubmitBundleResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitBundleResponse) ProtoMessage()    {}

// Msg defines the Msg service
type MsgServer interface {
	// SubmitBundle defines a method to submit a new bundle
	SubmitBundle(context.Context, *MsgSubmitBundle) (*MsgSubmitBundleResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations
type UnimplementedMsgServer struct{}

func (*UnimplementedMsgServer) SubmitBundle(context.Context, *MsgSubmitBundle) (*MsgSubmitBundleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitBundle not implemented")
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sei.mev.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitBundle",
			Handler:    _Msg_SubmitBundle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sei/mev/v1beta1/tx.proto",
}

func _Msg_SubmitBundle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSubmitBundle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SubmitBundle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sei.mev.v1beta1.Msg/SubmitBundle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SubmitBundle(ctx, req.(*MsgSubmitBundle))
	}
	return interceptor(ctx, in, info, handler)
}
