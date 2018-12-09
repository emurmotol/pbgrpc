package service

import (
	"context"

	"github.com/Sirupsen/logrus"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"
)

// EthereumService describes the service.
type EthereumService interface {
	CreateAccount(ctx context.Context, passphrase string) (address string, err error)
}

type basicEthereumService struct {
	c  *ethclient.Client
	ks *keystore.KeyStore
}

func (b *basicEthereumService) CreateAccount(ctx context.Context, passphrase string) (address string, err error) {
	account, err := b.ks.NewAccount(passphrase)
	return account.Address.String(), err
}

// NewBasicEthereumService returns a naive, stateless implementation of EthereumService.
func NewBasicEthereumService() EthereumService {
	c, err := ethclient.Dial("https://ropsten.infura.io/v3/cb0f868c966e4cf8ae1da8d3a7e92af9")
	if err != nil {
		logrus.Fatal(err)
	}
	ks := keystore.NewKeyStore("/c/Users/emur/Appdata/Roaming/Ethereum/keystore", keystore.StandardScryptN, keystore.StandardScryptP)
	return &basicEthereumService{c: c, ks, ks}
}

// New returns a EthereumService with all of the expected middleware wired in.
func New(middleware []Middleware) EthereumService {
	var svc EthereumService = NewBasicEthereumService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
