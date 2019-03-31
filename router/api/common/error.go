package common

// Error is error format response
// swagger:response Error
type Error struct {
	Error string `json:"error"`
}

// ErrorUnauthorized is an unauthorized error response
// swagger:response ErrorUnauthorized
type ErrorUnauthorized struct {
	// The error
	// in: body
	Body Error
}

// BadRequest is an bad request error response
// swagger:response BadRequest
type BadRequest struct {
	// The error
	// in: body
	Body Error
}

// ErrorNotFound is a not found format response
// swagger:response ErrorNotFound
type ErrorNotFound struct {
	// The error
	// in: body
	Body Error
}

// InternalServerError is a InternalServerError format response
// swagger:response InternalServerError
type InternalServerError struct {
	// The error
	// in: body
	Body Error
}
