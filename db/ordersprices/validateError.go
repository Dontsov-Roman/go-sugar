package ordersprices

// ValidateError struct for validate user by table
type ValidateError struct {
	OrderIDPriceID string
	ErrorMessage   string
}

// AddToErrorMessage concat string to ErrorMessage
func (e *ValidateError) AddToErrorMessage(msg string) {
	e.ErrorMessage = e.ErrorMessage + msg
	e.addDotToErrorMessage()
}
func (e *ValidateError) addDotToErrorMessage() {
	if e.ErrorMessage != "" {
		e.ErrorMessage = e.ErrorMessage + ". "
	}
}
