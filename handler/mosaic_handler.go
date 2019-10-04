package handler

// MosaicLambda represent the main lambda and it's dependencies
type MosaicLambda struct {
	Name string `json:"name,omitempty"`
}

// HandleRequest is the main entry point for the lambda function
func (m *MosaicLambda) HandleRequest() (int, error) {
	return 0, nil
}
