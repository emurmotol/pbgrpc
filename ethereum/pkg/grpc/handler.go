package grpc

import (
	"context"

	endpoint "github.com/emurmotol/pbgrpc/ethereum/pkg/endpoint"
	pb "github.com/emurmotol/pbgrpc/ethereum/pkg/grpc/pb"
	grpc "github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"
)

// makeCreateAccountHandler creates the handler logic
func makeCreateAccountHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.CreateAccountEndpoint, decodeCreateAccountRequest, encodeCreateAccountResponse, options...)
}

// decodeCreateAccountResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain sum request.
func decodeCreateAccountRequest(_ context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.CreateAccountRequest)
	return endpoint.CreateAccountRequest{Passphrase: req.Passphrase}, nil
}

// encodeCreateAccountResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
func encodeCreateAccountResponse(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(endpoint.CreateAccountResponse)
	var err string
	if res.Err != nil {
		err = res.Err.Error()
	}
	return &pb.CreateAccountReply{Address: res.Address, Err: err}, nil
}

func (g *grpcServer) CreateAccount(ctx context1.Context, req *pb.CreateAccountRequest) (*pb.CreateAccountReply, error) {
	_, rep, err := g.createAccount.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CreateAccountReply), nil
}
