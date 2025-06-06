// Code generated by go-swagger; DO NOT EDIT.

// SPDX-FileCopyrightText: Copyright 2025 SAP SE or an SAP affiliate company
//
// SPDX-License-Identifier: Apache-2.0

package pools

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/sapcc/andromeda/models"
)

// PutPoolsPoolIDHandlerFunc turns a function with the right signature into a put pools pool ID handler
type PutPoolsPoolIDHandlerFunc func(PutPoolsPoolIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PutPoolsPoolIDHandlerFunc) Handle(params PutPoolsPoolIDParams) middleware.Responder {
	return fn(params)
}

// PutPoolsPoolIDHandler interface for that can handle valid put pools pool ID params
type PutPoolsPoolIDHandler interface {
	Handle(PutPoolsPoolIDParams) middleware.Responder
}

// NewPutPoolsPoolID creates a new http.Handler for the put pools pool ID operation
func NewPutPoolsPoolID(ctx *middleware.Context, handler PutPoolsPoolIDHandler) *PutPoolsPoolID {
	return &PutPoolsPoolID{Context: ctx, Handler: handler}
}

/*
	PutPoolsPoolID swagger:route PUT /pools/{pool_id} Pools putPoolsPoolId

Update a pool
*/
type PutPoolsPoolID struct {
	Context *middleware.Context
	Handler PutPoolsPoolIDHandler
}

func (o *PutPoolsPoolID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPutPoolsPoolIDParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PutPoolsPoolIDAcceptedBody put pools pool ID accepted body
//
// swagger:model PutPoolsPoolIDAcceptedBody
type PutPoolsPoolIDAcceptedBody struct {

	// pool
	Pool *models.Pool `json:"pool,omitempty"`
}

// Validate validates this put pools pool ID accepted body
func (o *PutPoolsPoolIDAcceptedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePool(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PutPoolsPoolIDAcceptedBody) validatePool(formats strfmt.Registry) error {
	if swag.IsZero(o.Pool) { // not required
		return nil
	}

	if o.Pool != nil {
		if err := o.Pool.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("putPoolsPoolIdAccepted" + "." + "pool")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("putPoolsPoolIdAccepted" + "." + "pool")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this put pools pool ID accepted body based on the context it is used
func (o *PutPoolsPoolIDAcceptedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePool(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PutPoolsPoolIDAcceptedBody) contextValidatePool(ctx context.Context, formats strfmt.Registry) error {

	if o.Pool != nil {
		if err := o.Pool.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("putPoolsPoolIdAccepted" + "." + "pool")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("putPoolsPoolIdAccepted" + "." + "pool")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PutPoolsPoolIDAcceptedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutPoolsPoolIDAcceptedBody) UnmarshalBinary(b []byte) error {
	var res PutPoolsPoolIDAcceptedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PutPoolsPoolIDBody put pools pool ID body
//
// swagger:model PutPoolsPoolIDBody
type PutPoolsPoolIDBody struct {

	// pool
	// Required: true
	Pool *models.Pool `json:"pool"`
}

// Validate validates this put pools pool ID body
func (o *PutPoolsPoolIDBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePool(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PutPoolsPoolIDBody) validatePool(formats strfmt.Registry) error {

	if err := validate.Required("pool"+"."+"pool", "body", o.Pool); err != nil {
		return err
	}

	if o.Pool != nil {
		if err := o.Pool.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pool" + "." + "pool")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pool" + "." + "pool")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this put pools pool ID body based on the context it is used
func (o *PutPoolsPoolIDBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePool(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PutPoolsPoolIDBody) contextValidatePool(ctx context.Context, formats strfmt.Registry) error {

	if o.Pool != nil {
		if err := o.Pool.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pool" + "." + "pool")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pool" + "." + "pool")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PutPoolsPoolIDBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PutPoolsPoolIDBody) UnmarshalBinary(b []byte) error {
	var res PutPoolsPoolIDBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
