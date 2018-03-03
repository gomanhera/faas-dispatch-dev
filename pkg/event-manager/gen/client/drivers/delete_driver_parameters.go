///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package drivers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteDriverParams creates a new DeleteDriverParams object
// with the default values initialized.
func NewDeleteDriverParams() *DeleteDriverParams {
	var ()
	return &DeleteDriverParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDriverParamsWithTimeout creates a new DeleteDriverParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteDriverParamsWithTimeout(timeout time.Duration) *DeleteDriverParams {
	var ()
	return &DeleteDriverParams{

		timeout: timeout,
	}
}

// NewDeleteDriverParamsWithContext creates a new DeleteDriverParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteDriverParamsWithContext(ctx context.Context) *DeleteDriverParams {
	var ()
	return &DeleteDriverParams{

		Context: ctx,
	}
}

// NewDeleteDriverParamsWithHTTPClient creates a new DeleteDriverParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteDriverParamsWithHTTPClient(client *http.Client) *DeleteDriverParams {
	var ()
	return &DeleteDriverParams{
		HTTPClient: client,
	}
}

/*DeleteDriverParams contains all the parameters to send to the API endpoint
for the delete driver operation typically these are written to a http.Request
*/
type DeleteDriverParams struct {

	/*DriverName
	  Name of the driver to work on

	*/
	DriverName string
	/*Tags
	  Filter based on tags

	*/
	Tags []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete driver params
func (o *DeleteDriverParams) WithTimeout(timeout time.Duration) *DeleteDriverParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete driver params
func (o *DeleteDriverParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete driver params
func (o *DeleteDriverParams) WithContext(ctx context.Context) *DeleteDriverParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete driver params
func (o *DeleteDriverParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete driver params
func (o *DeleteDriverParams) WithHTTPClient(client *http.Client) *DeleteDriverParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete driver params
func (o *DeleteDriverParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDriverName adds the driverName to the delete driver params
func (o *DeleteDriverParams) WithDriverName(driverName string) *DeleteDriverParams {
	o.SetDriverName(driverName)
	return o
}

// SetDriverName adds the driverName to the delete driver params
func (o *DeleteDriverParams) SetDriverName(driverName string) {
	o.DriverName = driverName
}

// WithTags adds the tags to the delete driver params
func (o *DeleteDriverParams) WithTags(tags []string) *DeleteDriverParams {
	o.SetTags(tags)
	return o
}

// SetTags adds the tags to the delete driver params
func (o *DeleteDriverParams) SetTags(tags []string) {
	o.Tags = tags
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDriverParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param driverName
	if err := r.SetPathParam("driverName", o.DriverName); err != nil {
		return err
	}

	valuesTags := o.Tags

	joinedTags := swag.JoinByFormat(valuesTags, "multi")
	// query array param tags
	if err := r.SetQueryParam("tags", joinedTags...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
