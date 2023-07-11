package ports

// Noop is the interface that wraps the methods to access the
// underlying data store.
type Noop interface {
	Get() string
}
