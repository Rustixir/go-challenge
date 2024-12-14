package errutil

import "net/http"

var (
	ErrNotFound            = NewKeyError("not found", http.StatusNotFound)
	ErrInternalServerError = NewKeyError("internal_server_error", http.StatusInternalServerError)
	ErrInvalidSegment      = NewKeyError("invalid_segment", http.StatusBadRequest)
	ErrInvalidUser         = NewKeyError("invalid_user", http.StatusBadRequest)
	ErrInvalidRequest      = NewKeyError("invalid_request", http.StatusBadRequest)
)
