// Package list provides an ordered set interface.
package list

type Interface interface {
	Len() int
	Get(i int) interface{}
}

type List struct {
	vals []interface{}
}

func (l *List) Len() int {
	return len(l.vals)
}

func (l *List) Get(i int) interface{} {
	return l.vals[i]
}

func (l *List) Set(i int, x interface{}) {
	l.vals[i] = x
}

func (l *List) Append(x interface{}) {
	l.vals = append(l.vals, x)
}

// New creates a
func New(n int) *List {
	vals := make([]interface{}, n)
	return &List{vals}
}
