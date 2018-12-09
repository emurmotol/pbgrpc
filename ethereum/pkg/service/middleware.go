package service

import (
	"context"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(EthereumService) EthereumService

type loggingMiddleware struct {
	logger log.Logger
	next   EthereumService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a EthereumService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next EthereumService) EthereumService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) CreateAccount(ctx context.Context, passphrase string) (address string, err error) {
	defer func() {
		l.logger.Log("method", "CreateAccount", "passphrase", passphrase, "address", address, "err", err)
	}()
	return l.next.CreateAccount(ctx, passphrase)
}
