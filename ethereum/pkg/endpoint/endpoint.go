package endpoint

import (
	"context"

	service "github.com/emurmotol/pbgrpc/ethereum/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
)

// CreateAccountRequest collects the request parameters for the CreateAccount method.
type CreateAccountRequest struct {
	Passphrase string `json:"passphrase"`
}

// CreateAccountResponse collects the response parameters for the CreateAccount method.
type CreateAccountResponse struct {
	Address string `json:"address"`
	Err     error  `json:"err"`
}

// MakeCreateAccountEndpoint returns an endpoint that invokes CreateAccount on the service.
func MakeCreateAccountEndpoint(s service.EthereumService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateAccountRequest)
		address, err := s.CreateAccount(ctx, req.Passphrase)
		return CreateAccountResponse{
			Address: address,
			Err:     err,
		}, nil
	}
}

// Failed implements Failer.
func (r CreateAccountResponse) Failed() error {
	return r.Err
}

// Failer is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// CreateAccount implements Service. Primarily useful in a client.
func (e Endpoints) CreateAccount(ctx context.Context, passphrase string) (address string, err error) {
	request := CreateAccountRequest{Passphrase: passphrase}
	response, err := e.CreateAccountEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateAccountResponse).Address, response.(CreateAccountResponse).Err
}
