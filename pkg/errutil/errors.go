package errutil

type KeyError struct {
	Key        string
	StatusCode int
}

func NewKeyError(key string, statusCode int) *KeyError {
	return &KeyError{
		Key:        key,
		StatusCode: statusCode,
	}
}

func (e *KeyError) Error() string {
	return e.Key
}
