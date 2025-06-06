// Code generated by go-swagger; DO NOT EDIT.

// SPDX-FileCopyrightText: Copyright 2025 SAP SE or an SAP affiliate company
//
// SPDX-License-Identifier: Apache-2.0

package monitors

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewPutMonitorsMonitorIDParams creates a new PutMonitorsMonitorIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutMonitorsMonitorIDParams() *PutMonitorsMonitorIDParams {
	return &PutMonitorsMonitorIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutMonitorsMonitorIDParamsWithTimeout creates a new PutMonitorsMonitorIDParams object
// with the ability to set a timeout on a request.
func NewPutMonitorsMonitorIDParamsWithTimeout(timeout time.Duration) *PutMonitorsMonitorIDParams {
	return &PutMonitorsMonitorIDParams{
		timeout: timeout,
	}
}

// NewPutMonitorsMonitorIDParamsWithContext creates a new PutMonitorsMonitorIDParams object
// with the ability to set a context for a request.
func NewPutMonitorsMonitorIDParamsWithContext(ctx context.Context) *PutMonitorsMonitorIDParams {
	return &PutMonitorsMonitorIDParams{
		Context: ctx,
	}
}

// NewPutMonitorsMonitorIDParamsWithHTTPClient creates a new PutMonitorsMonitorIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutMonitorsMonitorIDParamsWithHTTPClient(client *http.Client) *PutMonitorsMonitorIDParams {
	return &PutMonitorsMonitorIDParams{
		HTTPClient: client,
	}
}

/*
PutMonitorsMonitorIDParams contains all the parameters to send to the API endpoint

	for the put monitors monitor ID operation.

	Typically these are written to a http.Request.
*/
type PutMonitorsMonitorIDParams struct {

	// Monitor.
	Monitor PutMonitorsMonitorIDBody

	/* MonitorID.

	   The UUID of the monitor

	   Format: uuid
	*/
	MonitorID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put monitors monitor ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutMonitorsMonitorIDParams) WithDefaults() *PutMonitorsMonitorIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put monitors monitor ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutMonitorsMonitorIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put monitors monitor ID params
func (o *PutMonitorsMonitorIDParams) WithTimeout(timeout time.Duration) *PutMonitorsMonitorIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put monitors monitor ID params
func (o *PutMonitorsMonitorIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put monitors monitor ID params
func (o *PutMonitorsMonitorIDParams) WithContext(ctx context.Context) *PutMonitorsMonitorIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put monitors monitor ID params
func (o *PutMonitorsMonitorIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put monitors monitor ID params
func (o *PutMonitorsMonitorIDParams) WithHTTPClient(client *http.Client) *PutMonitorsMonitorIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put monitors monitor ID params
func (o *PutMonitorsMonitorIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMonitor adds the monitor to the put monitors monitor ID params
func (o *PutMonitorsMonitorIDParams) WithMonitor(monitor PutMonitorsMonitorIDBody) *PutMonitorsMonitorIDParams {
	o.SetMonitor(monitor)
	return o
}

// SetMonitor adds the monitor to the put monitors monitor ID params
func (o *PutMonitorsMonitorIDParams) SetMonitor(monitor PutMonitorsMonitorIDBody) {
	o.Monitor = monitor
}

// WithMonitorID adds the monitorID to the put monitors monitor ID params
func (o *PutMonitorsMonitorIDParams) WithMonitorID(monitorID strfmt.UUID) *PutMonitorsMonitorIDParams {
	o.SetMonitorID(monitorID)
	return o
}

// SetMonitorID adds the monitorId to the put monitors monitor ID params
func (o *PutMonitorsMonitorIDParams) SetMonitorID(monitorID strfmt.UUID) {
	o.MonitorID = monitorID
}

// WriteToRequest writes these params to a swagger request
func (o *PutMonitorsMonitorIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Monitor); err != nil {
		return err
	}

	// path param monitor_id
	if err := r.SetPathParam("monitor_id", o.MonitorID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
