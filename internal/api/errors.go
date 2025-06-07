package api

import (
	"errors"
	"fmt"
	"net/http"
)

// UrlqueryApiError represents an error returned by the API.
type UrlqueryApiError struct {
	StatusCode int
	Message    string
}

func (e *UrlqueryApiError) Error() string {
	return fmt.Sprintf("API Error (HTTP StatusCode: %d) %s", e.StatusCode, e.Message)
}

var (
	ErrNotFound            = errors.New("not found")
	ErrForbidden           = errors.New("forbidden")
	ErrBadRequest          = errors.New("bad request")
	ErrUnauthorized        = errors.New("unauthorized")
	ErrNotAcceptable       = errors.New("not acceptable")
	ErrTooManyRequests     = errors.New("too many requests")
	ErrUnprocessableEntity = errors.New("unprocessable entity")
	ErrUnexpectedStatus    = errors.New("unexpected status code")
)

func handleResponseError(resp *http.Response) error {
	switch resp.StatusCode {
	// Success status codes
	case http.StatusOK,
		http.StatusCreated,
		http.StatusAccepted,
		http.StatusNoContent:
		return nil

	// Client errors
	case http.StatusBadRequest:
		return &UrlqueryApiError{
			StatusCode: resp.StatusCode,
			Message:    fmt.Sprintf("%s: The request was malformed or invalid", ErrBadRequest.Error()),
		}

	case http.StatusUnauthorized:
		return &UrlqueryApiError{
			StatusCode: resp.StatusCode,
			Message:    fmt.Sprintf("%s: Invalid or missing API key", ErrUnauthorized.Error()),
		}

	case http.StatusForbidden:
		return &UrlqueryApiError{
			StatusCode: resp.StatusCode,
			Message:    fmt.Sprintf("%s: Access denied - insufficient permissions", ErrForbidden.Error()),
		}

	case http.StatusNotFound:
		return &UrlqueryApiError{
			StatusCode: resp.StatusCode,
			Message:    fmt.Sprintf("%s: The requested resource was not found", ErrNotFound.Error()),
		}

	case http.StatusNotAcceptable:
		return &UrlqueryApiError{
			StatusCode: resp.StatusCode,
			Message:    fmt.Sprintf("%s: The request format is not acceptable", ErrNotAcceptable.Error()),
		}

	case http.StatusUnprocessableEntity:
		return &UrlqueryApiError{
			StatusCode: resp.StatusCode,
			Message:    fmt.Sprintf("%s: The request data could not be processed", ErrUnprocessableEntity.Error()),
		}

	case http.StatusTooManyRequests:
		return &UrlqueryApiError{
			StatusCode: resp.StatusCode,
			Message:    fmt.Sprintf("%s: Rate limit exceeded - please try again later", ErrTooManyRequests.Error()),
		}

	// Server errors
	case http.StatusInternalServerError:
		return &UrlqueryApiError{
			StatusCode: resp.StatusCode,
			Message:    "Internal server error - please try again later",
		}

	case http.StatusBadGateway:
		return &UrlqueryApiError{
			StatusCode: resp.StatusCode,
			Message:    "Bad gateway - service temporarily unavailable",
		}

	case http.StatusServiceUnavailable:
		return &UrlqueryApiError{
			StatusCode: resp.StatusCode,
			Message:    "Service unavailable - please try again later",
		}

	case http.StatusGatewayTimeout:
		return &UrlqueryApiError{
			StatusCode: resp.StatusCode,
			Message:    "Gateway timeout - request took too long to process",
		}

	default:
		return &UrlqueryApiError{
			StatusCode: resp.StatusCode,
			Message:    fmt.Sprintf("%s: HTTP %d", ErrUnexpectedStatus.Error(), resp.StatusCode),
		}
	}
}
