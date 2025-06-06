// Code generated by go-swagger; DO NOT EDIT.

// SPDX-FileCopyrightText: Copyright 2025 SAP SE or an SAP affiliate company
//
// SPDX-License-Identifier: Apache-2.0

package members

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/sapcc/andromeda/models"
)

// GetMembersMemberIDHandlerFunc turns a function with the right signature into a get members member ID handler
type GetMembersMemberIDHandlerFunc func(GetMembersMemberIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetMembersMemberIDHandlerFunc) Handle(params GetMembersMemberIDParams) middleware.Responder {
	return fn(params)
}

// GetMembersMemberIDHandler interface for that can handle valid get members member ID params
type GetMembersMemberIDHandler interface {
	Handle(GetMembersMemberIDParams) middleware.Responder
}

// NewGetMembersMemberID creates a new http.Handler for the get members member ID operation
func NewGetMembersMemberID(ctx *middleware.Context, handler GetMembersMemberIDHandler) *GetMembersMemberID {
	return &GetMembersMemberID{Context: ctx, Handler: handler}
}

/*
	GetMembersMemberID swagger:route GET /members/{member_id} Members getMembersMemberId

Show member detail
*/
type GetMembersMemberID struct {
	Context *middleware.Context
	Handler GetMembersMemberIDHandler
}

func (o *GetMembersMemberID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetMembersMemberIDParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetMembersMemberIDOKBody get members member ID o k body
//
// swagger:model GetMembersMemberIDOKBody
type GetMembersMemberIDOKBody struct {

	// member
	Member *models.Member `json:"member,omitempty"`
}

// Validate validates this get members member ID o k body
func (o *GetMembersMemberIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMember(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetMembersMemberIDOKBody) validateMember(formats strfmt.Registry) error {
	if swag.IsZero(o.Member) { // not required
		return nil
	}

	if o.Member != nil {
		if err := o.Member.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getMembersMemberIdOK" + "." + "member")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getMembersMemberIdOK" + "." + "member")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get members member ID o k body based on the context it is used
func (o *GetMembersMemberIDOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateMember(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetMembersMemberIDOKBody) contextValidateMember(ctx context.Context, formats strfmt.Registry) error {

	if o.Member != nil {
		if err := o.Member.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getMembersMemberIdOK" + "." + "member")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getMembersMemberIdOK" + "." + "member")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetMembersMemberIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetMembersMemberIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetMembersMemberIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
