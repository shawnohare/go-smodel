package collection

// The Collection interface represents a generic abstraction of
// any data source that we can iterate over.
type Interface interface {
	Len() int
	Get(interface{}) interface{}
}
