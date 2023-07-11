package noop

import (
	"context"

	"github.com/katallaxie/pkg/server"
	"github.com/katallaxie/template-go/internal/controllers"
)

var _ server.Listener = (*NoopSrv)(nil)

// NoopSrv is the server that implements the Noop interface.
type NoopSrv struct {
	ctrl *controllers.Noop
}

// New returns a new instance of NoopSrv.
func New(ctrl *controllers.Noop) *NoopSrv {
	return &NoopSrv{
		ctrl: ctrl,
	}
}

// Start starts the server.
func (n *NoopSrv) Start(ctx context.Context, ready server.ReadyFunc, run server.RunFunc) func() error {
	return func() error {
		ready()

		<-ctx.Done()

		return nil
	}
}
