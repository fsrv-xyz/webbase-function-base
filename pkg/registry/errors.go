package registry

type FunctionAlreadyRegisteredError struct{}

func (e *FunctionAlreadyRegisteredError) Error() string {
	return "function already registered"
}
