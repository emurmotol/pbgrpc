package grpc

import (
	"context"
	"errors"

	endpoint1 "github.com/emurmotol/pbgrpc/ethereum/pkg/endpoint"
	pb "github.com/emurmotol/pbgrpc/ethereum/pkg/grpc/pb"
	service "github.com/emurmotol/pbgrpc/ethereum/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
	grpc1 "github.com/go-kit/kit/transport/grpc"
	grpc "google.golang.org/grpc"
)

// New returns an AddService backed by a gRPC server at the other end
//  of the conn. The caller is responsible for constructing the conn, and
// eventually closing the underlying transport. We bake-in certain middlewares,
// implementing the client library pattern.
func New(conn *grpc.ClientConn, options map[string][]grpc1.ClientOption) (service.EthereumService, error) {
	var createAccountEndpoint endpoint.Endpoint
	{
		createAccountEndpoint = grpc1.NewClient(conn, "pb.Ethereum", "CreateAccount", encodeCreateAccountRequest, decodeCreateAccountResponse, pb.CreateAccountReply{}, options["CreateAccount"]...).Endpoint()
	}

	return endpoint1.Endpoints{CreateAccountEndpoint: createAccountEndpoint}, nil
}

// encodeCreateAccountRequest is a transport/grpc.EncodeRequestFunc that converts a
//  user-domain sum request to a gRPC request.
func encodeCreateAccountRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoint1.CreateAccountRequest)
	return &pb.CreateAccountRequest{Passphrase: req.Passphrase}, nil
}

// decodeCreateAccountResponse is a transport/grpc.DecodeResponseFunc that converts
// a gRPC concat reply to a user-domain concat response.
func decodeCreateAccountResponse(_ context.Context, reply interface{}) (interface{}, error) {
	res := reply.(*pb.CreateAccountReply)
	return endpoint1.CreateAccountResponse{Address: res.Address, Err: errors.New(res.Err)}, nil
}
