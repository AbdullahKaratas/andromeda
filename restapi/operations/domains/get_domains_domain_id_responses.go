// Code generated by go-swagger; DO NOT EDIT.

// SPDX-FileCopyrightText: Copyright 2025 SAP SE or an SAP affiliate company
//
// SPDX-License-Identifier: Apache-2.0

package domains

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/sapcc/andromeda/models"
)

// GetDomainsDomainIDOKCode is the HTTP code returned for type GetDomainsDomainIDOK
const GetDomainsDomainIDOKCode int = 200

/*
GetDomainsDomainIDOK Shows the details of a specific domain.

swagger:response getDomainsDomainIdOK
*/
type GetDomainsDomainIDOK struct {

	/*
	  In: Body
	*/
	Payload *GetDomainsDomainIDOKBody `json:"body,omitempty"`
}

// NewGetDomainsDomainIDOK creates GetDomainsDomainIDOK with default headers values
func NewGetDomainsDomainIDOK() *GetDomainsDomainIDOK {

	return &GetDomainsDomainIDOK{}
}

// WithPayload adds the payload to the get domains domain Id o k response
func (o *GetDomainsDomainIDOK) WithPayload(payload *GetDomainsDomainIDOKBody) *GetDomainsDomainIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get domains domain Id o k response
func (o *GetDomainsDomainIDOK) SetPayload(payload *GetDomainsDomainIDOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDomainsDomainIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetDomainsDomainIDNotFoundCode is the HTTP code returned for type GetDomainsDomainIDNotFound
const GetDomainsDomainIDNotFoundCode int = 404

/*
GetDomainsDomainIDNotFound Not Found

swagger:response getDomainsDomainIdNotFound
*/
type GetDomainsDomainIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetDomainsDomainIDNotFound creates GetDomainsDomainIDNotFound with default headers values
func NewGetDomainsDomainIDNotFound() *GetDomainsDomainIDNotFound {

	return &GetDomainsDomainIDNotFound{}
}

// WithPayload adds the payload to the get domains domain Id not found response
func (o *GetDomainsDomainIDNotFound) WithPayload(payload *models.Error) *GetDomainsDomainIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get domains domain Id not found response
func (o *GetDomainsDomainIDNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDomainsDomainIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
GetDomainsDomainIDDefault Unexpected Error

swagger:response getDomainsDomainIdDefault
*/
type GetDomainsDomainIDDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetDomainsDomainIDDefault creates GetDomainsDomainIDDefault with default headers values
func NewGetDomainsDomainIDDefault(code int) *GetDomainsDomainIDDefault {
	if code <= 0 {
		code = 500
	}

	return &GetDomainsDomainIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get domains domain ID default response
func (o *GetDomainsDomainIDDefault) WithStatusCode(code int) *GetDomainsDomainIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get domains domain ID default response
func (o *GetDomainsDomainIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get domains domain ID default response
func (o *GetDomainsDomainIDDefault) WithPayload(payload *models.Error) *GetDomainsDomainIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get domains domain ID default response
func (o *GetDomainsDomainIDDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDomainsDomainIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
