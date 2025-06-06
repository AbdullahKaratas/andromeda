// Code generated by go-swagger; DO NOT EDIT.

// SPDX-FileCopyrightText: Copyright 2025 SAP SE or an SAP affiliate company
//
// SPDX-License-Identifier: Apache-2.0

package pools

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewPutPoolsPoolIDParams creates a new PutPoolsPoolIDParams object
//
// There are no default values defined in the spec.
func NewPutPoolsPoolIDParams() PutPoolsPoolIDParams {

	return PutPoolsPoolIDParams{}
}

// PutPoolsPoolIDParams contains all the bound params for the put pools pool ID operation
// typically these are obtained from a http.Request
//
// swagger:parameters PutPoolsPoolID
type PutPoolsPoolIDParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: body
	*/
	Pool PutPoolsPoolIDBody
	/*The UUID of the pool
	  Required: true
	  In: path
	*/
	PoolID strfmt.UUID
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPutPoolsPoolIDParams() beforehand.
func (o *PutPoolsPoolIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body PutPoolsPoolIDBody
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("pool", "body", ""))
			} else {
				res = append(res, errors.NewParseError("pool", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(r.Context())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Pool = body
			}
		}
	} else {
		res = append(res, errors.Required("pool", "body", ""))
	}

	rPoolID, rhkPoolID, _ := route.Params.GetOK("pool_id")
	if err := o.bindPoolID(rPoolID, rhkPoolID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindPoolID binds and validates parameter PoolID from path.
func (o *PutPoolsPoolIDParams) bindPoolID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("pool_id", "path", "strfmt.UUID", raw)
	}
	o.PoolID = *(value.(*strfmt.UUID))

	if err := o.validatePoolID(formats); err != nil {
		return err
	}

	return nil
}

// validatePoolID carries on validations for parameter PoolID
func (o *PutPoolsPoolIDParams) validatePoolID(formats strfmt.Registry) error {

	if err := validate.FormatOf("pool_id", "path", "uuid", o.PoolID.String(), formats); err != nil {
		return err
	}
	return nil
}
