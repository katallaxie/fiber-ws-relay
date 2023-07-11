package adapters

import (
	"github.com/katallaxie/template-go/internal/ports"
)

var _ ports.Noop = (*Noop)(nil)

// Noop is the adapter that implements the Noop interface.
type Noop struct{}

// NewNoop returns a new instance of Noop.
func NewNoop() *Noop {
	return &Noop{}
}

// Get returns a string.
func (n *Noop) Get() string {
	return "Hello, World!"
}
