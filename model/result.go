package model

type ExecutionResult struct {
	Image string

	Target string

	Success bool

	Cached bool

	Error error
}
