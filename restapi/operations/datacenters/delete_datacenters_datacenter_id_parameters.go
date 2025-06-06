// Code generated by go-swagger; DO NOT EDIT.

// SPDX-FileCopyrightText: Copyright 2025 SAP SE or an SAP affiliate company
//
// SPDX-License-Identifier: Apache-2.0

package datacenters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewDeleteDatacentersDatacenterIDParams creates a new DeleteDatacentersDatacenterIDParams object
//
// There are no default values defined in the spec.
func NewDeleteDatacentersDatacenterIDParams() DeleteDatacentersDatacenterIDParams {

	return DeleteDatacentersDatacenterIDParams{}
}

// DeleteDatacentersDatacenterIDParams contains all the bound params for the delete datacenters datacenter ID operation
// typically these are obtained from a http.Request
//
// swagger:parameters DeleteDatacentersDatacenterID
type DeleteDatacentersDatacenterIDParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The UUID of the datacenter
	  Required: true
	  In: path
	*/
	DatacenterID strfmt.UUID
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteDatacentersDatacenterIDParams() beforehand.
func (o *DeleteDatacentersDatacenterIDParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rDatacenterID, rhkDatacenterID, _ := route.Params.GetOK("datacenter_id")
	if err := o.bindDatacenterID(rDatacenterID, rhkDatacenterID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindDatacenterID binds and validates parameter DatacenterID from path.
func (o *DeleteDatacentersDatacenterIDParams) bindDatacenterID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("datacenter_id", "path", "strfmt.UUID", raw)
	}
	o.DatacenterID = *(value.(*strfmt.UUID))

	if err := o.validateDatacenterID(formats); err != nil {
		return err
	}

	return nil
}

// validateDatacenterID carries on validations for parameter DatacenterID
func (o *DeleteDatacentersDatacenterIDParams) validateDatacenterID(formats strfmt.Registry) error {

	if err := validate.FormatOf("datacenter_id", "path", "uuid", o.DatacenterID.String(), formats); err != nil {
		return err
	}
	return nil
}
