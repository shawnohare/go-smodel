// Package smodel provides a simple Go interface for statistical models
// as well as some particular implementations used to score nodes in a
// social network.
package smodel

// Interface for a generic statistical model.
type Interface interface {
	Eval(interface{}) interface{}
	// TODO
	// - Train method that accepts a collection of input data
	//   and modifies the model params to find a best fit.
	// - Anything more specific should probably be off-loaded to
	//   less general model interfaces (e.g., a Linear Regression interface.)
}

// The Eval function applys the model to a single input.
func Eval(m Interface, x interface{}) interface{} {
	return m.Eval(x)
}

// Map the model over a collection of inputs.
func Map(m Interface, inputs []interface{}) []interface{} {
	outputs := make([]interface{}, len(inputs))
	for i, x := range inputs {
		outputs[i] = Eval(m, x)
	}
	return outputs
}

// Map the model over a list of inputs.
// func Map(m Interface, l list.Interface) list.Interface {
// TODO
// }

// Map a model over a collection of inputs.
// func Map(m Interface, c Collection) Collection {
// TODO
// 	for i := 0; i < c.Len(); i++ {

// 	}
// }
