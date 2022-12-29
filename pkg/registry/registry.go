package registry

import (
	"net/http"
)

type Function struct {
	HttpFunction http.HandlerFunc
}

type Registry struct {
	functions map[string]*Function
}

func NewRegistry() *Registry {
	return &Registry{
		functions: make(map[string]*Function),
	}
}

var defaultRegistry = NewRegistry()

func (r *Registry) RegisterHttpFunction(name string, f http.HandlerFunc) error {
	if _, ok := r.functions[name]; ok {
		return &FunctionAlreadyRegisteredError{}
	}
	r.functions[name] = &Function{HttpFunction: f}
	return nil
}

func GetRegisteredFunctions() map[string]*Function {
	return defaultRegistry.functions
}

func RegisterHttpFunction(name string, f http.HandlerFunc) error {
	return defaultRegistry.RegisterHttpFunction(name, f)
}
