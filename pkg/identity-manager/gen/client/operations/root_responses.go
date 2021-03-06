///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/vmware/dispatch/pkg/identity-manager/gen/models"
)

// RootReader is a Reader for the Root structure.
type RootReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RootReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRootOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewRootDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRootOK creates a RootOK with default headers values
func NewRootOK() *RootOK {
	return &RootOK{}
}

/*RootOK handles this case with default header values.

home page
*/
type RootOK struct {
	Payload *models.Message
}

func (o *RootOK) Error() string {
	return fmt.Sprintf("[GET /][%d] rootOK  %+v", 200, o.Payload)
}

func (o *RootOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Message)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRootDefault creates a RootDefault with default headers values
func NewRootDefault(code int) *RootDefault {
	return &RootDefault{
		_statusCode: code,
	}
}

/*RootDefault handles this case with default header values.

error
*/
type RootDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the root default response
func (o *RootDefault) Code() int {
	return o._statusCode
}

func (o *RootDefault) Error() string {
	return fmt.Sprintf("[GET /][%d] root default  %+v", o._statusCode, o.Payload)
}

func (o *RootDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
