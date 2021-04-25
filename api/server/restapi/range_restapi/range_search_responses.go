// Code generated by go-swagger; DO NOT EDIT.

package range_restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// RangeSearchOKCode is the HTTP code returned for type RangeSearchOK
const RangeSearchOKCode int = 200

/*RangeSearchOK Request was processed successfully.

swagger:response rangeSearchOK
*/
type RangeSearchOK struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewRangeSearchOK creates RangeSearchOK with default headers values
func NewRangeSearchOK() *RangeSearchOK {

	return &RangeSearchOK{}
}

// WithPayload adds the payload to the range search o k response
func (o *RangeSearchOK) WithPayload(payload string) *RangeSearchOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the range search o k response
func (o *RangeSearchOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RangeSearchOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// RangeSearchBadRequestCode is the HTTP code returned for type RangeSearchBadRequest
const RangeSearchBadRequestCode int = 400

/*RangeSearchBadRequest Hash prefix must be exactly 5 characters long.

swagger:response rangeSearchBadRequest
*/
type RangeSearchBadRequest struct {
}

// NewRangeSearchBadRequest creates RangeSearchBadRequest with default headers values
func NewRangeSearchBadRequest() *RangeSearchBadRequest {

	return &RangeSearchBadRequest{}
}

// WriteResponse to the client
func (o *RangeSearchBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(400)
}

// RangeSearchNotFoundCode is the HTTP code returned for type RangeSearchNotFound
const RangeSearchNotFoundCode int = 404

/*RangeSearchNotFound No results found.

swagger:response rangeSearchNotFound
*/
type RangeSearchNotFound struct {
}

// NewRangeSearchNotFound creates RangeSearchNotFound with default headers values
func NewRangeSearchNotFound() *RangeSearchNotFound {

	return &RangeSearchNotFound{}
}

// WriteResponse to the client
func (o *RangeSearchNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// RangeSearchInternalServerErrorCode is the HTTP code returned for type RangeSearchInternalServerError
const RangeSearchInternalServerErrorCode int = 500

/*RangeSearchInternalServerError Server encountered an error.

swagger:response rangeSearchInternalServerError
*/
type RangeSearchInternalServerError struct {
}

// NewRangeSearchInternalServerError creates RangeSearchInternalServerError with default headers values
func NewRangeSearchInternalServerError() *RangeSearchInternalServerError {

	return &RangeSearchInternalServerError{}
}

// WriteResponse to the client
func (o *RangeSearchInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
