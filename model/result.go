package model

type ExecutionResult struct {
	Image string

	Success bool

	Cached bool

	Error error
}
