///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package organization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/vmware/dispatch/pkg/api/v1"
)

// GetOrganizationsReader is a Reader for the GetOrganizations structure.
type GetOrganizationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetOrganizationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetOrganizationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetOrganizationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetOrganizationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetOrganizationsOK creates a GetOrganizationsOK with default headers values
func NewGetOrganizationsOK() *GetOrganizationsOK {
	return &GetOrganizationsOK{}
}

/*GetOrganizationsOK handles this case with default header values.

Successful operation
*/
type GetOrganizationsOK struct {
	Payload []*v1.Organization
}

func (o *GetOrganizationsOK) Error() string {
	return fmt.Sprintf("[GET /v1/iam/organization][%d] getOrganizationsOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsUnauthorized creates a GetOrganizationsUnauthorized with default headers values
func NewGetOrganizationsUnauthorized() *GetOrganizationsUnauthorized {
	return &GetOrganizationsUnauthorized{}
}

/*GetOrganizationsUnauthorized handles this case with default header values.

Unauthorized Request
*/
type GetOrganizationsUnauthorized struct {
	Payload *v1.Error
}

func (o *GetOrganizationsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/iam/organization][%d] getOrganizationsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOrganizationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsForbidden creates a GetOrganizationsForbidden with default headers values
func NewGetOrganizationsForbidden() *GetOrganizationsForbidden {
	return &GetOrganizationsForbidden{}
}

/*GetOrganizationsForbidden handles this case with default header values.

access to this resource is forbidden
*/
type GetOrganizationsForbidden struct {
	Payload *v1.Error
}

func (o *GetOrganizationsForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/iam/organization][%d] getOrganizationsForbidden  %+v", 403, o.Payload)
}

func (o *GetOrganizationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsDefault creates a GetOrganizationsDefault with default headers values
func NewGetOrganizationsDefault(code int) *GetOrganizationsDefault {
	return &GetOrganizationsDefault{
		_statusCode: code,
	}
}

/*GetOrganizationsDefault handles this case with default header values.

Unexpected Error
*/
type GetOrganizationsDefault struct {
	_statusCode int

	Payload *v1.Error
}

// Code gets the status code for the get organizations default response
func (o *GetOrganizationsDefault) Code() int {
	return o._statusCode
}

func (o *GetOrganizationsDefault) Error() string {
	return fmt.Sprintf("[GET /v1/iam/organization][%d] getOrganizations default  %+v", o._statusCode, o.Payload)
}

func (o *GetOrganizationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(v1.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}