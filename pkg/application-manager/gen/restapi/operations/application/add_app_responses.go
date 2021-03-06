///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vmware/dispatch/pkg/application-manager/gen/models"
)

// AddAppOKCode is the HTTP code returned for type AddAppOK
const AddAppOKCode int = 200

/*AddAppOK Application created

swagger:response addAppOK
*/
type AddAppOK struct {

	/*
	  In: Body
	*/
	Payload *models.Application `json:"body,omitempty"`
}

// NewAddAppOK creates AddAppOK with default headers values
func NewAddAppOK() *AddAppOK {
	return &AddAppOK{}
}

// WithPayload adds the payload to the add app o k response
func (o *AddAppOK) WithPayload(payload *models.Application) *AddAppOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add app o k response
func (o *AddAppOK) SetPayload(payload *models.Application) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddAppOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddAppBadRequestCode is the HTTP code returned for type AddAppBadRequest
const AddAppBadRequestCode int = 400

/*AddAppBadRequest Invalid Input

swagger:response addAppBadRequest
*/
type AddAppBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddAppBadRequest creates AddAppBadRequest with default headers values
func NewAddAppBadRequest() *AddAppBadRequest {
	return &AddAppBadRequest{}
}

// WithPayload adds the payload to the add app bad request response
func (o *AddAppBadRequest) WithPayload(payload *models.Error) *AddAppBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add app bad request response
func (o *AddAppBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddAppBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddAppUnauthorizedCode is the HTTP code returned for type AddAppUnauthorized
const AddAppUnauthorizedCode int = 401

/*AddAppUnauthorized Unauthorized Request

swagger:response addAppUnauthorized
*/
type AddAppUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddAppUnauthorized creates AddAppUnauthorized with default headers values
func NewAddAppUnauthorized() *AddAppUnauthorized {
	return &AddAppUnauthorized{}
}

// WithPayload adds the payload to the add app unauthorized response
func (o *AddAppUnauthorized) WithPayload(payload *models.Error) *AddAppUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add app unauthorized response
func (o *AddAppUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddAppUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddAppConflictCode is the HTTP code returned for type AddAppConflict
const AddAppConflictCode int = 409

/*AddAppConflict Already Exists

swagger:response addAppConflict
*/
type AddAppConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddAppConflict creates AddAppConflict with default headers values
func NewAddAppConflict() *AddAppConflict {
	return &AddAppConflict{}
}

// WithPayload adds the payload to the add app conflict response
func (o *AddAppConflict) WithPayload(payload *models.Error) *AddAppConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add app conflict response
func (o *AddAppConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddAppConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddAppInternalServerErrorCode is the HTTP code returned for type AddAppInternalServerError
const AddAppInternalServerErrorCode int = 500

/*AddAppInternalServerError Internal Error

swagger:response addAppInternalServerError
*/
type AddAppInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddAppInternalServerError creates AddAppInternalServerError with default headers values
func NewAddAppInternalServerError() *AddAppInternalServerError {
	return &AddAppInternalServerError{}
}

// WithPayload adds the payload to the add app internal server error response
func (o *AddAppInternalServerError) WithPayload(payload *models.Error) *AddAppInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add app internal server error response
func (o *AddAppInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddAppInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
