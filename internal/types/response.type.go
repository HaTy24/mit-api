package types

type OperationResult struct {
	Success  bool
	Data     interface{}
	Message  interface{}
	Code     string
	HttpCode int
}
