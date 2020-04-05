package errors

type BusinessServiceError struct {
	Errorcode string
	Message string
	Cause string
}
