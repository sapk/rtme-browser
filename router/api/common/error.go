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
}

// BadRequest is an bad request error response
// swagger:response BadRequest
type BadRequest struct {
	// The error
}

// ErrorNotFound is a not found format response
// swagger:response ErrorNotFound
type ErrorNotFound struct {
	// The error
	// in: body
	Body Error
}
