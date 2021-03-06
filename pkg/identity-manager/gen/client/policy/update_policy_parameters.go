///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/vmware/dispatch/pkg/identity-manager/gen/models"
)

// NewUpdatePolicyParams creates a new UpdatePolicyParams object
// with the default values initialized.
func NewUpdatePolicyParams() *UpdatePolicyParams {
	var ()
	return &UpdatePolicyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdatePolicyParamsWithTimeout creates a new UpdatePolicyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdatePolicyParamsWithTimeout(timeout time.Duration) *UpdatePolicyParams {
	var ()
	return &UpdatePolicyParams{

		timeout: timeout,
	}
}

// NewUpdatePolicyParamsWithContext creates a new UpdatePolicyParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdatePolicyParamsWithContext(ctx context.Context) *UpdatePolicyParams {
	var ()
	return &UpdatePolicyParams{

		Context: ctx,
	}
}

// NewUpdatePolicyParamsWithHTTPClient creates a new UpdatePolicyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdatePolicyParamsWithHTTPClient(client *http.Client) *UpdatePolicyParams {
	var ()
	return &UpdatePolicyParams{
		HTTPClient: client,
	}
}

/*UpdatePolicyParams contains all the parameters to send to the API endpoint
for the update policy operation typically these are written to a http.Request
*/
type UpdatePolicyParams struct {

	/*Body
	  Policy object

	*/
	Body *models.Policy
	/*PolicyName
	  Name of Policy to work on

	*/
	PolicyName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update policy params
func (o *UpdatePolicyParams) WithTimeout(timeout time.Duration) *UpdatePolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update policy params
func (o *UpdatePolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update policy params
func (o *UpdatePolicyParams) WithContext(ctx context.Context) *UpdatePolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update policy params
func (o *UpdatePolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update policy params
func (o *UpdatePolicyParams) WithHTTPClient(client *http.Client) *UpdatePolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update policy params
func (o *UpdatePolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update policy params
func (o *UpdatePolicyParams) WithBody(body *models.Policy) *UpdatePolicyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update policy params
func (o *UpdatePolicyParams) SetBody(body *models.Policy) {
	o.Body = body
}

// WithPolicyName adds the policyName to the update policy params
func (o *UpdatePolicyParams) WithPolicyName(policyName string) *UpdatePolicyParams {
	o.SetPolicyName(policyName)
	return o
}

// SetPolicyName adds the policyName to the update policy params
func (o *UpdatePolicyParams) SetPolicyName(policyName string) {
	o.PolicyName = policyName
}

// WriteToRequest writes these params to a swagger request
func (o *UpdatePolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models.Policy)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param policyName
	if err := r.SetPathParam("policyName", o.PolicyName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
