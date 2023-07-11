package controllers

import (
	"github.com/katallaxie/template-go/internal/ports"
)

// Noop is the controller that uses the Noop interface.
type Noop struct {
	Provider ports.Noop
}

// New returns a new instance of Noop.
func New(provider ports.Noop) *Noop {
	return &Noop{
		Provider: provider,
	}
}

// Get returns a string.
func (n *Noop) Get() string {
	return n.Provider.Get()
}
